package usecase

import (
	"context"

	"github.com/csc13010-student-management/internal/config"
	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/opentracing/opentracing-go"
)

type configUsecase struct {
	cr config.IConfigRepository
	lg *logger.LoggerZap
}

func NewConfigUsecase(
	cr config.IConfigRepository,
	lg *logger.LoggerZap,
) config.IConfigUsecase {
	return &configUsecase{
		cr: cr,
		lg: lg,
	}
}

func (c *configUsecase) GetConfig(ctx context.Context) (*models.Config, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "configUsecase.GetConfig")
	defer span.Finish()

	cf, err := c.cr.GetConfig(ctx)
	if err != nil {
		return nil, err
	}

	return cf, nil
}

func (c *configUsecase) UpdateConfig(ctx context.Context, cf *models.Config) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "configUsecase.UpdateConfig")
	defer span.Finish()

	err := c.cr.UpdateConfig(ctx, cf)
	if err != nil {
		return err
	}

	return nil
}
