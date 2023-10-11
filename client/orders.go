package client

import (
	"encoding/json"
	"errors"
	"rabbitx-client/model"
	"reflect"
)

// CreateOrder is a method that creates a new order on the exchange.
// This method requires an OrderCreateRequest object as input.
// The method returns a pointer to an OrderCreateResponse object and an error object.
func (c *RbClient) CreateOrder(data *OrderCreateRequest) (*OrderCreateResponse, error) {
	apiKey, apiSecret, _, err := c.GetSecrets()
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		API_KEY_HEADER: apiKey,
	}

	respBody, err := c.post(PATH_ORDERS, data, headers, &secretKey{
		apiKey:    apiKey,
		apiSecret: apiSecret,
	})
	if err != nil {
		return nil, err
	}

	var resp Response[*OrderCreateResponse]

	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, errors.New(resp.Error)
	}

	return resp.Result[0], nil
}

// CancelOrder is a method that cancels an existing order on the exchange.
// This method requires an OrderCancelRequest object as input.
// The method returns a pointer to an OrderCancelResponse object and an error object.
func (c *RbClient) CancelOrder(data *OrderCancelRequest) (*OrderCancelResponse, error) {
	apiKey, apiSecret, _, err := c.GetSecrets()
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		API_KEY_HEADER: apiKey,
	}

	respBody, err := c.delete(PATH_ORDERS, data, headers, &secretKey{
		apiKey:    apiKey,
		apiSecret: apiSecret,
	})
	if err != nil {
		return nil, err
	}

	var resp Response[*OrderCancelResponse]

	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, errors.New(resp.Error)
	}

	return resp.Result[0], nil
}

// ListOrders is a method that lists all orders on the exchange.
// This method requires an OrderListRequest object as input.
// The method returns a slice of OrderData objects and an error object.
func (c *RbClient) ListOrders(data *OrderListRequest) ([]model.OrderData, error) {
	apiKey, _, _, err := c.GetSecrets()
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		API_KEY_HEADER: apiKey,
	}

	queryParams := makeQueryParams(reflect.ValueOf(*data))

	respBody, err := c.get(PATH_ORDERS, queryParams, headers)
	if err != nil {
		return nil, err
	}

	var resp Response[model.OrderData]

	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, errors.New(resp.Error)
	}

	return resp.Result, nil
}
