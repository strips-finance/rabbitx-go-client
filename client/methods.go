package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"rabbitx-client/auth"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// Avoid acquiring any lock inside methods
// If a signature is required, the caller needs to acquire the lock
// Then, it can be passed down the call chain
type secretKey struct {
	apiKey    string
	apiSecret string
}

// setHeaders sets the headers for a given http request.
func (c *RbClient) setHeaders(req *http.Request, headers map[string]string) {
	req.Header.Set("Content-Type", "application/json")

	for title, value := range headers {
		req.Header.Set(title, value)
	}
}

// get sends a GET request to the specified path with the provided parameters and headers.
func (c *RbClient) get(path string, params map[string]string, headers map[string]string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.apiUrl, path)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	c.setHeaders(req, headers)

	q := req.URL.Query()
	for paramKey, paramValue := range params {
		q.Add(paramKey, paramValue)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

// post sends a POST request to the specified path with the provided body, headers, and secret key.
func (c *RbClient) post(path string, body interface{}, headers map[string]string, secret *secretKey) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.apiUrl, path)
	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	c.setHeaders(req, headers)

	return c.doRequest(req, secret)
}

// put sends a PUT request to the specified path with the provided body, headers, and secret key.
func (c *RbClient) put(path string, body interface{}, headers map[string]string, secret *secretKey) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.apiUrl, path)
	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	c.setHeaders(req, headers)

	return c.doRequest(req, secret)
}

// delete sends a DELETE request to the specified path with the provided body, headers, and secret key.
func (c *RbClient) delete(path string, body interface{}, headers map[string]string, secret *secretKey) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.apiUrl, path)
	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	c.setHeaders(req, headers)

	return c.doRequest(req, secret)
}

// doRequest sends the request and returns the response or any occurring error.
func (c *RbClient) doRequest(req *http.Request, secret *secretKey) ([]byte, error) {
	if secret != nil {
		_, err := c.setSignatureHeaders(req, secret)
		if err != nil {
			return nil, err
		}
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	logrus.
		WithField("Request URL: ", req.URL).
		WithField("Response status: ", resp.Status).
		WithField("Response code: ", resp.StatusCode).
		Warnf("data: %s", string(data))

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(resp.Status)
	}

	return data, err
}

// setSignatureHeaders sets the signature headers for a given http request.
func (c *RbClient) setSignatureHeaders(req *http.Request, secret *secretKey) (string, error) {

	payload, err := c.parsePayload(req)
	if err != nil {
		return "", err
	}

	if payload == nil {
		return "", nil
	}

	timestamp := time.Now().Unix() + SIGNATURE_LIFETIME

	signature, err := auth.PayloadSignature(payload, secret.apiSecret, timestamp)
	if err != nil {
		return "", err
	}

	c.setHeaders(req, map[string]string{
		API_SECRET_SIGNATURE_HEADER: signature,
		API_SECRET_TIMESTAMP_HEADER: strconv.FormatInt(timestamp, 10),
	})

	return signature, nil
}

// parsePayload parses the payload from a given http request.
func (c *RbClient) parsePayload(req *http.Request) (map[string]string, error) {
	var data map[string]json.RawMessage

	rMethod := req.Method
	if !(rMethod == http.MethodPost || rMethod == http.MethodPut || rMethod == http.MethodDelete) {
		return nil, nil
	}

	jsonData, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	req.Body = io.NopCloser(bytes.NewReader(jsonData))
	if err = json.Unmarshal(jsonData, &data); err != nil {
		return nil, err
	}

	payloadData := map[string]string{}
	for k, v := range data {
		payloadData[k] = strings.Trim(string(v), "\"")
	}

	payloadData["method"] = rMethod
	payloadData["path"] = req.URL.Path

	return payloadData, nil
}
