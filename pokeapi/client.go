package pokeapi

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

// Client represents a PokeAPI client with configurable options.
type Client struct {
	baseURL       string
	httpClient    *http.Client
	timeout       time.Duration
	userAgent     string
	retryAttempts int
	retryDelay    time.Duration
	logger        *zap.Logger
}

/*
NewClient creates a new PokeAPI client with the provided options.
It uses default values for any option not explicitly set.
*/
func NewClient(options ...Option) *Client {
	client := &Client{
		baseURL:       "https://pokeapi.co/api/v2",
		httpClient:    &http.Client{},
		timeout:       30 * time.Second,
		userAgent:     "PokeAPI-Go-Client/1.0",
		retryAttempts: 3,
		retryDelay:    time.Second,
	}

	for _, opt := range options {
		opt(client)
	}

	if client.logger == nil {
		client.logger, _ = zap.NewProduction()
	}

	return client
}
