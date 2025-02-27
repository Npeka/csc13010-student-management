package repository

import (
	"context"
	"errors"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/student"
	"github.com/csc13010-student-management/internal/student/dtos"
	"gorm.io/gorm"
)

type studentRepository struct {
	db *gorm.DB
}

// NewStudentRepository initializes the repository with the database
func NewStudentRepository(db *gorm.DB) student.IStudentRepository {
	return &studentRepository{db: db}
}

// GetStudents retrieves the list of all students
func (s *studentRepository) GetStudents(ctx context.Context) ([]*models.Student, error) {
	var students []*models.Student
	if err := s.db.WithContext(ctx).Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

// CreateStudents adds multiple students to the database
func (s *studentRepository) CreateStudents(ctx context.Context, students []models.Student) error {
	return s.db.WithContext(ctx).Create(&students).Error
}

// GetStudentByStudentID retrieves student information by ID
func (s *studentRepository) GetStudentByStudentID(ctx context.Context, studentID string) (*models.Student, error) {
	student := &models.Student{}
	if err := s.db.WithContext(ctx).Where("student_id = ?", studentID).First(student).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return student, nil
}

// GetFullInfoStudentByStudentID retrieves full student information by ID
func (s *studentRepository) GetFullInfoStudentByStudentID(ctx context.Context, studentID string) (*dtos.StudentDTO, error) {
	studentDTO := &dtos.StudentDTO{}
	err := s.db.WithContext(ctx).
		Table("students").
		Select("students.id, students.student_id, students.full_name, students.birth_date, students.gender_id, genders.name as gender_name, students.faculty_id, faculties.name as faculty_name, students.course_id, courses.name as course_name, students.program_id, programs.name as program_name, students.status_id, statuses.name as status_name").
		Joins("left join genders on students.gender_id = genders.id").
		Joins("left join faculties on students.faculty_id = faculties.id").
		Joins("left join courses on students.course_id = courses.id").
		Joins("left join programs on students.program_id = programs.id").
		Joins("left join statuses on students.status_id = statuses.id").
		Where("students.student_id = ?", studentID).
		First(studentDTO).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return studentDTO, nil
}

// CreateStudent adds a student to the database
func (s *studentRepository) CreateStudent(ctx context.Context, student *models.Student) error {
	return s.db.WithContext(ctx).Create(student).Error
}

// UpdateStudent updates student information
func (s *studentRepository) UpdateStudent(ctx context.Context, student *models.Student) error {
	return s.db.WithContext(ctx).Where("student_id = ?", student.StudentID).Updates(student).Error
}

// DeleteStudent removes a student from the database by ID
func (s *studentRepository) DeleteStudent(ctx context.Context, studentID string) error {
	return s.db.WithContext(ctx).Where("student_id = ?", studentID).Delete(&models.Student{}).Error
}

// GetOptions retrieves various options for students
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
