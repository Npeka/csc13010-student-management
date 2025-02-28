package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/student"
	"github.com/csc13010-student-management/internal/student/dtos"
	"github.com/csc13010-student-management/internal/student/mocks"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/golang/mock/gomock"
)

func TestNewStudentUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIStudentRepository(ctrl)
	mockLogger := logger.NewLoggerTest()

	type args struct {
		sr student.IStudentRepository
		lg *logger.LoggerZap
	}
	tests := []struct {
		name string
		args args
		want student.IStudentUsecase
	}{
		{
			name: "Success - Create Student Usecase",
			args: args{
				sr: mockRepo,
				lg: mockLogger,
			},
			want: NewStudentUsecase(mockRepo, mockLogger),
		},
		{
			name: "Failed - Create Student Usecase",
			args: args{
				sr: nil,
				lg: nil,
			},
			want: NewStudentUsecase(nil, nil),
		},
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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIStudentRepository(ctrl)
	mockLogger := logger.NewLoggerTest()

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
		{
			name: "Log and return error",
			fields: fields{
				sr: mockRepo,
				lg: mockLogger,
			},
			args: args{
				msg: "Test error",
				err: errors.New("test error"),
			},
			wantErr: true,
		},
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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIStudentRepository(ctrl)
	mockLogger := logger.NewLoggerTest()
	mockStudents := []*models.Student{
		{StudentID: "22127180", FullName: "Nguyen Phuc Khang"},
		{StudentID: "22127181", FullName: "John Doe"},
	}

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
		setup        func()
	}{
		{
			name: "Success - Get Students",
			fields: fields{
				sr: mockRepo,
				lg: mockLogger,
			},
			args: args{
				ctx: context.Background(),
			},
			wantStudents: mockStudents,
			wantErr:      false,
			setup: func() {
				mockRepo.EXPECT().GetStudents(gomock.Any()).Return(mockStudents, nil)
			},
		},
		{
			name: "Failed - Get Students",
			fields: fields{
				sr: mockRepo,
				lg: mockLogger,
			},
			args: args{
				ctx: context.Background(),
			},
			wantStudents: nil,
			wantErr:      true,
			setup: func() {
				mockRepo.EXPECT().GetStudents(gomock.Any()).Return(nil, errors.New("failed to get students"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
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

func Test_studentUsecase_GetFullInfoStudentByStudentID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIStudentRepository(ctrl)
	mockLogger := logger.NewLoggerTest()
	mockStudent := &dtos.StudentDTO{StudentID: "22127180", FullName: "Nguyen Phuc Khang"}

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
		setup       func()
	}{
		{
			name: "Success - Get Full Info Student By Student ID",
			fields: fields{
				sr: mockRepo,
				lg: mockLogger,
			},
			args: args{
				ctx:       context.Background(),
				studentID: "22127180",
			},
			wantStudent: mockStudent,
			wantErr:     false,
			setup: func() {
				mockRepo.EXPECT().GetFullInfoStudentByStudentID(gomock.Any(), "22127180").Return(mockStudent, nil)
			},
		},
		{
			name: "Failed - Get Full Info Student By Student ID",
			fields: fields{
				sr: mockRepo,
				lg: mockLogger,
			},
			args: args{
				ctx:       context.Background(),
				studentID: "22127180",
			},
			wantStudent: nil,
			wantErr:     true,
			setup: func() {
				mockRepo.EXPECT().GetFullInfoStudentByStudentID(gomock.Any(), "22127180").Return(nil, errors.New("failed to get student"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIStudentRepository(ctrl)
	mockLogger := logger.NewLoggerTest()
	mockStudent := &models.Student{StudentID: "22127180", FullName: "Nguyen Phuc Khang"}

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
		setup   func()
	}{
		{
			name: "Success - Create Student",
			fields: fields{
				sr: mockRepo,
				lg: mockLogger,
			},
			args: args{
				ctx:     context.Background(),
				student: mockStudent,
			},
			wantErr: false,
			setup: func() {
				mockRepo.EXPECT().CreateStudent(gomock.Any(), mockStudent).Return(nil)
			},
		},
		{
			name: "Failed - Create Student",
			fields: fields{
				sr: mockRepo,
				lg: mockLogger,
			},
			args: args{
				ctx:     context.Background(),
				student: mockStudent,
			},
			wantErr: true,
			setup: func() {
				mockRepo.EXPECT().CreateStudent(gomock.Any(), mockStudent).Return(errors.New("failed to create student"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIStudentRepository(ctrl)
	mockLogger := logger.NewLoggerTest()
	mockStudent := &models.Student{StudentID: "22127180", FullName: "Nguyen Phuc Khang"}

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
		setup   func()
	}{
		{
			name: "Success - Update Student",
			fields: fields{
				sr: mockRepo,
				lg: mockLogger,
			},
			args: args{
				ctx:     context.Background(),
				student: mockStudent,
			},
			wantErr: false,
			setup: func() {
				mockRepo.EXPECT().UpdateStudent(gomock.Any(), mockStudent).Return(nil)
			},
		},
		{
			name: "Failed - Update Student",
			fields: fields{
				sr: mockRepo,
				lg: mockLogger,
			},
			args: args{
				ctx:     context.Background(),
				student: mockStudent,
			},
			wantErr: true,
			setup: func() {
				mockRepo.EXPECT().UpdateStudent(gomock.Any(), mockStudent).Return(errors.New("failed to update student"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIStudentRepository(ctrl)
	mockLogger := logger.NewLoggerTest()

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
		setup   func()
	}{
		{
			name: "Success - Delete Student",
			fields: fields{
				sr: mockRepo,
				lg: mockLogger,
			},
			args: args{
				ctx:       context.Background(),
				studentID: "22127180",
			},
			wantErr: false,
			setup: func() {
				mockRepo.EXPECT().DeleteStudent(gomock.Any(), "22127180").Return(nil)
			},
		},
		{
			name: "Failed - Delete Student",
			fields: fields{
				sr: mockRepo,
				lg: mockLogger,
			},
			args: args{
				ctx:       context.Background(),
				studentID: "22127180",
			},
			wantErr: true,
			setup: func() {
				mockRepo.EXPECT().DeleteStudent(gomock.Any(), "22127180").Return(errors.New("failed to delete student"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIStudentRepository(ctrl)
	mockLogger := logger.NewLoggerTest()
	mockOptions := &dtos.OptionDTO{
		Genders:   []*dtos.Option{{ID: 1, Name: "Male"}, {ID: 2, Name: "Female"}},
		Faculties: []*dtos.Option{{ID: 1, Name: "Engineering"}, {ID: 2, Name: "Arts"}},
		Courses:   []*dtos.Option{{ID: 1, Name: "Computer Science"}, {ID: 2, Name: "Mathematics"}},
		Programs:  []*dtos.Option{{ID: 1, Name: "Undergraduate"}, {ID: 2, Name: "Postgraduate"}},
		Statuses:  []*dtos.Option{{ID: 1, Name: "Active"}, {ID: 2, Name: "Inactive"}},
	}

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
		setup       func()
	}{
		{
			name: "Success - Get Options",
			fields: fields{
				sr: mockRepo,
				lg: mockLogger,
			},
			args: args{
				ctx: context.Background(),
			},
			wantOptions: mockOptions,
			wantErr:     false,
			setup: func() {
				mockRepo.EXPECT().GetOptions(gomock.Any()).Return(mockOptions, nil)
			},
		},
		{
			name: "Failed - Get Options",
			fields: fields{
				sr: mockRepo,
				lg: mockLogger,
			},
			args: args{
				ctx: context.Background(),
			},
			wantOptions: nil,
			wantErr:     true,
			setup: func() {
				mockRepo.EXPECT().GetOptions(gomock.Any()).Return(nil, errors.New("failed to get options"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
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
