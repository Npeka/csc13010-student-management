package dtos

// UserRegisterRequestDTO represents the request body for user registration.
type UserRegisterRequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
