package student

import (
	"context"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/student/dtos"
)

type IStudentUsecase interface {
	GetStudents(ctx context.Context) ([]*models.Student, error)
	GetStudentByStudentID(ctx context.Context, student_id string) (*models.Student, error)
	CreateStudent(ctx context.Context, student *models.Student) error
	UpdateStudent(ctx context.Context, student *models.Student) error
	DeleteStudent(ctx context.Context, student_id string) error
	GetOptions(ctx context.Context) (*dtos.OptionDTO, error)
	ImportStudents(ctx context.Context, filePath string) error
	ExportStudents(ctx context.Context, format string) (string, error)
}
