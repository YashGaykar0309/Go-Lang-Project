package client

import (
	"net/http"
	"time"
)

type Client struct {
	Endpoint   string
	HTTPClient *http.Client
}

func New(endpoint string) *Client {
	return &Client{
		Endpoint: endpoint,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
