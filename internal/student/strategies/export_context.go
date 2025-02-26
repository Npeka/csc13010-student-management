package strategies

import (
	"context"
	"errors"
	"path/filepath"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/student"
)

type ExportContext struct {
	strategy student.ExportStrategy
}

func (c *ExportContext) SetStrategy(strategy student.ExportStrategy) {
	c.strategy = strategy
}

func (c *ExportContext) ExecuteExport(ctx context.Context, students []models.Student, filePath string) error {
	if c.strategy == nil {
		return errors.New("export strategy not set")
	}
	return c.strategy.Export(ctx, students, filePath)
}

func NewExportContext(filePath string) (*ExportContext, error) {
	ctx := &ExportContext{}
	switch filepath.Ext(filePath) {
	case ".csv":
		ctx.SetStrategy(&CSVExportStrategy{})
	case ".json":
		ctx.SetStrategy(&JSONExportStrategy{})
	default:
		return nil, errors.New("unsupported file format")
	}
	return ctx, nil
}
