package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type HTCL struct {
	Client *http.Client
}

func NewHttpClient(transport http.RoundTripper, timeout time.Duration) *HTCL {
	client := &http.Client{
		Transport: transport,
		Timeout:   timeout,
	}
	return &HTCL{
		Client: client,
	}
}

func REQ(
	ct *HTCL,
	method, url string,
	data interface{},

) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data: %w", err)
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := ct.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("POST request failed: %w", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		err := fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(responseBody))
		return responseBody, err
	}

	return responseBody, nil
}

func (hc *HTCL) POST(url string, body interface{}, onRes func([]byte)) error {
	res, err := REQ(hc, "POST", url, body)
	onRes(res)
	return err
}

func (hc *HTCL) GET(url string, body interface{}, onRes func([]byte)) error {
	res, err := REQ(hc, "GET", url, body)
	onRes(res)
	return err
}
