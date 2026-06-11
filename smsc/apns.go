package smsc

import (
	"context"

	"github.com/philodi-dev/rdcpass-notification-service/golang-sdk/internal/api"
)

// APNS exposes Apple Push Notification endpoints.
type APNS struct {
	client *Client
}

// APNS returns the APNS API (multi-stage; reuses cached session).
func (c *Client) APNS() *APNS {
	return &APNS{client: c}
}

// Send delivers a push notification synchronously.
func (a *APNS) Send(ctx context.Context, req SendAPNSRequest) (*NotificationResponse, error) {
	var resp NotificationResponse
	if err := a.client.doAuthed(ctx, "POST", api.PathAPNSSend, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendAsync enqueues a push notification for background delivery.
func (a *APNS) SendAsync(ctx context.Context, req SendAPNSRequest) (*NotificationResponse, error) {
	var resp NotificationResponse
	if err := a.client.doAuthed(ctx, "POST", api.PathAPNSSendAsync, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendAPNS is a convenience wrapper around APNS().Send.
func (c *Client) SendAPNS(ctx context.Context, req SendAPNSRequest) (*NotificationResponse, error) {
	return c.APNS().Send(ctx, req)
}

// SendAPNSAsync is a convenience wrapper around APNS().SendAsync.
func (c *Client) SendAPNSAsync(ctx context.Context, req SendAPNSRequest) (*NotificationResponse, error) {
	return c.APNS().SendAsync(ctx, req)
}
