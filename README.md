# RDCPASS Notification Service — Go SDK

Go client for the [RDCPASS Notification Service](https://github.com/philodi-dev/rdcpass-notification-service) API (SMS, email, OTP, APNS).

App registration is **not** part of this SDK. Obtain `app_id` and `secret_key` from your administrator, then configure the client with your service `base_url`.

## Installation

```bash
go get github.com/philodi-dev/rdcpass-notification-go-sdk/smsc
```

For local development inside this repository:

```bash
go mod edit -replace github.com/philodi-dev/rdcpass-notification-go-sdk=../rdcpass-notification-go-sdk
```

## Project layout

```
rdcpass-notification-go-sdk/
├── smsc/                  # Public API (import this package)
│   ├── client.go          # Client entry point
│   ├── config.go          # Config struct
│   ├── options.go         # Functional options + New()
│   ├── auth.go            # Session management (multi-stage)
│   ├── quick.go           # Single-stage helpers
│   ├── sms.go / otp.go / email.go / apns.go / platform.go / requests.go
│   └── types.go / errors.go
├── internal/
│   ├── api/               # Route path constants
│   ├── transport/         # HTTP + TLS layer
│   └── session/           # Bearer token cache
└── examples/basic/        # Runnable demo
```

## Quick start

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/philodi-dev/rdcpass-notification-go-sdk/smsc"
)

func main() {
    client, err := smsc.New(
        "https://smsc.example.com",
        "your-app-id",
        "your-secret-key",
        smsc.WithTLSInsecureSkipVerify(), // dev/UAT only
    )
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()

    // Single-stage: session + action in one call
    resp, err := client.Quick().SendOTP(ctx, "+2434445079")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("OTP sent:", resp.RequestID, resp.Status)
}
```

## Configuration

| Field / Option | Required | Description |
|----------------|----------|-------------|
| `BaseURL` | yes | Service root URL (e.g. `https://smsc.example.com`) |
| `AppID` | yes | Application identifier |
| `SecretKey` | yes | Application secret |
| `WithHTTPClient()` | no | Custom `*http.Client` |
| `WithTimeout()` | no | Default HTTP timeout (30s) |
| `WithTLSInsecureSkipVerify()` | no | Skip TLS verification (dev/UAT only) |
| `WithTLSCACertFile()` / `WithTLSCACertPEM()` | no | PEM CA bundle for private certs |

```go
// Config struct (alternative to functional options)
client, err := smsc.NewClient(smsc.Config{
    BaseURL:   baseURL,
    AppID:     appID,
    SecretKey: secretKey,
})
```

## API surface

The SDK exposes two complementary styles:

### 1. Grouped services (recommended for multi-call flows)

| Accessor | Methods |
|----------|---------|
| `client.Auth()` | `CreateSession`, `ClearSession` |
| `client.Platform()` | `Health` |
| `client.SMS()` | `Send`, `SendAsync` |
| `client.Email()` | `Send`, `SendAsync`, `SendHTML`, `SendHTMLAsync` |
| `client.OTP()` | `Send`, `SendAsync`, `Verify`, `VerifyAsync` |
| `client.APNS()` | `Send`, `SendAsync` |
| `client.Requests()` | `GetStatus` |
| `client.Quick()` | `SendSMS`, `SendOTP`, `VerifyOTP`, `SendEmail`, `SendHTMLEmail` |

```go
session, _ := client.Auth().CreateSession(ctx)

sms, _ := client.SMS().Send(ctx, smsc.SendSMSRequest{
    Phone:   "+2434445079",
    Content: "Hello",
})

otp, _ := client.OTP().Send(ctx, smsc.SendOTPRequest{Phone: "+2434445079"})

email, _ := client.Email().Send(ctx, smsc.SendEmailRequest{
    To:      "user@example.com",
    Subject: "Hello",
    Body:    "Welcome to RDCPASS.",
})
```

Subsequent calls reuse the cached session automatically (refreshed 10s before expiry).

### 2. Quick (single-stage)

`client.Quick()` creates a **fresh session** per call and does not touch the multi-stage cache:

```go
client.Quick().SendSMS(ctx, phone, content)
client.Quick().SendOTP(ctx, phone)
client.Quick().VerifyOTP(ctx, phone, code)
client.Quick().SendEmail(ctx, to, subject, body)
client.Quick().SendHTMLEmail(ctx, to, subject, html)
```

### Convenience methods

Top-level shortcuts mirror the grouped API for brevity:

- `client.SendSMS`, `client.SendEmail`, `client.SendOTP`, `client.VerifyOTP`, …
- `client.SendSMSSingle`, `client.SendEmailSingle`, `client.SendOTPSingle`, `client.VerifyOTPSingle`

## Error handling

```go
resp, err := client.SMS().Send(ctx, req)
if err != nil {
    var apiErr *smsc.APIError
    if errors.As(err, &apiErr) {
        if apiErr.IsUnauthorized() { /* bad credentials */ }
        if apiErr.IsThrottled()     { /* rate limited */ }
    }
    log.Fatal(err)
}
```

## Example project

```bash
cd examples/basic
export SMSC_BASE_URL=https://your-host/smsc
export SMSC_APP_ID=your-app-id
export SMSC_SECRET_KEY=your-secret-key
export SMSC_PHONE=+243819989641
export SMSC_TLS_INSECURE=true   # UAT self-signed certs
./run.sh
```

**macOS 26 note:** Go 1.22 may fail with `dyld: missing LC_UUID`. Use `CGO_ENABLED=0 go run .`, upgrade to Go 1.24+, or build with `-ldflags="-linkmode=external"`.

## Types

- `NotificationResponse` — SMS / email / OTP send / APNS result
- `SendEmailRequest` / `SendHTMLEmailRequest` — plain-text and HTML email payloads
- `OtpVerifyResponse` — OTP verify result (`request_id`, `status`, `signature`)
- `SessionResponse` — session token and TTL
- `HealthResponse` — platform health

**Delivery status:** `pending`, `processing`, `delivered`, `failed`

**OTP verify status:** `verified`, `failed`, `pending`
