package pokeapi

import (
	"time"

	"go.uber.org/zap"
)

// Option is a functional option type that allows us to configure the Client
type Option func(*Client)

// WithBaseURL sets the base URL for PokeAPI
func WithBaseURL(url string) Option {
	return func(c *Client) {
		c.baseURL = url
	}
}

// WithTimeout sets the HTTP client timeout
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.timeout = timeout
	}
}

// WithUserAgent sets the HTTP client user agent.
func WithUserAgent(userAgent string) Option {
	return func(c *Client) {
		c.userAgent = userAgent
	}
}

// WithRetryPolicy sets the retry policy for failed requests.
func WithRetryPolicy(attempts int, delay time.Duration) Option {
	return func(c *Client) {
		c.retryAttempts = attempts
		c.retryDelay = delay
	}
}

// WithLogger sets the logger for the client.
func WithLogger(logger *zap.Logger) Option {
	return func(c *Client) {
		c.logger = logger
	}
}
