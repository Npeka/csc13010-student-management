package repository

import (
	"context"
	"reflect"
	"testing"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/status"
	"github.com/csc13010-student-management/internal/status/mocks"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
)

func TestNewStatusRepository(t *testing.T) {
	t.Parallel()

	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want status.IStatusRepository
	}{
		{
			name: "Success - Create Status Repository",
			args: args{
				db: &gorm.DB{},
			},
			want: NewStatusRepository(&gorm.DB{}),
		},
		{
			name: "Failed - Create Status Repository",
			args: args{
				db: nil,
			},
			want: NewStatusRepository(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStatusRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStatusRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_statusRepository_GetStatuses(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIStatusRepository(ctrl)
	mockStatuses := []*models.Status{
		{Name: "Active"},
		{Name: "Inactive"},
	}

	type fields struct {
		sr status.IStatusRepository
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
			sr := tt.fields.sr
			got, err := sr.GetStatuses(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("statusRepository.GetStatuses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("statusRepository.GetStatuses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_statusRepository_CreateStatus(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIStatusRepository(ctrl)
	mockStatus := &models.Status{Name: "Active"}

	type fields struct {
		sr status.IStatusRepository
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
			},
			args: args{
				ctx:    context.Background(),
				status: mockStatus,
			},
			wantErr: false,
			setup: func() {
				mockRepo.EXPECT().CreateStatus(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name: "Failed - Create Status (Invalid Data)",
			fields: fields{
				sr: mockRepo,
			},
			args: args{
				ctx:    context.Background(),
				status: mockStatus,
			},
			wantErr: true,
			setup: func() {
				mockRepo.EXPECT().CreateStatus(gomock.Any(), gomock.Any()).Return(gorm.ErrInvalidData)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			sr := tt.fields.sr
			if err := sr.CreateStatus(tt.args.ctx, tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("statusRepository.CreateStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_statusRepository_DeleteStatus(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIStatusRepository(ctrl)
	mockStatusID := uint(1)

	type fields struct {
		sr status.IStatusRepository
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
			},
			args: args{
				ctx: context.Background(),
				id:  mockStatusID,
			},
			wantErr: false,
			setup: func() {
				mockRepo.EXPECT().DeleteStatus(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name: "Failed - Delete Status (Record Not Found)",
			fields: fields{
				sr: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				id:  mockStatusID,
			},
			wantErr: true,
			setup: func() {
				mockRepo.EXPECT().DeleteStatus(gomock.Any(), gomock.Any()).Return(gorm.ErrRecordNotFound)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			sr := tt.fields.sr
			if err := sr.DeleteStatus(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("statusRepository.DeleteStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
