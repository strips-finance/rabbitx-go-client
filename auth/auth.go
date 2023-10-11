package auth

import (
	"crypto/ecdsa"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"sort"
	"strconv"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	ONBOARDING_MESSAGE = "Welcome to Rabbit DEX"
)

// OnboardingSiganture converts internal message string to metamask
// EIP-191 format. So later message can be verified as signed
// by Ethereum secret key holder via metamask.
// timestamp is used to verify every message
// signed via ECDSA algo by metamask.
func OnboardingSiganture(privateKey *ecdsa.PrivateKey, timestamp int64) (string, error) {
	metamaskMessage := fmt.Sprintf("%s\n%d", ONBOARDING_MESSAGE, timestamp)
	eip191Message := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(metamaskMessage), metamaskMessage)

	message := crypto.Keccak256Hash([]byte(eip191Message)).Bytes()
	signature, err := crypto.Sign(message, privateKey)
	if err != nil {
		return "", err
	}

	return hexutil.Encode(signature), nil
}

// PayloadSignature returns HMAC-SHA256 signature after signing payload hash with
// provided by user secret. Later this signature is used to ensure
// signer of payload was valid. From high level overview payload
// can be signed by frontend user by rotating random secret or by
// market maker with constant api key secret.
func PayloadSignature(payload map[string]string, secret string, timestamp int64) (string, error) {
	secretBytes, err := hexutil.Decode(secret)
	if err != nil {
		return "", err
	}

	// Sort payload keys and prepare an alphabetically ordered string.
	var message string
	sortedKeys := make([]string, 0, len(payload))
	for k := range payload {
		sortedKeys = append(sortedKeys, k)
	}

	sort.Strings(sortedKeys)
	for _, k := range sortedKeys {
		message += fmt.Sprintf("%s=%s", k, payload[k])
	}

	timestampStr := strconv.FormatInt(timestamp, 10)

	// Calculate hash itself with given input.
	input := make([]byte, 0, len(message)+len(timestampStr))
	input = append(input, []byte(message)...)
	input = append(input, []byte(timestampStr)...)
	hash := sha256.Sum256(input)

	mac := hmac.New(sha256.New, secretBytes)
	mac.Write(hash[:])
	signature := mac.Sum(nil)

	return hexutil.Encode(signature), nil
}
