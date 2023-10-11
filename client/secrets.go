package client

import (
	"encoding/json"
	"errors"
	"rabbitx-client/model"
)

// RefreshSecrets invalidates the current secrets and automatically increases the expiration time.
// This function requires a signature with api_secret for this request.
// The api_key also needs to be set up as a header.
// It returns a pointer to a Secret object and an error object.
// If the request fails, the function returns an error.
func (c *RbClient) RefreshSecrets(apiKey, apiSecret, refreshToken string) (*model.Secret, error) {
	// Setting up the headers with the API key.
	headers := map[string]string{
		API_KEY_HEADER: apiKey,
	}

	// Making a POST request to the PATH_SECRETS_REFRESH endpoint.
	// The request includes the refresh token and the headers.
	// The secret key is also included in the request.
	respBody, err := c.post(PATH_SECRETS_REFRESH, SecretRefreshRequest{
		RefreshToken: refreshToken,
	}, headers, &secretKey{apiKey: apiKey, apiSecret: apiSecret})

	// If there is an error in making the request, return the error.
	if err != nil {
		return nil, err
	}

	// Declaring a variable to hold the response.
	var resp Response[*model.Secret]

	// Unmarshalling the response body into the resp variable.
	// If there is an error in unmarshalling, return the error.
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, err
	}

	// If the length of the result in the response is less than or equal to 0,
	// panic with the message "NEVER happened for resp.Result".
	if len(resp.Result) <= 0 {
		panic("NEVER happened for resp.Result")
	}

	// If the request was not successful, return an error with the error message from the response.
	if !resp.Success {
		return nil, errors.New(resp.Error)
	}

	// If the request was successful, return the first result from the response and no error.
	return resp.Result[0], nil
}
