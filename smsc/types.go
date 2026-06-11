package smsc

// DeliveryStatus represents notification delivery state.
type DeliveryStatus string

const (
	StatusPending    DeliveryStatus = "pending"
	StatusProcessing DeliveryStatus = "processing"
	StatusDelivered  DeliveryStatus = "delivered"
	StatusFailed     DeliveryStatus = "failed"
)

// OtpVerifyStatus represents OTP verification outcome.
type OtpVerifyStatus string

const (
	OtpVerified OtpVerifyStatus = "verified"
	OtpFailed   OtpVerifyStatus = "failed"
	OtpPending  OtpVerifyStatus = "pending"
)

// NotificationType identifies the kind of notification.
type NotificationType string

const (
	TypeSMS       NotificationType = "sms"
	TypeOTPSend   NotificationType = "otp_send"
	TypeOTPVerify NotificationType = "otp_verify"
	TypeAPNS      NotificationType = "apns"
)

// SessionResponse is returned by POST /api/auth/session.
type SessionResponse struct {
	SessionToken string `json:"session_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
}

// NotificationResponse is returned by SMS, OTP send, and APNS endpoints.
type NotificationResponse struct {
	RequestID         string           `json:"request_id"`
	AppID             string           `json:"app_id"`
	NotificationType  NotificationType `json:"notification_type"`
	Status            DeliveryStatus   `json:"status"`
	CreatedAt         string           `json:"created_at"`
	CompletedAt       *string          `json:"completed_at"`
	Time              string           `json:"time"`
	ClientIP          *string          `json:"client_ip"`
	SenderID          *string          `json:"sender_id"`
	ProviderMessageID *string          `json:"provider_message_id"`
	ErrorMessage      *string          `json:"error_message"`
}

// OtpVerifyResponse is returned by OTP verify endpoints.
type OtpVerifyResponse struct {
	RequestID string          `json:"request_id"`
	Status    OtpVerifyStatus `json:"status"`
	Signature string          `json:"signature"`
}

// HealthResponse is returned by GET /api/platform/health.
type HealthResponse struct {
	Status      string `json:"status"`
	Description string `json:"description"`
	Timestamp   string `json:"timestamp"`
}

// SendSMSRequest is the body for POST /api/sms/send.
type SendSMSRequest struct {
	Phone   string `json:"phone"`
	Content string `json:"content"`
}

// SendOTPRequest is the body for POST /api/otp/send.
type SendOTPRequest struct {
	Phone string `json:"phone"`
}

// VerifyOTPRequest is the body for POST /api/otp/verify.
type VerifyOTPRequest struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

// SendAPNSRequest is the body for POST /api/apns/send.
type SendAPNSRequest struct {
	DeviceToken string                 `json:"device_token"`
	Title       string                 `json:"title"`
	Body        string                 `json:"body"`
	Badge       *uint32                `json:"badge,omitempty"`
	UserInfo    map[string]interface{} `json:"user_info,omitempty"`
}

// RequestStatus is returned by GET /api/requests/{id}.
type RequestStatus struct {
	Notification *NotificationResponse
	OtpVerify    *OtpVerifyResponse
}
