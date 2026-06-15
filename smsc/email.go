package smsc

import (
	"context"

	"github.com/philodi-dev/rdcpass-notification-go-sdk/internal/api"
)

// Email exposes email notification endpoints.
type Email struct {
	client *Client
}

// Email returns the email API (multi-stage; reuses cached session).
func (c *Client) Email() *Email {
	return &Email{client: c}
}

// Send delivers a plain-text email synchronously.
func (e *Email) Send(ctx context.Context, req SendEmailRequest) (*NotificationResponse, error) {
	var resp NotificationResponse
	if err := e.client.doAuthed(ctx, "POST", api.PathEmailSend, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendAsync enqueues a plain-text email for background delivery.
func (e *Email) SendAsync(ctx context.Context, req SendEmailRequest) (*NotificationResponse, error) {
	var resp NotificationResponse
	if err := e.client.doAuthed(ctx, "POST", api.PathEmailSendAsync, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendHTML delivers an HTML email synchronously.
func (e *Email) SendHTML(ctx context.Context, req SendHTMLEmailRequest) (*NotificationResponse, error) {
	var resp NotificationResponse
	if err := e.client.doAuthed(ctx, "POST", api.PathEmailSendHTML, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendHTMLAsync enqueues an HTML email for background delivery.
func (e *Email) SendHTMLAsync(ctx context.Context, req SendHTMLEmailRequest) (*NotificationResponse, error) {
	var resp NotificationResponse
	if err := e.client.doAuthed(ctx, "POST", api.PathEmailSendHTMLAsync, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendEmail is a convenience wrapper around Email().Send.
func (c *Client) SendEmail(ctx context.Context, req SendEmailRequest) (*NotificationResponse, error) {
	return c.Email().Send(ctx, req)
}

// SendEmailAsync is a convenience wrapper around Email().SendAsync.
func (c *Client) SendEmailAsync(ctx context.Context, req SendEmailRequest) (*NotificationResponse, error) {
	return c.Email().SendAsync(ctx, req)
}

// SendHTMLEmail is a convenience wrapper around Email().SendHTML.
func (c *Client) SendHTMLEmail(ctx context.Context, req SendHTMLEmailRequest) (*NotificationResponse, error) {
	return c.Email().SendHTML(ctx, req)
}

// SendHTMLEmailAsync is a convenience wrapper around Email().SendHTMLAsync.
func (c *Client) SendHTMLEmailAsync(ctx context.Context, req SendHTMLEmailRequest) (*NotificationResponse, error) {
	return c.Email().SendHTMLAsync(ctx, req)
}
