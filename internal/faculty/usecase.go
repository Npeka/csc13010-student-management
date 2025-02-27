package faculty

import (
	"context"

	"github.com/csc13010-student-management/internal/models"
)

type IFacultyUsecase interface {
	GetFaculties(ctx context.Context) ([]*models.Faculty, error)
	CreateFaculty(ctx context.Context, faculty *models.Faculty) error
	DeleteFaculty(ctx context.Context, id int) error
}
