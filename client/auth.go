package client

import (
	"encoding/json"
	"fmt"
	"gridcp-proxmox-go/types"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func (c *Client) Login() error {
	loginURL := fmt.Sprintf("%s/api2/json/access/ticket", c.BaseURL)
	data := url.Values{}
	data.Set("username", c.Username)
	data.Set("password", c.Password)

	req, _ := http.NewRequest("POST", loginURL, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("login failed: %s", body)
	}

	var authResp types.AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return err
	}

	c.Ticket = authResp.Data.Ticket
	c.CSRFToken = authResp.Data.CSRFPreventionToken

	return nil
}
