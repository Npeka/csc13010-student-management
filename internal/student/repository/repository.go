package repository

import (
	"context"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/student"
	"github.com/csc13010-student-management/internal/student/dtos"
	"gorm.io/gorm"
)

type studentRepository struct {
	db *gorm.DB
}

// NewStudentRepository khởi tạo repository với database
func NewStudentRepository(db *gorm.DB) student.IStudentRepository {
	return &studentRepository{db: db}
}

// GetStudents lấy danh sách tất cả sinh viên
func (s *studentRepository) GetStudents(ctx context.Context) ([]*models.Student, error) {
	var students []*models.Student
	err := s.db.WithContext(ctx).Find(&students).Error
	if err != nil {
		return nil, err
	}
	return students, err
}

// CreateStudent thêm sinh viên vào database
func (s *studentRepository) CreateStudent(ctx context.Context, student *models.Student) error {
	return s.db.WithContext(ctx).Create(&student).Error
}

// UpdateStudent cập nhật thông tin sinh viên
func (s *studentRepository) UpdateStudent(ctx context.Context, student *models.Student) error {
	return s.db.WithContext(ctx).Where("id = ?", student.ID).Updates(&student).Error
}

// DeleteStudent xóa sinh viên khỏi database theo ID
func (s *studentRepository) DeleteStudent(ctx context.Context, student_id string) error {
	return s.db.WithContext(ctx).Where("id = ?", student_id).Delete(&models.Student{}).Error
}

func (s *studentRepository) GetOptions(ctx context.Context) (*dtos.OptionDTO, error) {
	optionDTO := &dtos.OptionDTO{}

	optionMap := map[string]*[]*dtos.Option{
		"genders":   &optionDTO.Genders,
		"faculties": &optionDTO.Faculties,
		"courses":   &optionDTO.Courses,
		"programs":  &optionDTO.Programs,
		"statuses":  &optionDTO.Statuses,
	}

	modelMap := map[string]interface{}{
		"genders":   &models.Gender{},
		"faculties": &models.Faculty{},
		"courses":   &models.Course{},
		"programs":  &models.Program{},
		"statuses":  &models.Status{},
	}

	for key, model := range modelMap {
		if err := s.db.Model(model).Select("id, name").Find(optionMap[key]).Error; err != nil {
			return nil, err
		}
	}

	return optionDTO, nil
}
