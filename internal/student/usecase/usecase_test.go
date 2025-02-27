package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/student"
	"github.com/csc13010-student-management/internal/student/dtos"
	"github.com/csc13010-student-management/pkg/logger"
)

func TestNewStudentUsecase(t *testing.T) {
	type args struct {
		sr student.IStudentRepository
		lg *logger.LoggerZap
	}
	tests := []struct {
		name string
		args args
		want student.IStudentUsecase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStudentUsecase(tt.args.sr, tt.args.lg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStudentUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_studentUsecase_logAndReturnError(t *testing.T) {
	type fields struct {
		sr student.IStudentRepository
		lg *logger.LoggerZap
	}
	type args struct {
		msg string
		err error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &studentUsecase{
				sr: tt.fields.sr,
				lg: tt.fields.lg,
			}
			if err := s.logAndReturnError(tt.args.msg, tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("studentUsecase.logAndReturnError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_studentUsecase_GetStudents(t *testing.T) {
	type fields struct {
		sr student.IStudentRepository
		lg *logger.LoggerZap
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantStudents []*models.Student
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &studentUsecase{
				sr: tt.fields.sr,
				lg: tt.fields.lg,
			}
			gotStudents, err := s.GetStudents(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("studentUsecase.GetStudents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStudents, tt.wantStudents) {
				t.Errorf("studentUsecase.GetStudents() = %v, want %v", gotStudents, tt.wantStudents)
			}
		})
	}
}

func Test_studentUsecase_GetStudentByStudentID(t *testing.T) {
	type fields struct {
		sr student.IStudentRepository
		lg *logger.LoggerZap
	}
	type args struct {
		ctx       context.Context
		studentID string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantStudent *models.Student
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &studentUsecase{
				sr: tt.fields.sr,
				lg: tt.fields.lg,
			}
			gotStudent, err := s.GetStudentByStudentID(tt.args.ctx, tt.args.studentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("studentUsecase.GetStudentByStudentID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStudent, tt.wantStudent) {
				t.Errorf("studentUsecase.GetStudentByStudentID() = %v, want %v", gotStudent, tt.wantStudent)
			}
		})
	}
}

func Test_studentUsecase_GetFullInfoStudentByStudentID(t *testing.T) {
	type fields struct {
		sr student.IStudentRepository
		lg *logger.LoggerZap
	}
	type args struct {
		ctx       context.Context
		studentID string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantStudent *dtos.StudentDTO
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &studentUsecase{
				sr: tt.fields.sr,
				lg: tt.fields.lg,
			}
			gotStudent, err := s.GetFullInfoStudentByStudentID(tt.args.ctx, tt.args.studentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("studentUsecase.GetFullInfoStudentByStudentID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStudent, tt.wantStudent) {
				t.Errorf("studentUsecase.GetFullInfoStudentByStudentID() = %v, want %v", gotStudent, tt.wantStudent)
			}
		})
	}
}

func Test_studentUsecase_CreateStudent(t *testing.T) {
	type fields struct {
		sr student.IStudentRepository
		lg *logger.LoggerZap
	}
	type args struct {
		ctx     context.Context
		student *models.Student
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &studentUsecase{
				sr: tt.fields.sr,
				lg: tt.fields.lg,
			}
			if err := s.CreateStudent(tt.args.ctx, tt.args.student); (err != nil) != tt.wantErr {
				t.Errorf("studentUsecase.CreateStudent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_studentUsecase_UpdateStudent(t *testing.T) {
	type fields struct {
		sr student.IStudentRepository
		lg *logger.LoggerZap
	}
	type args struct {
		ctx     context.Context
		student *models.Student
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &studentUsecase{
				sr: tt.fields.sr,
				lg: tt.fields.lg,
			}
			if err := s.UpdateStudent(tt.args.ctx, tt.args.student); (err != nil) != tt.wantErr {
				t.Errorf("studentUsecase.UpdateStudent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_studentUsecase_DeleteStudent(t *testing.T) {
	type fields struct {
		sr student.IStudentRepository
		lg *logger.LoggerZap
	}
	type args struct {
		ctx       context.Context
		studentID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &studentUsecase{
				sr: tt.fields.sr,
				lg: tt.fields.lg,
			}
			if err := s.DeleteStudent(tt.args.ctx, tt.args.studentID); (err != nil) != tt.wantErr {
				t.Errorf("studentUsecase.DeleteStudent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_studentUsecase_GetOptions(t *testing.T) {
	type fields struct {
		sr student.IStudentRepository
		lg *logger.LoggerZap
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantOptions *dtos.OptionDTO
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &studentUsecase{
				sr: tt.fields.sr,
				lg: tt.fields.lg,
			}
			gotOptions, err := s.GetOptions(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("studentUsecase.GetOptions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOptions, tt.wantOptions) {
				t.Errorf("studentUsecase.GetOptions() = %v, want %v", gotOptions, tt.wantOptions)
			}
		})
	}
}
