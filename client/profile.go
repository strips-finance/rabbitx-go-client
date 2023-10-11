package client

import (
	"encoding/json"
	"errors"
	"rabbitx-client/model"
)

// GetProfile is a method that retrieves the profile data of the user.
// This method requires the API key which is obtained from the GetSecrets method.
// The API key is used to authenticate the request.
// The method returns a pointer to a ProfileData object and an error object.
// If the API key is not valid or the request fails, the method returns an error.
func (c *RbClient) GetProfile() (*model.ProfileData, error) {
	apiKey, _, _, err := c.GetSecrets()
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		API_KEY_HEADER: apiKey,
	}

	respBody, err := c.get(PATH_ACCOUNT, nil, headers)
	if err != nil {
		return nil, err
	}

	var resp Response[*model.ProfileData]

	err = json.Unmarshal(respBody, &resp)
	if err != nil {
		return nil, err
	}

	if len(resp.Result) <= 0 {
		return nil, errors.New("unexpected response: no result data")
	}

	if !resp.Success {
		return nil, errors.New(resp.Error)
	}

	return resp.Result[0], nil
}
