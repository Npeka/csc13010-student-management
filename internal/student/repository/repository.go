package repository

import (
	"context"
	"errors"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/student"
	"github.com/csc13010-student-management/internal/student/dtos"
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) student.IStudentRepository {
	return &studentRepository{db: db}
}

func (s *studentRepository) GetStudents(ctx context.Context) ([]*models.Student, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "studentRepository.GetStudents")
	defer span.Finish()

	var students []*models.Student
	if err := s.db.WithContext(ctx).
		Preload("Gender", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name")
		}).
		Preload("Faculty", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name")
		}).
		Preload("Course", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name")
		}).
		Preload("Program", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name")
		}).
		Preload("Status", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name")
		}).
		Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func (s *studentRepository) CreateStudents(ctx context.Context, students []models.Student) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "studentRepository.CreateStudents")
	defer span.Finish()

	return s.db.WithContext(ctx).Create(&students).Error
}

func (s *studentRepository) GetStudentByStudentID(ctx context.Context, studentID string) (*models.Student, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "studentRepository.GetStudentByStudentID")
	defer span.Finish()

	student := &models.Student{}
	if err := s.db.WithContext(ctx).Where("student_id = ?", studentID).First(student).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return student, nil
}

func (s *studentRepository) GetFullInfoStudentByStudentID(ctx context.Context, studentID string) (*dtos.StudentDTO, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "studentRepository.GetFullInfoStudentByStudentID")
	defer span.Finish()

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

func (s *studentRepository) CreateStudent(ctx context.Context, student *models.Student) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "studentRepository.CreateStudent")
	defer span.Finish()

	return s.db.WithContext(ctx).Create(student).Error
}

func (s *studentRepository) UpdateStudent(ctx context.Context, student *models.Student) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "studentRepository.UpdateStudent")
	defer span.Finish()

	return s.db.WithContext(ctx).Where("student_id = ?", student.StudentID).Updates(student).Error
}

func (s *studentRepository) UpdateUserIDByUsername(ctx context.Context, studentID string, userID uint) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "studentRepository.UpdateUserIDByUsername")
	defer span.Finish()

	return s.db.WithContext(ctx).
		Model(&models.Student{}).
		Where("student_id = ?", studentID).
		UpdateColumn("user_id", userID).
		Error
}

func (s *studentRepository) DeleteStudent(ctx context.Context, studentID string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "studentRepository.DeleteStudent")
	defer span.Finish()

	return s.db.WithContext(ctx).Where("student_id = ?", studentID).Delete(&models.Student{}).Error
}

func (s *studentRepository) GetOptions(ctx context.Context) (*dtos.OptionDTO, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "studentRepository.GetOptions")
	defer span.Finish()

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
		if err := s.db.WithContext(ctx).Model(model).Find(optionMap[key]).Error; err != nil {
			return nil, err
		}
	}

	return optionDTO, nil
}
