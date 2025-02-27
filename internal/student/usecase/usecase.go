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

func (s *studentUsecase) logAndReturnError(msg string, err error) error {
	s.lg.Error(msg, zap.Error(err))
	return err
}

func (s *studentUsecase) GetStudents(ctx context.Context) (students []*models.Student, err error) {
	students, err = s.sr.GetStudents(ctx)
	if err != nil {
		return nil, s.logAndReturnError("Failed to get students", err)
	}
	s.lg.Info("Successfully fetched students")
	return students, nil
}

func (s *studentUsecase) GetStudentByStudentID(ctx context.Context, studentID string) (student *models.Student, err error) {
	student, err = s.sr.GetStudentByStudentID(ctx, studentID)
	if err != nil {
		return nil, s.logAndReturnError("Failed to get student", err)
	}
	s.lg.Info("Successfully fetched student", zap.String("id", studentID))
	return student, nil
}

func (s *studentUsecase) GetFullInfoStudentByStudentID(ctx context.Context, studentID string) (student *dtos.StudentDTO, err error) {
	student, err = s.sr.GetFullInfoStudentByStudentID(ctx, studentID)
	if err != nil {
		return nil, s.logAndReturnError("Failed to get full info student", err)
	}
	return student, nil
}

func (s *studentUsecase) CreateStudent(ctx context.Context, student *models.Student) error {
	if err := s.sr.CreateStudent(ctx, student); err != nil {
		return s.logAndReturnError("Failed to create student", err)
	}
	s.lg.Info("Successfully created student", zap.Int("id", student.ID))
	return nil
}

func (s *studentUsecase) UpdateStudent(ctx context.Context, student *models.Student) error {
	if err := s.sr.UpdateStudent(ctx, student); err != nil {
		return s.logAndReturnError("Failed to update student", err)
	}
	s.lg.Info("Successfully updated student", zap.Int("id", student.ID))
	return nil
}

func (s *studentUsecase) DeleteStudent(ctx context.Context, studentID string) error {
	if err := s.sr.DeleteStudent(ctx, studentID); err != nil {
		return s.logAndReturnError("Failed to delete student", err)
	}
	s.lg.Info("Successfully deleted student", zap.String("id", studentID))
	return nil
}

func (s *studentUsecase) GetOptions(ctx context.Context) (options *dtos.OptionDTO, err error) {
	options, err = s.sr.GetOptions(ctx)
	if err != nil {
		return nil, s.logAndReturnError("Failed to get student options", err)
	}
	s.lg.Info("Successfully fetched student options")
	return options, nil
}
