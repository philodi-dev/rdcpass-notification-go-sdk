package smsc

import (
	"context"

	"github.com/philodi-dev/rdcpass-notification-service/golang-sdk/internal/api"
)

// Auth exposes session management (multi-stage authentication).
type Auth struct {
	client *Client
}

// Auth returns the session API for multi-stage usage.
func (c *Client) Auth() *Auth {
	return &Auth{client: c}
}

// CreateSession exchanges app_id + secret_key for a bearer token and caches it.
func (a *Auth) CreateSession(ctx context.Context) (*SessionResponse, error) {
	resp, err := a.client.exchangeSession(ctx)
	if err != nil {
		return nil, err
	}
	a.client.sessions.Set(resp.SessionToken, resp.ExpiresIn)
	return resp, nil
}

// ClearSession removes the cached bearer token.
func (a *Auth) ClearSession() {
	a.client.sessions.Clear()
}

// CreateSession is a convenience wrapper around Auth().CreateSession.
func (c *Client) CreateSession(ctx context.Context) (*SessionResponse, error) {
	return c.Auth().CreateSession(ctx)
}

// ClearSession is a convenience wrapper around Auth().ClearSession.
func (c *Client) ClearSession() {
	c.Auth().ClearSession()
}

// exchangeSession uses the auth path constant.
func (c *Client) exchangeSession(ctx context.Context) (*SessionResponse, error) {
	body := map[string]string{
		"app_id":     c.appID,
		"secret_key": c.secretKey,
	}

	var resp SessionResponse
	if err := c.do(ctx, "POST", api.PathAuthSession, body, "", &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
