package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/program"
	"github.com/csc13010-student-management/internal/program/mocks"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
)

func TestNewProgramUsecase(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIProgramRepository(ctrl)
	mockLogger := logger.NewLoggerTest()

	type args struct {
		pr program.IProgramRepository
		lg *logger.LoggerZap
	}
	tests := []struct {
		name string
		args args
		want program.IProgramUsecase
	}{
		{
			name: "Success - Create Program Usecase",
			args: args{
				pr: mockRepo,
				lg: mockLogger,
			},
			want: NewProgramUsecase(mockRepo, mockLogger),
		},
		{
			name: "Failed - Create Program Usecase",
			args: args{
				pr: nil,
				lg: nil,
			},
			want: NewProgramUsecase(nil, nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProgramUsecase(tt.args.pr, tt.args.lg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProgramUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_programUsecase_GetPrograms(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIProgramRepository(ctrl)
	mockLogger := logger.NewLoggerTest()
	mockPrograms := []*models.Program{
		{Name: "Computer Science"},
		{Name: "Mathematics"},
	}

	type fields struct {
		pr program.IProgramRepository
		lg *logger.LoggerZap
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.Program
		wantErr bool
		setup   func()
	}{
		{
			name: "Success - Get Programs",
			fields: fields{
				pr: mockRepo,
				lg: mockLogger,
			},
			args: args{
				ctx: context.Background(),
			},
			want:    mockPrograms,
			wantErr: false,
			setup: func() {
				mockRepo.EXPECT().GetPrograms(gomock.Any()).Return(mockPrograms, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			pu := &programUsecase{
				pr: tt.fields.pr,
				lg: tt.fields.lg,
			}
			got, err := pu.GetPrograms(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("programUsecase.GetPrograms() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("programUsecase.GetPrograms() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_programUsecase_CreateProgram(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIProgramRepository(ctrl)
	mockLogger := logger.NewLoggerTest()
	mockProgram := &models.Program{Name: "Computer Science"}

	type fields struct {
		pr program.IProgramRepository
		lg *logger.LoggerZap
	}
	type args struct {
		ctx     context.Context
		program *models.Program
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		setup   func()
	}{
		{
			name: "Success - Create Program",
			fields: fields{
				pr: mockRepo,
				lg: mockLogger,
			},
			args: args{
				ctx:     context.Background(),
				program: mockProgram,
			},
			wantErr: false,
			setup: func() {
				mockRepo.EXPECT().CreateProgram(gomock.Any(), mockProgram).Return(nil)
			},
		},
		{
			name: "Failed - Create Program (Invalid Data)",
			fields: fields{
				pr: mockRepo,
				lg: mockLogger,
			},
			args: args{
				ctx:     context.Background(),
				program: mockProgram,
			},
			wantErr: true,
			setup: func() {
				mockRepo.EXPECT().CreateProgram(gomock.Any(), mockProgram).Return(gorm.ErrInvalidData)
			},
		},
		{
			name: "Failed - Create Program (Duplicate Key)",
			fields: fields{
				pr: mockRepo,
				lg: mockLogger,
			},
			args: args{
				ctx:     context.Background(),
				program: mockProgram,
			},
			wantErr: true,
			setup: func() {
				mockRepo.EXPECT().CreateProgram(gomock.Any(), mockProgram).Return(gorm.ErrDuplicatedKey)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			pu := &programUsecase{
				pr: tt.fields.pr,
				lg: tt.fields.lg,
			}
			if err := pu.CreateProgram(tt.args.ctx, tt.args.program); (err != nil) != tt.wantErr {
				t.Errorf("programUsecase.CreateProgram() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_programUsecase_DeleteProgram(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIProgramRepository(ctrl)
	mockLogger := logger.NewLoggerTest()
	mockProgramID := 1

	type fields struct {
		pr program.IProgramRepository
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
		{
			name: "Success - Delete Program",
			fields: fields{
				pr: mockRepo,
				lg: mockLogger,
			},
			args: args{
				ctx: context.Background(),
				id:  mockProgramID,
			},
			wantErr: false,
			setup: func() {
				mockRepo.EXPECT().DeleteProgram(gomock.Any(), mockProgramID).Return(nil)
			},
		},
		{
			name: "Failed - Delete Program (Record Not Found)",
			fields: fields{
				pr: mockRepo,
				lg: mockLogger,
			},
			args: args{
				ctx: context.Background(),
				id:  mockProgramID,
			},
			wantErr: true,
			setup: func() {
				mockRepo.EXPECT().DeleteProgram(gomock.Any(), mockProgramID).Return(gorm.ErrRecordNotFound)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			pu := &programUsecase{
				pr: tt.fields.pr,
				lg: tt.fields.lg,
			}
			if err := pu.DeleteProgram(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("programUsecase.DeleteProgram() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
