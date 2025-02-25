package repository

import (
	"context"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/student"
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

// SearchStudents tìm kiếm sinh viên theo tên hoặc email
func (s *studentRepository) SearchStudents(ctx context.Context, query string) ([]*models.Student, error) {
	var students []*models.Student
	err := s.db.WithContext(ctx).
		Where("LOWER(full_name) LIKE LOWER(?) OR LOWER(email) LIKE LOWER(?)", "%"+query+"%", "%"+query+"%").
		Find(&students).Error
	return students, err
}
