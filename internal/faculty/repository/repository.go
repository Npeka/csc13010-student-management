package repository

import (
	"context"
	"fmt"

	"github.com/csc13010-student-management/internal/faculty"
	"github.com/csc13010-student-management/internal/models"
	"gorm.io/gorm"
)

type facultyRepository struct {
	db *gorm.DB
}

func NewFacultyRepository(db *gorm.DB) faculty.IFacultyRepository {
	return &facultyRepository{db: db}
}

func (fr *facultyRepository) GetFaculties(ctx context.Context) ([]*models.Faculty, error) {
	var faculties []*models.Faculty
	err := fr.db.WithContext(ctx).Find(&faculties).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get faculties: %w", err)
	}
	return faculties, nil
}

func (fr *facultyRepository) CreateFaculty(ctx context.Context, faculty *models.Faculty) error {
	var existing models.Faculty
	if err := fr.db.WithContext(ctx).Where("name = ?", faculty.Name).First(&existing).Error; err == nil {
		return fmt.Errorf("faculty already exists: %w", err)
	}
	if err := fr.db.WithContext(ctx).Create(&faculty).Error; err != nil {
		return fmt.Errorf("failed to create faculty: %w", err)
	}
	return nil
}

func (fr *facultyRepository) DeleteFaculty(ctx context.Context, id int) error {
	if err := fr.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Faculty{}).Error; err != nil {
		return fmt.Errorf("failed to delete faculty: %w", err)
	}
	return nil
}
