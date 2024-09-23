package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/cenkalti/backoff/v4"
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

func (c *Client) doRequest(method, endpoint string, body io.Reader) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", c.baseURL, endpoint)

	var resp *http.Response
	var err error

	operation := func() error {
		req, err := http.NewRequest(method, url, body)
		if err != nil {
			return fmt.Errorf("error creating request: %w", err)
		}

		req.Header.Set("User-Agent", c.userAgent)
		req.Header.Set("Accept", "application/json")

		c.logger.Info("Making request",
			zap.String("method", method),
			zap.String("url", url))

		resp, err := c.httpClient.Do(req)
		if err != nil {
			c.logger.Error("request failed",
				zap.String("method", method),
				zap.String("url", url),
				zap.Error(err))

			return fmt.Errorf("error making request: %w", err)
		}

		if resp.StatusCode >= 500 {
			return fmt.Errorf("server error: %s", resp.Status)
		}

		return nil
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = c.timeout

	err = backoff.Retry(operation, b)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

/*
GetPokemon retrieves information about a Pokemon by its ID or name.
It returns a Pokemon struct containing the data or an error if the request fails.
*/
func (c *Client) GetPokemon(idOrName string) (*Pokemon, error) {
	resp, err := c.doRequest("GET", fmt.Sprintf("pokemon/%s", idOrName), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var pokemon Pokemon
	if err := json.NewDecoder(resp.Body).Decode(&pokemon); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &pokemon, nil
}
