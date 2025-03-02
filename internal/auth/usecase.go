package auth

import (
	"context"

	"github.com/csc13010-student-management/internal/auth/dtos"
)

type IAuthUsecase interface {
	Register(ctx context.Context, registerReq *dtos.UserRegisterRequestDTO) error
	Login(ctx context.Context, loginReq *dtos.UserLoginRequestDTO) (*dtos.UserLoginResponseDTO, error)
	Logout(ctx context.Context)
	Refresh(ctx context.Context)
}
