package usecase

import (
	"context"
	"encoding/json"

	"github.com/casbin/casbin/v2"
	"github.com/csc13010-student-management/internal/events"
	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/student"
	"github.com/csc13010-student-management/internal/student/dtos"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/csc13010-student-management/pkg/logger"
	"github.com/segmentio/kafka-go"
)

type studentUsecase struct {
	sr student.IStudentRepository
	lg *logger.LoggerZap
	kw *kafka.Writer
	e  *casbin.Enforcer
}

func NewStudentUsecase(
	sr student.IStudentRepository,
	lg *logger.LoggerZap,
	kw *kafka.Writer,
	e *casbin.Enforcer,
) student.IStudentUsecase {
	return &studentUsecase{
		sr: sr,
		lg: lg,
		kw: kw,
		e:  e,
	}
}

func (su *studentUsecase) GetStudents(ctx context.Context) (students []*models.Student, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "studentUsecase.GetStudents")
	defer span.Finish()

	students, err = su.sr.GetStudents(ctx)
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (su *studentUsecase) GetStudentByStudentID(ctx context.Context, studentID string) (student *models.Student, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "studentUsecase.GetStudentByStudentID")
	defer span.Finish()

	student, err = su.sr.GetStudentByStudentID(ctx, studentID)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (su *studentUsecase) GetFullInfoStudentByStudentID(ctx context.Context, studentID string) (student *dtos.StudentDTO, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "studentUsecase.GetFullInfoStudentByStudentID")
	defer span.Finish()

	student, err = su.sr.GetFullInfoStudentByStudentID(ctx, studentID)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (su *studentUsecase) CreateStudent(ctx context.Context, student *models.Student) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "studentUsecase.CreateStudent")
	defer span.Finish()

	// Create Student
	if err := su.sr.CreateStudent(ctx, student); err != nil {
		return err
	}

	// Publish Kafka event
	studentJson, err := json.Marshal(events.StudentCreatedEvent{
		Username: student.StudentID,
		Role:     models.RoleStudent,
	})
	if err != nil {
		return errors.Wrap(err, "studentUsecase.CreateStudent.json.Marshal")
	}
	if err := su.kw.WriteMessages(ctx, kafka.Message{Value: studentJson}); err != nil {
		return errors.Wrap(err, "studentUsecase.CreateStudent.kw.WriteMessages")
	}
	return nil
}

func (su *studentUsecase) UpdateStudent(ctx context.Context, student *models.Student) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "studentUsecase.UpdateStudent")
	defer span.Finish()

	if err := su.sr.UpdateStudent(ctx, student); err != nil {
		return err
	}
	return nil
}

func (su *studentUsecase) UpdateUserIDByUsername(ctx context.Context, studentID string, userID uint) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "studentUsecase.UpdateUserIDByUsername")
	defer span.Finish()

	if err := su.sr.UpdateUserIDByUsername(ctx, studentID, userID); err != nil {
		return err
	}
	return nil
}

func (su *studentUsecase) DeleteStudent(ctx context.Context, studentID string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "studentUsecase.DeleteStudent")
	defer span.Finish()

	if err := su.sr.DeleteStudent(ctx, studentID); err != nil {
		return err
	}
	return nil
}

func (su *studentUsecase) GetOptions(ctx context.Context) (optionsu *dtos.OptionDTO, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "studentUsecase.GetOptions")
	defer span.Finish()

	options, err := su.sr.GetOptions(ctx)
	if err != nil {
		return nil, err
	}
	return options, nil
}
