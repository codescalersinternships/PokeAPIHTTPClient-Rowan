package api

import "net/http"

// Client struct which wraps up our pokemon client
type Client struct {
	httpClient *http.Client
}

// NewClient is our Client initializer
func NewClient() *Client {
	client := &Client{
		httpClient: http.DefaultClient,
	}
	return client
}
