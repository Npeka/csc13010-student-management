package repository

import (
	"context"
	"reflect"
	"testing"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/student"
	"github.com/csc13010-student-management/internal/student/dtos"
	"gorm.io/gorm"
)

func TestNewStudentRepository(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want student.IStudentRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStudentRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStudentRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_studentRepository_GetStudents(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.Student
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &studentRepository{
				db: tt.fields.db,
			}
			got, err := s.GetStudents(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("studentRepository.GetStudents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("studentRepository.GetStudents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_studentRepository_CreateStudents(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx      context.Context
		students []models.Student
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
			s := &studentRepository{
				db: tt.fields.db,
			}
			if err := s.CreateStudents(tt.args.ctx, tt.args.students); (err != nil) != tt.wantErr {
				t.Errorf("studentRepository.CreateStudents() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_studentRepository_GetStudentByStudentID(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx       context.Context
		studentID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Student
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &studentRepository{
				db: tt.fields.db,
			}
			got, err := s.GetStudentByStudentID(tt.args.ctx, tt.args.studentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("studentRepository.GetStudentByStudentID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("studentRepository.GetStudentByStudentID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_studentRepository_GetFullInfoStudentByStudentID(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx       context.Context
		studentID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dtos.StudentDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &studentRepository{
				db: tt.fields.db,
			}
			got, err := s.GetFullInfoStudentByStudentID(tt.args.ctx, tt.args.studentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("studentRepository.GetFullInfoStudentByStudentID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("studentRepository.GetFullInfoStudentByStudentID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_studentRepository_CreateStudent(t *testing.T) {
	type fields struct {
		db *gorm.DB
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
			s := &studentRepository{
				db: tt.fields.db,
			}
			if err := s.CreateStudent(tt.args.ctx, tt.args.student); (err != nil) != tt.wantErr {
				t.Errorf("studentRepository.CreateStudent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_studentRepository_UpdateStudent(t *testing.T) {
	type fields struct {
		db *gorm.DB
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
			s := &studentRepository{
				db: tt.fields.db,
			}
			if err := s.UpdateStudent(tt.args.ctx, tt.args.student); (err != nil) != tt.wantErr {
				t.Errorf("studentRepository.UpdateStudent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_studentRepository_DeleteStudent(t *testing.T) {
	type fields struct {
		db *gorm.DB
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
			s := &studentRepository{
				db: tt.fields.db,
			}
			if err := s.DeleteStudent(tt.args.ctx, tt.args.studentID); (err != nil) != tt.wantErr {
				t.Errorf("studentRepository.DeleteStudent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_studentRepository_GetOptions(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dtos.OptionDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &studentRepository{
				db: tt.fields.db,
			}
			got, err := s.GetOptions(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("studentRepository.GetOptions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("studentRepository.GetOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}
