package repository

import (
	"context"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/program"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type programRepository struct {
	db *gorm.DB
}

func NewProgramRepository(db *gorm.DB) program.IProgramRepository {
	return &programRepository{db: db}
}

func (pr *programRepository) GetPrograms(ctx context.Context) ([]*models.Program, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "programRepository.GetPrograms")
	defer span.Finish()

	var programs []*models.Program
	if err := pr.db.WithContext(ctx).Find(&programs).Error; err != nil {
		return nil, errors.Wrap(err, "programRepository.GetPrograms.Find")
	}
	return programs, nil
}

func (pr *programRepository) CreateProgram(ctx context.Context, program *models.Program) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "programRepository.CreateProgram")
	defer span.Finish()

	if err := pr.db.WithContext(ctx).Create(program).Error; err != nil {
		return errors.Wrap(err, "programRepository.CreateProgram.Create")
	}
	return nil
}

func (pr *programRepository) UpdateProgram(ctx context.Context, program *models.Program) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "programRepository.UpdateProgram")
	defer span.Finish()

	if err := pr.db.WithContext(ctx).Model(program).Updates(program).Error; err != nil {
		return errors.Wrap(err, "programRepository.UpdateProgram.Updates")
	}
	return nil
}

func (pr *programRepository) DeleteProgram(ctx context.Context, id uint) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "programRepository.DeleteProgram")
	defer span.Finish()

	if err := pr.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Program{}).Error; err != nil {
		return errors.Wrap(err, "programRepository.DeleteProgram.Delete")
	}
	return nil
}
