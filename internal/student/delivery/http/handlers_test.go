package http

import (
	"reflect"
	"testing"

	"github.com/csc13010-student-management/internal/student"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/gin-gonic/gin"
)

func TestNewStudentHandlers(t *testing.T) {
	type args struct {
		su student.IStudentUsecase
		lg *logger.LoggerZap
	}
	tests := []struct {
		name string
		args args
		want student.IStudentHandlers
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStudentHandlers(tt.args.su, tt.args.lg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStudentHandlers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_studentHandlers_GetStudents(t *testing.T) {
	type fields struct {
		su student.IStudentUsecase
		lg *logger.LoggerZap
	}
	tests := []struct {
		name   string
		fields fields
		want   gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &studentHandlers{
				su: tt.fields.su,
				lg: tt.fields.lg,
			}
			if got := s.GetStudents(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("studentHandlers.GetStudents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_studentHandlers_GetStudentByStudentID(t *testing.T) {
	type fields struct {
		su student.IStudentUsecase
		lg *logger.LoggerZap
	}
	tests := []struct {
		name   string
		fields fields
		want   gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &studentHandlers{
				su: tt.fields.su,
				lg: tt.fields.lg,
			}
			if got := s.GetStudentByStudentID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("studentHandlers.GetStudentByStudentID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_studentHandlers_GetFullInfoStudentByStudentID(t *testing.T) {
	type fields struct {
		su student.IStudentUsecase
		lg *logger.LoggerZap
	}
	tests := []struct {
		name   string
		fields fields
		want   gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &studentHandlers{
				su: tt.fields.su,
				lg: tt.fields.lg,
			}
			if got := s.GetFullInfoStudentByStudentID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("studentHandlers.GetFullInfoStudentByStudentID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_studentHandlers_CreateStudent(t *testing.T) {
	type fields struct {
		su student.IStudentUsecase
		lg *logger.LoggerZap
	}
	tests := []struct {
		name   string
		fields fields
		want   gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &studentHandlers{
				su: tt.fields.su,
				lg: tt.fields.lg,
			}
			if got := s.CreateStudent(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("studentHandlers.CreateStudent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_studentHandlers_UpdateStudent(t *testing.T) {
	type fields struct {
		su student.IStudentUsecase
		lg *logger.LoggerZap
	}
	tests := []struct {
		name   string
		fields fields
		want   gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &studentHandlers{
				su: tt.fields.su,
				lg: tt.fields.lg,
			}
			if got := s.UpdateStudent(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("studentHandlers.UpdateStudent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_studentHandlers_DeleteStudent(t *testing.T) {
	type fields struct {
		su student.IStudentUsecase
		lg *logger.LoggerZap
	}
	tests := []struct {
		name   string
		fields fields
		want   gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &studentHandlers{
				su: tt.fields.su,
				lg: tt.fields.lg,
			}
			if got := s.DeleteStudent(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("studentHandlers.DeleteStudent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_studentHandlers_GetOptions(t *testing.T) {
	type fields struct {
		su student.IStudentUsecase
		lg *logger.LoggerZap
	}
	tests := []struct {
		name   string
		fields fields
		want   gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &studentHandlers{
				su: tt.fields.su,
				lg: tt.fields.lg,
			}
			if got := s.GetOptions(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("studentHandlers.GetOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}
