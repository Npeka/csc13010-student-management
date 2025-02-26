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

	ImportStudents(ctx context.Context, filePath string) error
	ExportStudents(ctx context.Context, format string) (string, error)

	GetOptions(ctx context.Context) (*dtos.OptionDTO, error)

	GetFaculties(ctx context.Context) ([]*models.Faculty, error)
	GetPrograms(ctx context.Context) ([]*models.Program, error)
	GetStatuses(ctx context.Context) ([]*models.Status, error)

	CreateFaculty(ctx context.Context, faculty *models.Faculty) error
	CreateProgram(ctx context.Context, program *models.Program) error
	CreateStatus(ctx context.Context, status *models.Status) error

	DeleteFaculty(ctx context.Context, faculty_id int) error
	DeleteProgram(ctx context.Context, program_id int) error
	DeleteStatus(ctx context.Context, status_id int) error
}
