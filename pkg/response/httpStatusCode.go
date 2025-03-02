package response

import "net/http"

const (
	// Token errors
	ErrMissingToken = 10001
	ErrInvalidToken = 10002

	// Permission errors
	ErrCheckPermission  = 10003
	ErrPermissionDenied = 10004

	// Register errors
	ErrEmailInvalid    = 20001
	ErrEmailExist      = 20002
	UserRegiterSuccess = 20003

	// Login errors
	UserLoginSuccess = 20004

	// Token and OTP errors
	ErrInvalidOTP   = 30002
	ErrSendEmailOTP = 30003

	// Parameter errors
	ErrParamInvalid = 40001

	// User errors
	ErrUserExist = 50001

	// OTP errors
	ErrOTPNotExists = 60001

	// Common errors
	ErrBadRequest     = 400
	ErrUnauthorized   = 401
	ErrInternalServer = 500

	// Success code
	ErrCodeSuccess = 0
)

var msg = map[int]struct {
	httpCode int
	message  string
}{
	ErrMissingToken: {httpCode: http.StatusUnauthorized, message: "Missing token"},
	ErrInvalidToken: {httpCode: http.StatusUnauthorized, message: "Invalid token"},

	ErrCheckPermission:  {httpCode: http.StatusInternalServerError, message: "Check permission error"},
	ErrPermissionDenied: {httpCode: http.StatusForbidden, message: "Permission denied"},

	ErrEmailInvalid:    {httpCode: 400, message: "Email is invalid"},
	ErrEmailExist:      {httpCode: 400, message: "Email is already exist"},
	UserRegiterSuccess: {httpCode: 200, message: "User register success"},

	UserLoginSuccess: {httpCode: 200, message: "User login success"},

	ErrInvalidOTP:     {httpCode: 400, message: "Invalid OTP"},
	ErrSendEmailOTP:   {httpCode: 500, message: "Failed to send email OTP"},
	ErrParamInvalid:   {httpCode: 400, message: "Invalid parameter"},
	ErrUnauthorized:   {httpCode: 401, message: "Unauthorized"},
	ErrInternalServer: {httpCode: 500, message: "Internal server error"},
	ErrCodeSuccess:    {httpCode: 200, message: "Success"},

	ErrUserExist: {httpCode: 400, message: "User is already exist"},

	ErrOTPNotExists: {httpCode: 400, message: "OTP not exists"},
}
