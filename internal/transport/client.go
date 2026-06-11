package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const defaultTimeout = 30 * time.Second

// HTTPError is returned for non-2xx API responses.
type HTTPError struct {
	StatusCode int
	Message    string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("http error (status %d): %s", e.StatusCode, e.Message)
}

// Options configures the HTTP transport layer.
type Options struct {
	BaseURL    string
	HTTPClient *http.Client
	Timeout    time.Duration
	TLS        TLSOptions
}

// Client performs JSON HTTP requests against the notification service.
type Client struct {
	baseURL string
	http    *http.Client
}

// New validates options and returns an HTTP client.
func New(opts Options) (*Client, error) {
	baseURL := strings.TrimRight(strings.TrimSpace(opts.BaseURL), "/")
	if baseURL == "" {
		return nil, fmt.Errorf("transport: BaseURL is required")
	}

	httpClient := opts.HTTPClient
	if httpClient == nil {
		timeout := opts.Timeout
		if timeout == 0 {
			timeout = defaultTimeout
		}

		transport := http.DefaultTransport.(*http.Transport).Clone()
		tlsCfg, err := buildTLSConfig(opts.TLS)
		if err != nil {
			return nil, err
		}
		if tlsCfg != nil {
			transport.TLSClientConfig = tlsCfg
		}

		httpClient = &http.Client{
			Timeout:   timeout,
			Transport: transport,
		}
	}

	return &Client{baseURL: baseURL, http: httpClient}, nil
}

// Do sends an HTTP request and decodes a JSON response into out when provided.
func (c *Client) Do(ctx context.Context, method, path string, body any, bearerToken string, out any) error {
	req, err := c.newRequest(ctx, method, path, body, bearerToken)
	if err != nil {
		return err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("transport: request failed: %w", err)
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("transport: read body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		msg := strings.TrimSpace(string(raw))
		if msg == "" {
			msg = resp.Status
		}
		return &HTTPError{StatusCode: resp.StatusCode, Message: msg}
	}

	if out != nil && len(raw) > 0 {
		if err := json.Unmarshal(raw, out); err != nil {
			return fmt.Errorf("transport: decode response: %w", err)
		}
	}

	return nil
}

func (c *Client) newRequest(ctx context.Context, method, path string, body any, bearerToken string) (*http.Request, error) {
	var reader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("transport: encode request: %w", err)
		}
		reader = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, reader)
	if err != nil {
		return nil, fmt.Errorf("transport: build request: %w", err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")

	if bearerToken != "" {
		req.Header.Set("Authorization", "Bearer "+bearerToken)
	}

	return req, nil
}
