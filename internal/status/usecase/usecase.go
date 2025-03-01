package usecase

import (
	"context"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/status"
	"github.com/csc13010-student-management/pkg/logger"
	"go.uber.org/zap"
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
	statuses, err := su.sr.GetStatuses(ctx)
	if err != nil {
		su.lg.Error("Failed to get statuses", zap.Error(err))
		return nil, err
	}
	su.lg.Info("Successfully fetched statuses")
	return statuses, nil
}

func (su *statusUsecase) CreateStatus(ctx context.Context, status *models.Status) error {
	err := su.sr.CreateStatus(ctx, status)
	if err != nil {
		su.lg.Error("Failed to create status", zap.Error(err))
		return err
	}
	su.lg.Info("Successfully created status")
	return nil
}

func (su *statusUsecase) DeleteStatus(ctx context.Context, id uint) error {
	err := su.sr.DeleteStatus(ctx, id)
	if err != nil {
		su.lg.Error("Failed to delete status", zap.Error(err))
		return err
	}
	su.lg.Info("Successfully deleted status")
	return nil
}
