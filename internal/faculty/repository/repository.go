package repository

import (
	"context"

	"github.com/csc13010-student-management/internal/faculty"
	"github.com/csc13010-student-management/internal/models"
	"gorm.io/gorm"
)

type facultyRepository struct {
	db *gorm.DB
}

func NewfacultyRepository(db *gorm.DB) faculty.IFacultyRepository {
	return &facultyRepository{db: db}
}

func (f *facultyRepository) CreateFaculty(ctx context.Context, faculty *models.Faculty) error {
	panic("unimplemented")
}

func (f *facultyRepository) DeleteFaculty(ctx context.Context, id int) error {
	panic("unimplemented")
}

func (f *facultyRepository) GetFaculties(ctx context.Context) ([]*models.Faculty, error) {
	panic("unimplemented")
}
