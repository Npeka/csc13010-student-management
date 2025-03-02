package auth

import (
	"context"

	"github.com/csc13010-student-management/internal/models"
)

type IAuthRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	FindByUsername(ctx context.Context, username string) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindRoleByName(ctx context.Context, name string) (*models.Role, error)
}
