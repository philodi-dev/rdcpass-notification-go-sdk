package smsc

import (
	"context"
	"strings"

	"github.com/philodi-dev/rdcpass-notification-service/golang-sdk/internal/session"
	"github.com/philodi-dev/rdcpass-notification-service/golang-sdk/internal/transport"
)

// Client talks to the RDCPASS Notification Service API.
type Client struct {
	http      *transport.Client
	sessions  *session.Store
	appID     string
	secretKey string
}

// NewClient validates config and returns a ready-to-use client.
func NewClient(cfg Config) (*Client, error) {
	if err := cfg.validate(); err != nil {
		return nil, err
	}

	httpClient, err := transport.New(cfg.transportOptions())
	if err != nil {
		return nil, wrapError(err)
	}

	return &Client{
		http:      httpClient,
		sessions:  &session.Store{},
		appID:     strings.TrimSpace(cfg.AppID),
		secretKey: strings.TrimSpace(cfg.SecretKey),
	}, nil
}

// Quick returns single-stage helpers (session + action in one call).
func (c *Client) Quick() *QuickClient {
	return &QuickClient{client: c}
}

func (c *Client) do(ctx context.Context, method, path string, body any, bearer string, out any) error {
	return wrapError(c.http.Do(ctx, method, path, body, bearer, out))
}

func (c *Client) doAuthed(ctx context.Context, method, path string, body any, out any) error {
	token, err := c.bearerToken(ctx)
	if err != nil {
		return err
	}
	return c.do(ctx, method, path, body, token, out)
}

func (c *Client) bearerToken(ctx context.Context) (string, error) {
	if token, ok := c.sessions.Token(); ok {
		return token, nil
	}

	resp, err := c.exchangeSession(ctx)
	if err != nil {
		return "", err
	}

	c.sessions.Set(resp.SessionToken, resp.ExpiresIn)
	return resp.SessionToken, nil
}

func (c *Client) ephemeralToken(ctx context.Context) (string, error) {
	resp, err := c.exchangeSession(ctx)
	if err != nil {
		return "", err
	}
	return resp.SessionToken, nil
}
