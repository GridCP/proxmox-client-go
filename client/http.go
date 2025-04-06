package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) Get(path string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, path)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Cookie", "PVEAuthCookie="+c.Ticket)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func (c *Client) Post(path string, payload any) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, path)
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "PVEAuthCookie="+c.Ticket)
	req.Header.Set("CSRFPreventionToken", c.CSRFToken)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
