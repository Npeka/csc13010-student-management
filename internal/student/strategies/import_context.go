package strategies

import (
	"context"
	"errors"
	"path/filepath"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/student"
)

type ImportContext struct {
	strategy student.ImportStrategy
}

func (c *ImportContext) SetStrategy(strategy student.ImportStrategy) {
	c.strategy = strategy
}

func (c *ImportContext) ExecuteImport(ctx context.Context, filePath string) ([]models.Student, error) {
	if c.strategy == nil {
		return nil, errors.New("import strategy not set")
	}
	return c.strategy.Import(ctx, filePath)
}

func NewImportContext(filePath string) (*ImportContext, error) {
	ctx := &ImportContext{}
	switch filepath.Ext(filePath) {
	case ".csv":
		ctx.SetStrategy(&CSVImportStrategy{})
	case ".json":
		ctx.SetStrategy(&JSONImportStrategy{})
	default:
		return nil, errors.New("unsupported file format")
	}
	return ctx, nil
}
