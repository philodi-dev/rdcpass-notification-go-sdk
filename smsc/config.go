package smsc

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/philodi-dev/rdcpass-notification-service/golang-sdk/internal/transport"
)

// Config holds client credentials and transport options.
type Config struct {
	BaseURL   string
	AppID     string
	SecretKey string

	HTTPClient *http.Client
	Timeout    time.Duration

	// TLSInsecureSkipVerify disables TLS certificate verification (dev/UAT only).
	TLSInsecureSkipVerify bool
	TLSCACertPEM          string
	TLSCACertFile         string
}

func (cfg Config) validate() error {
	if strings.TrimSpace(cfg.BaseURL) == "" {
		return fmt.Errorf("smsc: BaseURL is required")
	}
	if strings.TrimSpace(cfg.AppID) == "" {
		return fmt.Errorf("smsc: AppID is required")
	}
	if strings.TrimSpace(cfg.SecretKey) == "" {
		return fmt.Errorf("smsc: SecretKey is required")
	}
	return nil
}

func (cfg Config) transportOptions() transport.Options {
	return transport.Options{
		BaseURL:    cfg.BaseURL,
		HTTPClient: cfg.HTTPClient,
		Timeout:    cfg.Timeout,
		TLS: transport.TLSOptions{
			InsecureSkipVerify: cfg.TLSInsecureSkipVerify,
			CACertPEM:          cfg.TLSCACertPEM,
			CACertFile:         cfg.TLSCACertFile,
		},
	}
}
