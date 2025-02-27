package repository

import (
	"context"
	"fmt"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/program"
	"gorm.io/gorm"
)

type programRepository struct {
	db *gorm.DB
}

func NewProgramRepository(db *gorm.DB) program.IProgramRepository {
	return &programRepository{db: db}
}

func (pr *programRepository) GetPrograms(ctx context.Context) ([]*models.Program, error) {
	var programs []*models.Program
	if err := pr.db.WithContext(ctx).Find(&programs).Error; err != nil {
		return nil, fmt.Errorf("failed to get programs: %w", err)
	}
	return programs, nil
}

func (pr *programRepository) CreateProgram(ctx context.Context, program *models.Program) error {
	var existing models.Program
	if err := pr.db.WithContext(ctx).Where("name = ?", program.Name).First(&existing).Error; err == nil {
		return fmt.Errorf("program already exists: %w", err)
	}
	if err := pr.db.WithContext(ctx).Create(program).Error; err != nil {
		return fmt.Errorf("failed to create program: %w", err)
	}
	return nil
}

func (pr *programRepository) DeleteProgram(ctx context.Context, programID int) error {
	if err := pr.db.WithContext(ctx).Where("id = ?", programID).Delete(&models.Program{}).Error; err != nil {
		return fmt.Errorf("failed to delete program: %w", err)
	}
	return nil
}
