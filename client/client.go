package client

import (
	"crypto/tls"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	BaseURL    string
	Username   string
	Password   string
	Ticket     string
	CSRFToken  string
	HttpClient *http.Client
}

func NewClient(baseURL, username, password string) *Client {
	return &Client{
		BaseURL:  strings.TrimRight(baseURL, "/"),
		Username: username,
		Password: password,
		HttpClient: &http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //Ignore Certificates.
			},
		},
	}
}
