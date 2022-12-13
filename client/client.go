package client

import "net/http"

// Config contains Kadena client configurable parameters.
type Config struct {
	NodeAddress string
	ChainName   string `help:"testnet01"`
	ChainID     int
	APIVersion  string `default:"0.0"`
}

// Client is implementation of Kadena API.
type Client struct {
	config Config

	http *http.Client
}

// New is constructor for Kadena API client.
func New(config Config) *Client {
	return &Client{
		config: config,
		http:   &http.Client{},
	}
}
