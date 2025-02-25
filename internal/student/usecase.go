package student

import (
	"context"

	"github.com/csc13010-student-management/internal/models"
)

type IStudentUsecase interface {
	GetStudents(ctx context.Context) ([]*models.Student, error)
	CreateStudent(ctx context.Context, student *models.Student) error
	UpdateStudent(ctx context.Context, student *models.Student) error
	DeleteStudent(ctx context.Context, student_id string) error
	SearchStudents(ctx context.Context, query string) ([]*models.Student, error)
}
