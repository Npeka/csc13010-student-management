package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/csc13010-student-management/internal/faculty"
	"github.com/csc13010-student-management/internal/faculty/mocks"
	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
)

func TestNewFacultyUsecase(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIFacultyRepository(ctrl)
	mockLogger := logger.NewLoggerTest()

	type args struct {
		fr faculty.IFacultyRepository
		lg *logger.LoggerZap
	}
	tests := []struct {
		name string
		args args
		want faculty.IFacultyUsecase
	}{
		// TODO: Add test cases.
		{
			name: "Success - Create Faculty Usecase",
			args: args{
				fr: mockRepo,
				lg: mockLogger,
			},
			want: NewFacultyUsecase(mockRepo, mockLogger),
		},
		{
			name: "Failed - Create Faculty Usecase",
			args: args{
				fr: nil,
				lg: nil,
			},
			want: NewFacultyUsecase(nil, nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFacultyUsecase(tt.args.fr, tt.args.lg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFacultyUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_facultyUsecase_GetFaculties(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIFacultyRepository(ctrl)
	mockLogger := logger.NewLoggerTest()
	mockFaculties := []*models.Faculty{
		{Name: "Computer Science"},
		{Name: "Mathematics"},
	}

	type fields struct {
		fr faculty.IFacultyRepository
		lg *logger.LoggerZap
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.Faculty
		wantErr bool
		setup   func()
	}{
		// TODO: Add test cases.
		{
			name: "Success - Get Faculties",
			fields: fields{
				fr: mockRepo,
				lg: mockLogger,
			},
			args: args{
				ctx: context.Background(),
			},
			want:    mockFaculties,
			wantErr: false,
			setup: func() {
				mockRepo.EXPECT().GetFaculties(gomock.Any()).Return(mockFaculties, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			fu := &facultyUsecase{
				fr: tt.fields.fr,
				lg: tt.fields.lg,
			}
			got, err := fu.GetFaculties(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("facultyUsecase.GetFaculties() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("facultyUsecase.GetFaculties() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_facultyUsecase_CreateFaculty(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIFacultyRepository(ctrl)
	mockLogger := logger.NewLoggerTest()
	mockFaculties := &models.Faculty{Name: "Computer Science"}

	type fields struct {
		fr faculty.IFacultyRepository
		lg *logger.LoggerZap
	}
	type args struct {
		ctx     context.Context
		faculty *models.Faculty
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		setup   func()
	}{
		// TODO: Add test cases.
		{
			name: "Success - Create Faculty",
			fields: fields{
				fr: mockRepo,
				lg: mockLogger,
			},
			args: args{
				ctx:     context.Background(),
				faculty: mockFaculties,
			},
			wantErr: false,
			setup: func() {
				mockRepo.EXPECT().CreateFaculty(gomock.Any(), mockFaculties).Return(nil)
			},
		},
		{
			name: "Failed - Create Faculty (Invalid Data)",
			fields: fields{
				fr: mockRepo,
				lg: mockLogger,
			},
			args: args{
				ctx:     context.Background(),
				faculty: mockFaculties,
			},
			wantErr: true,
			setup: func() {
				mockRepo.EXPECT().CreateFaculty(gomock.Any(), mockFaculties).Return(gorm.ErrInvalidData)
			},
		},
		{
			name: "Failed - Create Faculty (Duplicate Key)",
			fields: fields{
				fr: mockRepo,
				lg: mockLogger,
			},
			args: args{
				ctx:     context.Background(),
				faculty: mockFaculties,
			},
			wantErr: true,
			setup: func() {
				mockRepo.EXPECT().CreateFaculty(gomock.Any(), mockFaculties).Return(gorm.ErrDuplicatedKey)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			fu := &facultyUsecase{
				fr: tt.fields.fr,
				lg: tt.fields.lg,
			}
			if err := fu.CreateFaculty(tt.args.ctx, tt.args.faculty); (err != nil) != tt.wantErr {
				t.Errorf("facultyUsecase.CreateFaculty() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_facultyUsecase_DeleteFaculty(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIFacultyRepository(ctrl)
	mockLogger := logger.NewLoggerTest()
	mockFacultyID := 1

	type fields struct {
		fr faculty.IFacultyRepository
		lg *logger.LoggerZap
	}
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		setup   func()
	}{
		// TODO: Add test cases.
		{
			name: "Success - Delete Faculty",
			fields: fields{
				fr: mockRepo,
				lg: mockLogger,
			},
			args: args{
				ctx: context.Background(),
				id:  mockFacultyID,
			},
			wantErr: false,
			setup: func() {
				mockRepo.EXPECT().DeleteFaculty(gomock.Any(), mockFacultyID).Return(nil)
			},
		},
		{
			name: "Failed - Delete Faculty (Record Not Found)",
			fields: fields{
				fr: mockRepo,
				lg: mockLogger,
			},
			args: args{
				ctx: context.Background(),
				id:  mockFacultyID,
			},
			wantErr: true,
			setup: func() {
				mockRepo.EXPECT().DeleteFaculty(gomock.Any(), mockFacultyID).Return(gorm.ErrRecordNotFound)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			fu := &facultyUsecase{
				fr: tt.fields.fr,
				lg: tt.fields.lg,
			}
			if err := fu.DeleteFaculty(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("facultyUsecase.DeleteFaculty() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
