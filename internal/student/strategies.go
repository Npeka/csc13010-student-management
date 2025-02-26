package student

import (
	"context"

	"github.com/csc13010-student-management/internal/models"
)

type ImportStrategy interface {
	Import(ctx context.Context, filePath string) ([]models.Student, error)
}

type ExportStrategy interface {
	Export(ctx context.Context, students []models.Student, filePath string) error
}
