package events

type DebeziumEvent struct {
	Payload struct {
		Before map[string]interface{} `json:"before"`
		After  map[string]interface{} `json:"after"`
		Source struct {
			Table string `json:"table"`
			Lsn   int    `json:"lsn"`
		} `json:"source"`
		Transaction string `json:"transaction"`
		Op          string `json:"op"`
	} `json:"payload"`
}

type DebeziumAllEvent[T1 any, T2 any] struct {
	Payload struct {
		Before T1 `json:"before"`
		After  T2 `json:"after"`
	} `json:"payload"`
}

type DebeziumAfterEvent[T any] struct {
	Payload struct {
		Before map[string]interface{} `json:"before"`
		After  T                      `json:"after"`
	} `json:"payload"`
}

type RspwOTPEvent struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

type NotiType string

const (
	AuthUserCreated = "auth.user.created"
	AuthCreateUser  = "auth.user.create"
	NotiCreate      = "noti.create"

	NotiUserResetPasswordOTP NotiType = "auth.user.rspw.otp"
	NotiStudentStatusChanged NotiType = "student.status.changed"
)

type NotificationEvent struct {
	Type  NotiType               `json:"type"`
	Email string                 `json:"email"`
	Data  map[string]interface{} `json:"data"`
}
