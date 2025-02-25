package usecase

import (
	"context"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/student"
	"github.com/csc13010-student-management/internal/student/dtos"
	"github.com/csc13010-student-management/pkg/logger"
	"go.uber.org/zap"
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
		s.lg.Error("Failed to get students", zap.Error(err))
		return nil, err
	}
	s.lg.Info("Successfully fetched students")
	return students, nil
}

func (s *studentUsecase) CreateStudent(ctx context.Context, student *models.Student) error {
	err := s.sr.CreateStudent(ctx, student)
	if err != nil {
		s.lg.Error("Failed to create student", zap.Error(err))
		return err
	}
	s.lg.Info("Successfully created student", zap.Int("id", student.ID))
	return nil
}

func (s *studentUsecase) UpdateStudent(ctx context.Context, student *models.Student) error {
	err := s.sr.UpdateStudent(ctx, student)
	if err != nil {
		s.lg.Error("Failed to update student", zap.Error(err))
		return err
	}
	s.lg.Info("Successfully updated student", zap.Int("id", student.ID))
	return nil
}

func (s *studentUsecase) DeleteStudent(ctx context.Context, student_id string) error {
	err := s.sr.DeleteStudent(ctx, student_id)
	if err != nil {
		s.lg.Error("Failed to delete student", zap.Error(err))
		return err
	}
	s.lg.Info("Successfully deleted student", zap.String("id", student_id))
	return nil
}

func (s *studentUsecase) GetOptions(ctx context.Context) (*dtos.OptionDTO, error) {
	options, err := s.sr.GetOptions(ctx)
	if err != nil {
		s.lg.Error("Failed to get student options", zap.Error(err))
		return nil, err
	}
	s.lg.Info("Successfully fetched student options")
	return options, nil
}
