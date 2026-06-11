// Package smsc is the official Go client for the RDCPASS Notification Service.
//
// Configure the client with a base URL, app ID, and secret key. App registration
// is performed outside this SDK (admin API).
//
// Two usage styles are supported:
//
//   - Multi-stage: create or reuse a cached session, then call service methods.
//   - Quick (single-stage): one-shot helpers that authenticate and act in one call.
//
// Example:
//
//	client, err := smsc.NewClient(smsc.Config{
//	    BaseURL:   "https://smsc.example.com",
//	    AppID:     "my-app",
//	    SecretKey: "secret",
//	})
//	resp, err := client.Quick().SendOTP(ctx, "+2434445079")
package smsc
