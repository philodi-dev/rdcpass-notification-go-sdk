package smsc

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/philodi-dev/rdcpass-notification-service/golang-sdk/internal/transport"
)

// APIError is returned when the notification service responds with a non-success status.
type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("smsc: api error (status %d): %s", e.StatusCode, e.Message)
}

// IsUnauthorized reports whether the error is a 401 response.
func (e *APIError) IsUnauthorized() bool {
	return e.StatusCode == http.StatusUnauthorized
}

// IsNotFound reports whether the error is a 404 response.
func (e *APIError) IsNotFound() bool {
	return e.StatusCode == http.StatusNotFound
}

// IsThrottled reports whether the error is a 429 response.
func (e *APIError) IsThrottled() bool {
	return e.StatusCode == http.StatusTooManyRequests
}

func wrapError(err error) error {
	if err == nil {
		return nil
	}

	var httpErr *transport.HTTPError
	if errors.As(err, &httpErr) {
		return &APIError{
			StatusCode: httpErr.StatusCode,
			Message:    httpErr.Message,
		}
	}

	return fmt.Errorf("smsc: %w", err)
}
