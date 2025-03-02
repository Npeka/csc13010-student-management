package usecase

import (
	"context"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/program"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/opentracing/opentracing-go"
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
	span, ctx := opentracing.StartSpanFromContext(ctx, "programUseCase.GetPrograms")
	defer span.Finish()

	programs, err := pu.pr.GetPrograms(ctx)
	if err != nil {
		return nil, err
	}
	return programs, nil
}

func (pu *programUsecase) CreateProgram(ctx context.Context, program *models.Program) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "programUseCase.CreateProgram")
	defer span.Finish()

	if err := pu.pr.CreateProgram(ctx, program); err != nil {
		return err
	}
	return nil
}

func (pu *programUsecase) UpdateProgram(ctx context.Context, program *models.Program) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "programUseCase.UpdateProgram")
	defer span.Finish()

	if err := pu.pr.UpdateProgram(ctx, program); err != nil {
		return err
	}
	return nil
}

func (pu *programUsecase) DeleteProgram(ctx context.Context, id uint) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "programUseCase.DeleteProgram")
	defer span.Finish()

	if err := pu.pr.DeleteProgram(ctx, id); err != nil {
		return err
	}
	return nil
}
