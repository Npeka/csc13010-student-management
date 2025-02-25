package response

const (
	// Register errors
	ErrEmailInvalid = 20001
	ErrEmailExist   = 20002

	// Token and OTP errors
	ErrInvalidToken = 30001
	ErrInvalidOTP   = 30002
	ErrSendEmailOTP = 30003

	// Parameter errors
	ErrParamInvalid = 40001

	// User errors
	ErrUserExist = 50001

	// OTP errors
	ErrOTPNotExists = 60001

	// Common errors
	ErrUnauthorized   = 401
	ErrInternalServer = 500

	// Success code
	ErrCodeSuccess = 0
)

var msg = map[int]string{
	ErrEmailInvalid:   "Email is invalid",
	ErrEmailExist:     "Email is already exist",
	ErrInvalidToken:   "Invalid token",
	ErrInvalidOTP:     "Invalid OTP",
	ErrSendEmailOTP:   "Failed to send email OTP",
	ErrParamInvalid:   "Invalid parameter",
	ErrUnauthorized:   "Unauthorized",
	ErrInternalServer: "Internal server error",
	ErrCodeSuccess:    "Success",

	ErrUserExist: "User is already exist",

	ErrOTPNotExists: "OTP not exists",
}
