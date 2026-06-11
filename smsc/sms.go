package smsc

import (
	"context"

	"github.com/philodi-dev/rdcpass-notification-go-sdk/internal/api"
)

// SMS exposes SMS notification endpoints.
type SMS struct {
	client *Client
}

// SMS returns the SMS API (multi-stage; reuses cached session).
func (c *Client) SMS() *SMS {
	return &SMS{client: c}
}

// Send delivers an SMS synchronously.
func (s *SMS) Send(ctx context.Context, req SendSMSRequest) (*NotificationResponse, error) {
	var resp NotificationResponse
	if err := s.client.doAuthed(ctx, "POST", api.PathSMSSend, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendAsync enqueues an SMS for background delivery.
func (s *SMS) SendAsync(ctx context.Context, req SendSMSRequest) (*NotificationResponse, error) {
	var resp NotificationResponse
	if err := s.client.doAuthed(ctx, "POST", api.PathSMSSendAsync, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendSMS is a convenience wrapper around SMS().Send.
func (c *Client) SendSMS(ctx context.Context, req SendSMSRequest) (*NotificationResponse, error) {
	return c.SMS().Send(ctx, req)
}

// SendSMSAsync is a convenience wrapper around SMS().SendAsync.
func (c *Client) SendSMSAsync(ctx context.Context, req SendSMSRequest) (*NotificationResponse, error) {
	return c.SMS().SendAsync(ctx, req)
}
