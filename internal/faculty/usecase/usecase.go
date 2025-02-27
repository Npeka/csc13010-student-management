package usecase

import (
	"context"

	"github.com/csc13010-student-management/internal/faculty"
	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/pkg/logger"
	"go.uber.org/zap"
)

type facultyUsecase struct {
	fr faculty.IFacultyRepository
	lg *logger.LoggerZap
}

func NewfacultyUsecase(
	fr faculty.IFacultyRepository,
	lg *logger.LoggerZap,
) faculty.IFacultyUsecase {
	return &facultyUsecase{
		fr: fr,
		lg: lg,
	}
}

func (fu *facultyUsecase) GetFaculties(ctx context.Context) ([]*models.Faculty, error) {
	faculties, err := fu.fr.GetFaculties(ctx)
	if err != nil {
		fu.lg.Error("Failed to get faculties", zap.Error(err))
		return nil, err
	}
	fu.lg.Info("Successfully fetched faculties")
	return faculties, nil
}

func (fu *facultyUsecase) CreateFaculty(ctx context.Context, faculty *models.Faculty) error {
	err := fu.fr.CreateFaculty(ctx, faculty)
	if err != nil {
		fu.lg.Error("Failed to create faculty", zap.Error(err))
		return err
	}
	fu.lg.Info("Successfully created faculty")
	return nil
}

func (fu *facultyUsecase) DeleteFaculty(ctx context.Context, id int) error {
	err := fu.fr.DeleteFaculty(ctx, id)
	if err != nil {
		fu.lg.Error("Failed to delete faculty", zap.Error(err))
		return err
	}
	fu.lg.Info("Successfully deleted faculty")
	return nil
}
