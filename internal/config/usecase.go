package config

import (
	"context"

	"github.com/csc13010-student-management/internal/models"
)

type IConfigUsecase interface {
	GetConfig(ctx context.Context) (*models.Config, error)
	UpdateConfig(ctx context.Context, cf *models.Config) error
}
