package smsc

import (
	"context"

	"github.com/philodi-dev/rdcpass-notification-go-sdk/internal/api"
)

// OTP exposes one-time password endpoints.
type OTP struct {
	client *Client
}

// OTP returns the OTP API (multi-stage; reuses cached session).
func (c *Client) OTP() *OTP {
	return &OTP{client: c}
}

// Send delivers an OTP via SMS synchronously.
func (o *OTP) Send(ctx context.Context, req SendOTPRequest) (*NotificationResponse, error) {
	var resp NotificationResponse
	if err := o.client.doAuthed(ctx, "POST", api.PathOTPSend, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendAsync enqueues OTP delivery for background processing.
func (o *OTP) SendAsync(ctx context.Context, req SendOTPRequest) (*NotificationResponse, error) {
	var resp NotificationResponse
	if err := o.client.doAuthed(ctx, "POST", api.PathOTPSendAsync, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Verify validates an OTP code synchronously.
func (o *OTP) Verify(ctx context.Context, req VerifyOTPRequest) (*OtpVerifyResponse, error) {
	var resp OtpVerifyResponse
	if err := o.client.doAuthed(ctx, "POST", api.PathOTPVerify, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// VerifyAsync enqueues OTP verification for background processing.
func (o *OTP) VerifyAsync(ctx context.Context, req VerifyOTPRequest) (*OtpVerifyResponse, error) {
	var resp OtpVerifyResponse
	if err := o.client.doAuthed(ctx, "POST", api.PathOTPVerifyAsync, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendOTP is a convenience wrapper around OTP().Send.
func (c *Client) SendOTP(ctx context.Context, req SendOTPRequest) (*NotificationResponse, error) {
	return c.OTP().Send(ctx, req)
}

// SendOTPAsync is a convenience wrapper around OTP().SendAsync.
func (c *Client) SendOTPAsync(ctx context.Context, req SendOTPRequest) (*NotificationResponse, error) {
	return c.OTP().SendAsync(ctx, req)
}

// VerifyOTP is a convenience wrapper around OTP().Verify.
func (c *Client) VerifyOTP(ctx context.Context, req VerifyOTPRequest) (*OtpVerifyResponse, error) {
	return c.OTP().Verify(ctx, req)
}

// VerifyOTPAsync is a convenience wrapper around OTP().VerifyAsync.
func (c *Client) VerifyOTPAsync(ctx context.Context, req VerifyOTPRequest) (*OtpVerifyResponse, error) {
	return c.OTP().VerifyAsync(ctx, req)
}
