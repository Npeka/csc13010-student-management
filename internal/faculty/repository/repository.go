package repository

import (
	"context"

	"github.com/csc13010-student-management/internal/faculty"
	"github.com/csc13010-student-management/internal/models"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type facultyRepository struct {
	db *gorm.DB
}

func NewFacultyRepository(db *gorm.DB) faculty.IFacultyRepository {
	return &facultyRepository{db: db}
}

func (fr *facultyRepository) GetFaculties(ctx context.Context) ([]*models.Faculty, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "facultyRepository.GetFaculties")
	defer span.Finish()

	var faculties []*models.Faculty
	err := fr.db.WithContext(ctx).Find(&faculties).Error
	if err != nil {
		return nil, errors.Wrap(err, "facultyRepository.GetFaculties.Find")
	}
	return faculties, nil
}

func (fr *facultyRepository) CreateFaculty(ctx context.Context, faculty *models.Faculty) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "facultyRepository.CreateFaculty")
	defer span.Finish()

	if err := fr.db.WithContext(ctx).Create(&faculty).Error; err != nil {
		return errors.Wrap(err, "facultyRepository.CreateFaculty.Create")
	}
	return nil
}

func (fr *facultyRepository) UpdateFaculty(ctx context.Context, faculty *models.Faculty) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "facultyRepository.UpdateFaculty")
	defer span.Finish()

	if err := fr.db.WithContext(ctx).Where("id = ?", faculty.ID).Updates(&faculty).Error; err != nil {
		return errors.Wrap(err, "facultyRepository.UpdateFaculty.Updates")
	}
	return nil
}

func (fr *facultyRepository) DeleteFaculty(ctx context.Context, id uint) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "facultyRepository.DeleteFaculty")
	defer span.Finish()

	if err := fr.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Faculty{}).Error; err != nil {
		return errors.Wrap(err, "facultyRepository.DeleteFaculty.Delete")
	}
	return nil
}
