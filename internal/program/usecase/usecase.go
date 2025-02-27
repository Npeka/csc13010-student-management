package usecase

import (
	"context"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/program"
	"github.com/csc13010-student-management/pkg/logger"
	"go.uber.org/zap"
)

type programUsecase struct {
	pr program.IProgramRepository
	lg *logger.LoggerZap
}

func NewProgramUsecase(
	pr program.IProgramRepository,
	lg *logger.LoggerZap,
) program.IProgramUsecase {
	return &programUsecase{
		pr: pr,
		lg: lg,
	}
}

func (pu *programUsecase) GetPrograms(ctx context.Context) ([]*models.Program, error) {
	programs, err := pu.pr.GetPrograms(ctx)
	if err != nil {
		pu.lg.Error("Failed to get programs", zap.Error(err))
		return nil, err
	}
	pu.lg.Info("Successfully fetched programs")
	return programs, nil
}

func (pu *programUsecase) CreateProgram(ctx context.Context, program *models.Program) error {
	if err := pu.pr.CreateProgram(ctx, program); err != nil {
		pu.lg.Error("Failed to create program", zap.Error(err))
		return err
	}
	pu.lg.Info("Successfully created program", zap.Int("id", program.ID))
	return nil
}

func (pu *programUsecase) DeleteProgram(ctx context.Context, id int) error {
	if err := pu.pr.DeleteProgram(ctx, id); err != nil {
		pu.lg.Error("Failed to delete program", zap.Error(err))
		return err
	}
	pu.lg.Info("Successfully deleted program", zap.Int("id", id))
	return nil
}
