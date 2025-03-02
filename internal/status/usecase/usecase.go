package usecase

import (
	"context"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/status"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/opentracing/opentracing-go"
)

type statusUsecase struct {
	sr status.IStatusRepository
	lg *logger.LoggerZap
}

func NewStatusUsecase(
	sr status.IStatusRepository,
	lg *logger.LoggerZap,
) status.IStatusUsecase {
	return &statusUsecase{
		sr: sr,
		lg: lg,
	}
}

func (su *statusUsecase) GetStatuses(ctx context.Context) ([]*models.Status, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "statusUsecase.GetStatuses")
	defer span.Finish()

	statuses, err := su.sr.GetStatuses(ctx)
	if err != nil {
		return nil, err
	}
	return statuses, nil
}

func (su *statusUsecase) CreateStatus(ctx context.Context, status *models.Status) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "statusUsecase.CreateStatus")
	defer span.Finish()

	err := su.sr.CreateStatus(ctx, status)
	if err != nil {
		return err
	}
	return nil
}

func (su *statusUsecase) UpdateStatus(ctx context.Context, status *models.Status) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "statusUsecase.UpdateStatus")
	defer span.Finish()

	err := su.sr.UpdateStatus(ctx, status)
	if err != nil {
		return err
	}
	return nil
}

func (su *statusUsecase) DeleteStatus(ctx context.Context, id uint) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "statusUsecase.DeleteStatus")
	defer span.Finish()

	err := su.sr.DeleteStatus(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
