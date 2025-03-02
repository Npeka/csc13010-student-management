package repository

import (
	"context"
	"reflect"
	"testing"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/program"
	"github.com/csc13010-student-management/internal/program/mocks"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
)

func TestNewProgramRepository(t *testing.T) {
	t.Parallel()

	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want program.IProgramRepository
	}{
		{
			name: "Success - Create Program Repository",
			args: args{
				db: &gorm.DB{},
			},
			want: NewProgramRepository(&gorm.DB{}),
		},
		{
			name: "Failed - Create Program Repository",
			args: args{
				db: nil,
			},
			want: NewProgramRepository(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProgramRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProgramRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_programRepository_GetPrograms(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIProgramRepository(ctrl)
	mockPrograms := []*models.Program{
		{Name: "High Quality"},
		{Name: "Regular"},
	}

	type fields struct {
		pr program.IProgramRepository
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
			pr := tt.fields.pr
			got, err := pr.GetPrograms(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("programRepository.GetPrograms() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("programRepository.GetPrograms() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_programRepository_CreateProgram(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIProgramRepository(ctrl)
	mockProgram := &models.Program{Name: "High Quality"}

	type fields struct {
		pr program.IProgramRepository
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
			},
			args: args{
				ctx:     context.Background(),
				program: mockProgram,
			},
			wantErr: false,
			setup: func() {
				mockRepo.EXPECT().CreateProgram(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name: "Failed - Create Program (Invalid Data)",
			fields: fields{
				pr: mockRepo,
			},
			args: args{
				ctx:     context.Background(),
				program: mockProgram,
			},
			wantErr: true,
			setup: func() {
				mockRepo.EXPECT().CreateProgram(gomock.Any(), gomock.Any()).Return(gorm.ErrInvalidData)
			},
		},
		{
			name: "Failed - Create Program (Duplicate Key)",
			fields: fields{
				pr: mockRepo,
			},
			args: args{
				ctx:     context.Background(),
				program: mockProgram,
			},
			wantErr: true,
			setup: func() {
				mockRepo.EXPECT().CreateProgram(gomock.Any(), gomock.Any()).Return(gorm.ErrDuplicatedKey)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			pr := tt.fields.pr
			if err := pr.CreateProgram(tt.args.ctx, tt.args.program); (err != nil) != tt.wantErr {
				t.Errorf("programRepository.CreateProgram() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_programRepository_DeleteProgram(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIProgramRepository(ctrl)
	mockProgramID := uint(1)

	type fields struct {
		pr program.IProgramRepository
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
			name: "Success - Delete Program",
			fields: fields{
				pr: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				id:  mockProgramID,
			},
			wantErr: false,
			setup: func() {
				mockRepo.EXPECT().DeleteProgram(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name: "Failed - Delete Program (Record Not Found)",
			fields: fields{
				pr: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				id:  mockProgramID,
			},
			wantErr: true,
			setup: func() {
				mockRepo.EXPECT().DeleteProgram(gomock.Any(), gomock.Any()).Return(gorm.ErrRecordNotFound)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			pr := tt.fields.pr
			if err := pr.DeleteProgram(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("programRepository.DeleteProgram() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_programRepository_UpdateProgram(t *testing.T) {
	type fields struct {
		db *gorm.DB
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
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := &programRepository{
				db: tt.fields.db,
			}
			if err := pr.UpdateProgram(tt.args.ctx, tt.args.program); (err != nil) != tt.wantErr {
				t.Errorf("programRepository.UpdateProgram() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
