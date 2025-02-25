package usecase

import (
	"context"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/student"
	"github.com/csc13010-student-management/pkg/logger"
)

type studentUsecase struct {
	sr student.IStudentRepository
	lg *logger.LoggerZap
}

func NewStudentUsecase(
	sr student.IStudentRepository,
	lg *logger.LoggerZap,
) student.IStudentUsecase {
	return &studentUsecase{
		sr: sr,
		lg: lg,
	}
}

func (s *studentUsecase) GetStudents(ctx context.Context) ([]*models.Student, error) {
	students, err := s.sr.GetStudents(ctx)
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (s *studentUsecase) CreateStudent(ctx context.Context, student *models.Student) error {
	err := s.sr.CreateStudent(ctx, student)
	if err != nil {
		return err
	}
	return nil
}

func (s *studentUsecase) UpdateStudent(ctx context.Context, student *models.Student) error {
	err := s.sr.UpdateStudent(ctx, student)
	if err != nil {
		return err
	}
	return nil
}

func (s *studentUsecase) DeleteStudent(ctx context.Context, student_id string) error {
	err := s.sr.DeleteStudent(ctx, student_id)
	if err != nil {
		return err
	}
	return nil
}

func (s *studentUsecase) SearchStudent(ctx context.Context, query string) ([]*models.Student, error) {
	students, err := s.sr.SearchStudents(ctx, query)
	if err != nil {
		return nil, err
	}
	return students, nil
}

// SearchStudents implements student.IStudentUsecase.
func (s *studentUsecase) SearchStudents(ctx context.Context, query string) ([]*models.Student, error) {
	panic("unimplemented")
}
