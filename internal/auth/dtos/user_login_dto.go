package dtos

type UserLoginRequestDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginResponseDTO struct {
	Token string `json:"token"`
}
