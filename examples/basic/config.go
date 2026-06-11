package main

import (
	"os"
	"strings"

	"github.com/philodi-dev/rdcpass-notification-go-sdk/smsc"
)

// Settings holds example runtime configuration from environment variables.
type Settings struct {
	BaseURL   string
	AppID     string
	SecretKey string
	Phone     string
	OTPCode   string
}

// LoadSettings reads configuration from the environment.
func LoadSettings() Settings {
	return Settings{
		BaseURL:   envOr("SMSC_BASE_URL", "http://localhost:3000"),
		AppID:     os.Getenv("SMSC_APP_ID"),
		SecretKey: os.Getenv("SMSC_SECRET_KEY"),
		Phone:     envOr("SMSC_PHONE", "+2434445079"),
		OTPCode:   os.Getenv("SMSC_OTP_CODE"),
	}
}

// NewClient builds an SDK client from settings.
func (s Settings) NewClient() (*smsc.Client, error) {
	cfg := smsc.Config{
		BaseURL:   s.BaseURL,
		AppID:     s.AppID,
		SecretKey: s.SecretKey,
	}

	if envBool("SMSC_TLS_INSECURE") {
		cfg.TLSInsecureSkipVerify = true
	}
	if caFile := os.Getenv("SMSC_TLS_CA_FILE"); caFile != "" {
		cfg.TLSCACertFile = caFile
	}

	return smsc.NewClient(cfg)
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func envBool(key string) bool {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(key))) {
	case "1", "true", "yes", "on":
		return true
	default:
		return false
	}
}
