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
