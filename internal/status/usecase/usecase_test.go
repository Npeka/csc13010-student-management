package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/status"
	"github.com/csc13010-student-management/internal/status/mocks"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
)

func TestNewStatusUsecase(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIStatusRepository(ctrl)
	mockLogger := logger.NewLoggerTest()

	type args struct {
		sr status.IStatusRepository
		lg *logger.LoggerZap
	}
	tests := []struct {
		name string
		args args
		want status.IStatusUsecase
	}{
		{
			name: "Success - Create Status Usecase",
			args: args{
				sr: mockRepo,
				lg: mockLogger,
			},
			want: NewStatusUsecase(mockRepo, mockLogger),
		},
		{
			name: "Failed - Create Status Usecase",
			args: args{
				sr: nil,
				lg: nil,
			},
			want: NewStatusUsecase(nil, nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStatusUsecase(tt.args.sr, tt.args.lg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStatusUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_statusUsecase_GetStatuses(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIStatusRepository(ctrl)
	mockLogger := logger.NewLoggerTest()
	mockStatuses := []*models.Status{
		{ID: 1, Name: "Active"},
		{ID: 2, Name: "Inactive"},
	}

	type fields struct {
		sr status.IStatusRepository
		lg *logger.LoggerZap
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.Status
		wantErr bool
		setup   func()
	}{
		{
			name: "Success - Get Statuses",
			fields: fields{
				sr: mockRepo,
				lg: mockLogger,
			},
			args: args{
				ctx: context.Background(),
			},
			want:    mockStatuses,
			wantErr: false,
			setup: func() {
				mockRepo.EXPECT().GetStatuses(gomock.Any()).Return(mockStatuses, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			su := &statusUsecase{
				sr: tt.fields.sr,
				lg: tt.fields.lg,
			}
			got, err := su.GetStatuses(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("statusUsecase.GetStatuses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("statusUsecase.GetStatuses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_statusUsecase_CreateStatus(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIStatusRepository(ctrl)
	mockLogger := logger.NewLoggerTest()
	mockStatus := &models.Status{Name: "Active"}

	type fields struct {
		sr status.IStatusRepository
		lg *logger.LoggerZap
	}
	type args struct {
		ctx    context.Context
		status *models.Status
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		setup   func()
	}{
		{
			name: "Success - Create Status",
			fields: fields{
				sr: mockRepo,
				lg: mockLogger,
			},
			args: args{
				ctx:    context.Background(),
				status: mockStatus,
			},
			wantErr: false,
			setup: func() {
				mockRepo.EXPECT().CreateStatus(gomock.Any(), mockStatus).Return(nil)
			},
		},
		{
			name: "Failed - Create Status (Invalid Data)",
			fields: fields{
				sr: mockRepo,
				lg: mockLogger,
			},
			args: args{
				ctx:    context.Background(),
				status: mockStatus,
			},
			wantErr: true,
			setup: func() {
				mockRepo.EXPECT().CreateStatus(gomock.Any(), mockStatus).Return(gorm.ErrInvalidData)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			su := &statusUsecase{
				sr: tt.fields.sr,
				lg: tt.fields.lg,
			}
			if err := su.CreateStatus(tt.args.ctx, tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("statusUsecase.CreateStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_statusUsecase_DeleteStatus(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIStatusRepository(ctrl)
	mockLogger := logger.NewLoggerTest()
	mockStatusID := uint(1)

	type fields struct {
		sr status.IStatusRepository
		lg *logger.LoggerZap
	}
	type args struct {
		ctx context.Context
		id  uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		setup   func()
	}{
		{
			name: "Success - Delete Status",
			fields: fields{
				sr: mockRepo,
				lg: mockLogger,
			},
			args: args{
				ctx: context.Background(),
				id:  mockStatusID,
			},
			wantErr: false,
			setup: func() {
				mockRepo.EXPECT().DeleteStatus(gomock.Any(), mockStatusID).Return(nil)
			},
		},
		{
			name: "Failed - Delete Status (Record Not Found)",
			fields: fields{
				sr: mockRepo,
				lg: mockLogger,
			},
			args: args{
				ctx: context.Background(),
				id:  mockStatusID,
			},
			wantErr: true,
			setup: func() {
				mockRepo.EXPECT().DeleteStatus(gomock.Any(), mockStatusID).Return(gorm.ErrRecordNotFound)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			su := &statusUsecase{
				sr: tt.fields.sr,
				lg: tt.fields.lg,
			}
			if err := su.DeleteStatus(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("statusUsecase.DeleteStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
