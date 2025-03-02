package usecase

import (
	"context"

	"github.com/csc13010-student-management/internal/faculty"
	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/opentracing/opentracing-go"
)

type facultyUsecase struct {
	fr faculty.IFacultyRepository
	lg *logger.LoggerZap
}

func NewFacultyUsecase(
	fr faculty.IFacultyRepository,
	lg *logger.LoggerZap,
) faculty.IFacultyUsecase {
	return &facultyUsecase{
		fr: fr,
		lg: lg,
	}
}

func (fu *facultyUsecase) GetFaculties(ctx context.Context) ([]*models.Faculty, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "facultyUsecase.GetFaculties")
	defer span.Finish()

	faculties, err := fu.fr.GetFaculties(ctx)
	if err != nil {
		return nil, err
	}
	return faculties, nil
}

func (fu *facultyUsecase) CreateFaculty(ctx context.Context, faculty *models.Faculty) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "facultyUsecase.CreateFaculty")
	defer span.Finish()

	err := fu.fr.CreateFaculty(ctx, faculty)
	if err != nil {
		return err
	}
	return nil
}

func (fu *facultyUsecase) UpdateFaculty(ctx context.Context, faculty *models.Faculty) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "facultyUsecase.UpdateFaculty")
	defer span.Finish()

	err := fu.fr.UpdateFaculty(ctx, faculty)
	if err != nil {
		return err
	}
	return nil
}

func (fu *facultyUsecase) DeleteFaculty(ctx context.Context, id uint) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "facultyUsecase.DeleteFaculty")
	defer span.Finish()

	err := fu.fr.DeleteFaculty(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
