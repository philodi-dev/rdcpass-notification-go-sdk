package smsc

import (
	"context"

	"github.com/philodi-dev/rdcpass-notification-go-sdk/internal/api"
)

// QuickClient provides single-stage helpers: authenticate and execute in one call.
// These methods do not update the multi-stage session cache.
type QuickClient struct {
	client *Client
}

// SendSMS creates a session and sends an SMS in one step.
func (q *QuickClient) SendSMS(ctx context.Context, phone, content string) (*NotificationResponse, error) {
	token, err := q.client.ephemeralToken(ctx)
	if err != nil {
		return nil, err
	}

	var resp NotificationResponse
	if err := q.client.do(ctx, "POST", api.PathSMSSend, SendSMSRequest{Phone: phone, Content: content}, token, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendOTP creates a session and sends an OTP in one step.
func (q *QuickClient) SendOTP(ctx context.Context, phone string) (*NotificationResponse, error) {
	token, err := q.client.ephemeralToken(ctx)
	if err != nil {
		return nil, err
	}

	var resp NotificationResponse
	if err := q.client.do(ctx, "POST", api.PathOTPSend, SendOTPRequest{Phone: phone}, token, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// VerifyOTP creates a session and verifies an OTP in one step.
func (q *QuickClient) VerifyOTP(ctx context.Context, phone, code string) (*OtpVerifyResponse, error) {
	token, err := q.client.ephemeralToken(ctx)
	if err != nil {
		return nil, err
	}

	var resp OtpVerifyResponse
	if err := q.client.do(ctx, "POST", api.PathOTPVerify, VerifyOTPRequest{Phone: phone, Code: code}, token, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendEmail creates a session and sends a plain-text email in one step.
func (q *QuickClient) SendEmail(ctx context.Context, to, subject, body string) (*NotificationResponse, error) {
	token, err := q.client.ephemeralToken(ctx)
	if err != nil {
		return nil, err
	}

	var resp NotificationResponse
	req := SendEmailRequest{To: to, Subject: subject, Body: body}
	if err := q.client.do(ctx, "POST", api.PathEmailSend, req, token, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendHTMLEmail creates a session and sends an HTML email in one step.
func (q *QuickClient) SendHTMLEmail(ctx context.Context, to, subject, html string) (*NotificationResponse, error) {
	token, err := q.client.ephemeralToken(ctx)
	if err != nil {
		return nil, err
	}

	var resp NotificationResponse
	req := SendHTMLEmailRequest{To: to, Subject: subject, HTML: html}
	if err := q.client.do(ctx, "POST", api.PathEmailSendHTML, req, token, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendSMSSingle is a convenience wrapper around Quick().SendSMS.
func (c *Client) SendSMSSingle(ctx context.Context, phone, content string) (*NotificationResponse, error) {
	return c.Quick().SendSMS(ctx, phone, content)
}

// SendOTPSingle is a convenience wrapper around Quick().SendOTP.
func (c *Client) SendOTPSingle(ctx context.Context, phone string) (*NotificationResponse, error) {
	return c.Quick().SendOTP(ctx, phone)
}

// VerifyOTPSingle is a convenience wrapper around Quick().VerifyOTP.
func (c *Client) VerifyOTPSingle(ctx context.Context, phone, code string) (*OtpVerifyResponse, error) {
	return c.Quick().VerifyOTP(ctx, phone, code)
}

// SendEmailSingle is a convenience wrapper around Quick().SendEmail.
func (c *Client) SendEmailSingle(ctx context.Context, to, subject, body string) (*NotificationResponse, error) {
	return c.Quick().SendEmail(ctx, to, subject, body)
}

// SendHTMLEmailSingle is a convenience wrapper around Quick().SendHTMLEmail.
func (c *Client) SendHTMLEmailSingle(ctx context.Context, to, subject, html string) (*NotificationResponse, error) {
	return c.Quick().SendHTMLEmail(ctx, to, subject, html)
}
