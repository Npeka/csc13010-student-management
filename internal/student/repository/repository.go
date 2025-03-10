package repository

import (
	"context"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/student"
	"github.com/csc13010-student-management/internal/student/dtos"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) student.IStudentRepository {
	return &studentRepository{db: db}
}

func (s *studentRepository) GetStudents(ctx context.Context) ([]*dtos.StudentResponseDTO, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "studentRepository.GetStudents")
	defer span.Finish()

	var students []*dtos.StudentResponseDTO

	err := s.db.WithContext(ctx).
		Table("students").
		Select(`
			students.id, students.student_id, students.full_name, students.birth_date, students.email, students.phone, students.address,
			genders.name AS gender, faculties.name AS faculty, 
			courses.name AS course, programs.name AS program, statuses.name AS status
		`).
		Joins("LEFT JOIN genders ON genders.id = students.gender_id").
		Joins("LEFT JOIN faculties ON faculties.id = students.faculty_id").
		Joins("LEFT JOIN courses ON courses.id = students.course_id").
		Joins("LEFT JOIN programs ON programs.id = students.program_id").
		Joins("LEFT JOIN statuses ON statuses.id = students.status_id").
		Scan(&students).Error

	if err != nil {
		return nil, err
	}

	return students, nil
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

func (s *studentRepository) GetFullInfoStudentByStudentID(ctx context.Context, studentID string) (*dtos.StudentResponseDTO, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "studentRepository.GetFullInfoStudentByStudentID")
	defer span.Finish()

	studentDTO := &dtos.StudentResponseDTO{}
	err := s.db.WithContext(ctx).
		Table("students").
		Select(`students.id, students.student_id, students.full_name, students.birth_date, students.email, students.phone,
			genders.name AS gender, faculties.name AS faculty, 
			courses.name AS course, programs.name AS program, statuses.name AS status`).
		Joins("LEFT JOIN genders ON genders.id = students.gender_id").
		Joins("LEFT JOIN faculties ON faculties.id = students.faculty_id").
		Joins("LEFT JOIN courses ON courses.id = students.course_id").
		Joins("LEFT JOIN programs ON programs.id = students.program_id").
		Joins("LEFT JOIN statuses ON statuses.id = students.status_id").
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

func (s *studentRepository) CreateStudents(ctx context.Context, students []models.Student) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "studentRepository.CreateStudents")
	defer span.Finish()

	return s.db.WithContext(ctx).Create(&students).Error
}

func (s *studentRepository) CreateStudent(ctx context.Context, student *models.Student) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "studentRepository.CreateStudent")
	defer span.Finish()

	return s.db.WithContext(ctx).Create(student).Error
}

func (s *studentRepository) UpdateStudent(ctx context.Context, student *models.Student) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "studentRepository.UpdateStudent")
	defer span.Finish()

	if err := s.db.WithContext(ctx).
		Model(student).
		Where("student_id = ?", student.StudentID).
		Updates(student).Error; err != nil {
		return errors.Wrap(err, "studentRepository.UpdateStudent.Save")
	}
	return nil
}

func (s *studentRepository) UpdateUserIDByUsername(ctx context.Context, studentID string, userID uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "studentRepository.UpdateUserIDByUsername")
	defer span.Finish()

	return s.db.WithContext(ctx).
		Model(&models.Student{}).
		Where("student_id = ?", studentID).
		UpdateColumn("user_id", userID).
		Error
}

func (s *studentRepository) DeleteStudent(ctx context.Context, student_id string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "studentRepository.DeleteStudent")
	defer span.Finish()

	return s.db.WithContext(ctx).Where("student_id = ?", student_id).Delete(&models.Student{}).Error
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

func (s *studentRepository) BatchUpdateUserIDs(ctx context.Context, studentIDs map[string]uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "studentRepository.BatchUpdateUserIDs")
	defer span.Finish()

	tx := s.db.Begin()
	for studentID, userID := range studentIDs {
		if err := tx.Model(&models.Student{}).
			Where("student_id = ?", studentID).
			UpdateColumn("user_id", userID).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}
