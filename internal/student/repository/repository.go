package repository

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"reflect"

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

// GetStudentByStudentID lấy thông tin sinh viên theo ID
func (s *studentRepository) GetStudentByStudentID(ctx context.Context, student_id string) (*models.Student, error) {
	student := &models.Student{}
	err := s.db.WithContext(ctx).Where("student_id = ?", student_id).First(&student).Error
	if err != nil {
		return nil, err
	}
	return student, err
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

// Lấy danh sách tag JSON từ struct
func getStructTags(v interface{}) []string {
	t := reflect.TypeOf(v)
	fields := make([]string, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		fields[i] = t.Field(i).Tag.Get("json")
	}
	return fields
}

// Chuyển struct thành slice string để ghi vào CSV
func structToSlice(v interface{}) ([]string, error) {
	val := reflect.ValueOf(v)
	var record []string

	for i := 0; i < val.NumField(); i++ {
		record = append(record, fmt.Sprintf("%v", val.Field(i).Interface()))
	}
	return record, nil
}

func (s *studentRepository) ExportStudentsToCSV(ctx context.Context) (string, error) {
	var students []models.Student
	if err := s.db.Find(&students).Error; err != nil {
		return "", err
	}

	// Tạo đường dẫn file CSV
	filePath := "exports/students.csv"
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	// Tạo header từ các field của struct Student
	headers := getStructTags(models.Student{})
	if err := writer.Write(headers); err != nil {
		return "", err
	}

	// Ghi dữ liệu của từng student
	for _, student := range students {
		record, err := structToSlice(student)
		if err != nil {
			return "", err
		}
		if err := writer.Write(record); err != nil {
			return "", err
		}
	}

	writer.Flush() // Đảm bảo tất cả dữ liệu được ghi vào file
	return filePath, nil
}

// ExportStudentsToJSON exports student data to a JSON file
func (s *studentRepository) GetAllStudents(ctx context.Context, students *[]models.Student) error {
	return s.db.Find(students).Error
}

func (s *studentRepository) BatchInsertStudents(ctx context.Context, students []models.Student) error {
	return s.db.Create(&students).Error
}
