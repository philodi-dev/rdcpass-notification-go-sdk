package smsc

import (
	"context"

	"github.com/philodi-dev/rdcpass-notification-service/golang-sdk/internal/api"
)

// Platform exposes unauthenticated platform endpoints.
type Platform struct {
	client *Client
}

// Platform returns the platform API.
func (c *Client) Platform() *Platform {
	return &Platform{client: c}
}

// Health checks service availability (no authentication required).
func (p *Platform) Health(ctx context.Context) (*HealthResponse, error) {
	var resp HealthResponse
	if err := p.client.do(ctx, "GET", api.PathPlatformHealth, nil, "", &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Health is a convenience wrapper around Platform().Health.
func (c *Client) Health(ctx context.Context) (*HealthResponse, error) {
	return c.Platform().Health(ctx)
}
