package smsc

import (
	"net/http"
	"time"
)

// Option customises client configuration.
type Option func(*Config)

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(client *http.Client) Option {
	return func(cfg *Config) {
		cfg.HTTPClient = client
	}
}

// WithTimeout sets the default HTTP timeout when no custom client is provided.
func WithTimeout(timeout time.Duration) Option {
	return func(cfg *Config) {
		cfg.Timeout = timeout
	}
}

// WithTLSInsecureSkipVerify skips TLS verification (dev/UAT only).
func WithTLSInsecureSkipVerify() Option {
	return func(cfg *Config) {
		cfg.TLSInsecureSkipVerify = true
	}
}

// WithTLSCACertPEM sets a PEM-encoded CA bundle for TLS verification.
func WithTLSCACertPEM(pem string) Option {
	return func(cfg *Config) {
		cfg.TLSCACertPEM = pem
	}
}

// WithTLSCACertFile sets a path to a PEM CA bundle for TLS verification.
func WithTLSCACertFile(path string) Option {
	return func(cfg *Config) {
		cfg.TLSCACertFile = path
	}
}

// New builds a client from credentials and optional functional options.
func New(baseURL, appID, secretKey string, opts ...Option) (*Client, error) {
	cfg := Config{
		BaseURL:   baseURL,
		AppID:     appID,
		SecretKey: secretKey,
	}
	for _, opt := range opts {
		opt(&cfg)
	}
	return NewClient(cfg)
}
