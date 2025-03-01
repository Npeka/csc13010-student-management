package program

import (
	"context"

	"github.com/csc13010-student-management/internal/models"
)

type IProgramRepository interface {
	GetPrograms(ctx context.Context) ([]*models.Program, error)
	CreateProgram(ctx context.Context, program *models.Program) error
	DeleteProgram(ctx context.Context, id uint) error
}
