package api

// API route paths (relative to BaseURL).
const (
	PathAuthSession      = "/api/auth/session"
	PathPlatformHealth   = "/api/platform/health"
	PathSMSSend          = "/api/sms/send"
	PathSMSSendAsync     = "/api/sms/send/async"
	PathOTPSend          = "/api/otp/send"
	PathOTPSendAsync     = "/api/otp/send/async"
	PathOTPVerify        = "/api/otp/verify"
	PathOTPVerifyAsync   = "/api/otp/verify/async"
	PathAPNSSend         = "/api/apns/send"
	PathAPNSSendAsync    = "/api/apns/send/async"
	PathEmailSend        = "/api/email/send"
	PathEmailSendAsync   = "/api/email/send/async"
	PathEmailSendHTML    = "/api/email/send/html"
	PathEmailSendHTMLAsync = "/api/email/send/html/async"
	PathRequestStatusFmt = "/api/requests/%s"
)
