package usecase

import (
	"context"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/student"
	"github.com/csc13010-student-management/internal/student/dtos"
	"github.com/csc13010-student-management/internal/student/strategies"
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

func (s *studentUsecase) GetStudentByStudentID(ctx context.Context, student_id string) (*models.Student, error) {
	student, err := s.sr.GetStudentByStudentID(ctx, student_id)
	if err != nil {
		s.lg.Error("Failed to get student", zap.Error(err))
		return nil, err
	}
	s.lg.Info("Successfully fetched student", zap.String("id", student_id))
	return student, nil
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

func (s *studentUsecase) ExportStudents(ctx context.Context, format string) (string, error) {
	var students []models.Student
	if err := s.sr.GetAllStudents(ctx, &students); err != nil {
		return "", err
	}

	// Xác định đường dẫn file export
	filePath := "exports/students." + format

	exportCtx, err := strategies.NewExportContext(filePath)
	if err != nil {
		return "", err
	}

	err = exportCtx.ExecuteExport(ctx, students, filePath)
	if err != nil {
		return "", err
	}

	return filePath, nil
}


func (s *studentUsecase) ImportStudents(ctx context.Context, filePath string) error {
	importCtx, err := strategies.NewImportContext(filePath)
	if err != nil {
		return err
	}

	students, err := importCtx.ExecuteImport(ctx, filePath)
	if err != nil {
		return err
	}

	// Lưu vào database
	return s.sr.BatchInsertStudents(ctx, students)
}
