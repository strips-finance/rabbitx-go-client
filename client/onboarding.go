package client

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"rabbitx-client/auth"
	"rabbitx-client/model"
	"strconv"
	"time"
)

// Onboarding is a method that creates a new user on the exchange or returns new secrets if the user already exists.
// This method requires a private key for onboarding and the wallet string.
// The private key is used to generate a signature which is passed in the header RBT-PK-SIGNATURE of the request.
// The method returns a pointer to an OnboardMarketMakerResult object and an error object.
// If the private key is nil, the method returns an error.
// If the onboarding signature cannot be generated, the method returns an error.
// If the post request to the PATH_ONBOARDING endpoint fails, the method returns an error.
// If the response body cannot be unmarshalled into a Response object, the method returns an error.
// If the response indicates a failure, the method returns an error.
func (c *RbClient) Onboarding(wallet string, privateKey *ecdsa.PrivateKey) (*model.OnboardMarketMakerResult, error) {
	if privateKey == nil {
		return nil, fmt.Errorf("private key required for onboarding")
	}

	timestamp := time.Now().Unix() + SIGNATURE_LIFETIME

	signature, err := auth.OnboardingSiganture(privateKey, timestamp)
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		API_SECRET_SIGNATURE_HEADER: signature,
		API_SECRET_TIMESTAMP_HEADER: strconv.FormatInt(timestamp, 10),
	}

	respBody, err := c.post(PATH_ONBOARDING, OnboardingRequest{
		IsClient:  false,
		Wallet:    wallet,
		Signature: signature,
	}, headers, nil)
	if err != nil {
		return nil, err
	}

	var resp Response[*model.OnboardMarketMakerResult]

	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, err
	}

	if len(resp.Result) <= 0 {
		return nil, errors.New("unexpected empty result")
	}

	if !resp.Success {
		return nil, errors.New(resp.Error)
	}

	return resp.Result[0], nil
}
