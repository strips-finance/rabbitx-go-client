// Package client implements a simple RabbitX client.
// It can be used with either a private key (optional) for automatic onboarding,
// or with a secret key only, which allows for trading operations but not withdrawals.
package client

import (
	"crypto/ecdsa"
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
	"os"
	"rabbitx-client/model"
	"strconv"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/crypto"
)

// RbClient is a demo client.
// If a private key is set up, it performs onboarding.
// If only an apiSecret is set up, it starts trading.
type RbClient struct {
	wallet       string
	apiUrl       string
	httpClient   http.Client
	privateKey   *ecdsa.PrivateKey
	refreshToken string
	jwtPrivate   string
	apiSecret    *model.APISecret
	mu           sync.Mutex
}

// NewRbClient creates a new RbClient instance.
// It accepts a private key or secret and performs onboarding based on that.
func NewRbClient(apiUrl, wallet, privateKey, apiKey, apiSecret, refreshToken, jwtPrivate string, keyExpired int64) *RbClient {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	rc := &RbClient{
		wallet:       wallet,
		apiUrl:       apiUrl,
		httpClient:   http.Client{Jar: jar},
		refreshToken: refreshToken,
		jwtPrivate:   jwtPrivate,
	}

	if privateKey != "" {
		if strings.HasPrefix(privateKey, "0x") {
			privateKey = privateKey[2:]
		}

		if len(privateKey) == 0 {
			panic("Invalid private key")
		}

		pk, err := crypto.HexToECDSA(privateKey)
		if err != nil {
			panic(err)
		}

		rc.privateKey = pk
	}

	isValid := apiKey != "" && refreshToken != "" && apiSecret != ""
	if isValid {
		rc.apiSecret = &model.APISecret{
			Key:        apiKey,
			Secret:     apiSecret,
			Expiration: uint(keyExpired),
		}
	}

	return rc
}

// GetSecrets retrieves the API secret key and secret.
// If the API secret is expired or nil, it will be updated automatically.
func (c *RbClient) GetSecrets() (apiKey string, apiSecret string, jwtPrivate string, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !isApiSecretExpired(c.apiSecret) {
		// Check if close to expired update it
		if isCloseToExpired(c.apiSecret) {
			secret, e := c.RefreshSecrets(c.apiSecret.Key, c.apiSecret.Secret, c.refreshToken)
			if e == nil {
				c.updateSecrets(secret.APISecret, secret.JwtPrivate, secret.RefreshToken)
			} else {
				return "", "", "", e
			}

		}
		return c.apiSecret.Key, c.apiSecret.Secret, c.jwtPrivate, nil
	}

	// Key expired we can update only by onboarding
	res, e := c.Onboarding(c.wallet, c.privateKey)
	if e == nil {
		c.updateSecrets(res.APISecret, res.Jwt, c.refreshToken)
		return c.apiSecret.Key, c.apiSecret.Secret, c.jwtPrivate, nil
	}

	return "", "", "", e
}

// SaveSecrets saves the API secrets to a file.
// It returns an error if the file cannot be opened or the secrets cannot be written to the file.
func (c *RbClient) SaveSecrets(fileName string) error {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	secrets := map[string]string{
		"apiSecretKey":    c.apiSecret.Key,
		"apiSecretSecret": c.apiSecret.Secret,
		"refreshToken":    c.refreshToken,
		"jwtPrivate":      c.jwtPrivate,
		"expiration":      strconv.FormatUint(uint64(c.apiSecret.Expiration), 10),
	}

	secretsJson, err := json.Marshal(secrets)
	if err != nil {
		return err
	}

	_, err = file.Write(secretsJson)
	if err != nil {
		return err
	}

	return nil
}

// updateSecrets updates the API secrets, JWT private key, and refresh token.
func (c *RbClient) updateSecrets(newApiSecret *model.APISecret, jwtPrivate, refreshtoken string) {
	c.apiSecret = newApiSecret
	c.jwtPrivate = jwtPrivate
	c.refreshToken = refreshtoken
}
