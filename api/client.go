package api

import (
	"net/http"
	"time"
)

var (
	// DefaultURL is the direct pokeapi URL
	DefaultURL = "https://pokeapi.co"
)

// Client struct which wraps up our pokemon client
type Client struct {
	httpClient *http.Client
	apiURL     string
}

// Option is used to custom initialize the client
type Option func(client *Client)

// NewClient is our Client initializer
func NewClient(options ...Option) *Client {
	client := &Client{
		httpClient: &http.Client{Timeout: 1 * time.Second},
		apiURL:     DefaultURL,
	}
	for _, option := range options {
		option(client)
	}
	return client
}

// CustomAPIURL is a function that allows custom other APIs
func CustomAPIURL(url string) Option {
	return func(client *Client) {
		client.apiURL = url
	}
}
