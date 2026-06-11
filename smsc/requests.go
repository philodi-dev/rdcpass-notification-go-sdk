package smsc

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/philodi-dev/rdcpass-notification-service/golang-sdk/internal/api"
)

// Requests exposes async request status polling.
type Requests struct {
	client *Client
}

// Requests returns the request status API.
func (c *Client) Requests() *Requests {
	return &Requests{client: c}
}

// GetStatus polls async or historical request status by ID.
func (r *Requests) GetStatus(ctx context.Context, requestID string) (*RequestStatus, error) {
	token, err := r.client.bearerToken(ctx)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(api.PathRequestStatusFmt, requestID)

	var raw json.RawMessage
	if err := r.client.do(ctx, "GET", path, nil, token, &raw); err != nil {
		return nil, err
	}

	status := &RequestStatus{}

	var notif NotificationResponse
	if err := json.Unmarshal(raw, &notif); err == nil && notif.RequestID != "" {
		status.Notification = &notif
		return status, nil
	}

	var verify OtpVerifyResponse
	if err := json.Unmarshal(raw, &verify); err == nil && verify.RequestID != "" {
		status.OtpVerify = &verify
		return status, nil
	}

	return nil, fmt.Errorf("smsc: unexpected request status payload")
}

// GetRequestStatus is a convenience wrapper around Requests().GetStatus.
func (c *Client) GetRequestStatus(ctx context.Context, requestID string) (*RequestStatus, error) {
	return c.Requests().GetStatus(ctx, requestID)
}
