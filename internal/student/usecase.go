package student

import (
	"context"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/student/dtos"
	"github.com/google/uuid"
)

type IStudentUsecase interface {
	GetStudents(ctx context.Context) ([]*dtos.StudentResponseDTO, error)
	GetStudentByStudentID(ctx context.Context, student_id string) (*models.Student, error)
	GetFullInfoStudentByStudentID(ctx context.Context, student_id string) (*dtos.StudentResponseDTO, error)

	CreateStudent(ctx context.Context, student *models.Student) error
	UpdateStudent(ctx context.Context, student *models.Student) error
	UpdateUserIDByUsername(ctx context.Context, student_id string, user_id uuid.UUID) error
	DeleteStudent(ctx context.Context, student_id string) error
	GetOptions(ctx context.Context) (*dtos.OptionDTO, error)
	BatchUpdateUserIDs(ctx context.Context, userIDs map[string]uuid.UUID) error
}
