# Basic example

Runnable demo for the RDCPASS Notification Service Go SDK.

## Structure

```
examples/basic/
├── main.go      # Demo flows (health, quick OTP, multi-stage SMS)
├── config.go    # Environment-based client setup
├── run.sh       # macOS Tahoe-compatible launcher
└── go.mod
```

## Prerequisites

- Running notification service (local or remote)
- Valid `app_id` and `secret_key` (created via admin API, not the SDK)

## Run

```bash
cd golang-sdk/examples/basic

export SMSC_BASE_URL=https://100.51.101.5/smsc
export SMSC_APP_ID=your-app-id
export SMSC_SECRET_KEY=your-secret-key
export SMSC_PHONE=+243819989641
export SMSC_TLS_INSECURE=true

# Optional: verify an OTP after sending
# export SMSC_OTP_CODE=12345678

./run.sh
```

## What it demonstrates

| Step | API style | Call |
|------|-----------|------|
| Health check | `Platform()` | `client.Platform().Health(ctx)` |
| Send OTP | `Quick()` (single-stage) | `client.Quick().SendOTP(ctx, phone)` |
| Send SMS | `Auth()` + `SMS()` (multi-stage) | `client.Auth().CreateSession` → `client.SMS().Send` |
| Verify OTP | `Quick()` (optional) | `client.Quick().VerifyOTP(ctx, phone, code)` |

### macOS 26 (Tahoe) + Go 1.22

If `go run .` crashes with `dyld: missing LC_UUID`, use `./run.sh` or `CGO_ENABLED=0 go run .`.
