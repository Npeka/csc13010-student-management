package repository

import (
	"context"
	"reflect"
	"testing"

	"github.com/csc13010-student-management/internal/faculty"
	"github.com/csc13010-student-management/internal/faculty/mocks"
	"github.com/csc13010-student-management/internal/models"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
)

func TestNewFacultyRepository(t *testing.T) {
	t.Parallel()
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want faculty.IFacultyRepository
	}{
		// TODO: Add test cases.
		{
			name: "Success - Create Faculty Repository",
			args: args{
				db: &gorm.DB{},
			},
			want: NewFacultyRepository(&gorm.DB{}),
		},
		{
			name: "Failed - Create Faculty Repository",
			args: args{
				db: nil,
			},
			want: NewFacultyRepository(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFacultyRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFacultyRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_facultyRepository_GetFaculties(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIFacultyRepository(ctrl)
	mockFaculties := []*models.Faculty{
		{Name: "Computer Science"},
		{Name: "Mathematics"},
	}

	type fields struct {
		fr faculty.IFacultyRepository
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
			fr := tt.fields.fr
			got, err := fr.GetFaculties(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("facultyRepository.GetFaculties() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("facultyRepository.GetFaculties() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_facultyRepository_CreateFaculty(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIFacultyRepository(ctrl)
	mockFaculties := &models.Faculty{Name: "Computer Science"}

	type fields struct {
		fr faculty.IFacultyRepository
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
			},
			args: args{
				ctx:     context.Background(),
				faculty: mockFaculties,
			},
			wantErr: false,
			setup: func() {
				mockRepo.EXPECT().CreateFaculty(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name: "Failed - Create Faculty (Invalid Data)",
			fields: fields{
				fr: mockRepo,
			},
			args: args{
				ctx:     context.Background(),
				faculty: mockFaculties,
			},
			wantErr: true,
			setup: func() {
				mockRepo.EXPECT().CreateFaculty(gomock.Any(), gomock.Any()).Return(gorm.ErrInvalidData)
			},
		},
		{
			name: "Failed - Create Faculty (Duplicate Key)",
			fields: fields{
				fr: mockRepo,
			},
			args: args{
				ctx:     context.Background(),
				faculty: mockFaculties,
			},
			wantErr: true,
			setup: func() {
				mockRepo.EXPECT().CreateFaculty(gomock.Any(), gomock.Any()).Return(gorm.ErrDuplicatedKey)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			fr := tt.fields.fr
			if err := fr.CreateFaculty(tt.args.ctx, tt.args.faculty); (err != nil) != tt.wantErr {
				t.Errorf("facultyRepository.CreateFaculty() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_facultyRepository_DeleteFaculty(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIFacultyRepository(ctrl)
	mockFacultyID := uint(1)

	type fields struct {
		fr faculty.IFacultyRepository
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
		// TODO: Add test cases.
		{
			name: "Success - Delete Faculty",
			fields: fields{
				fr: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				id:  mockFacultyID,
			},
			wantErr: false,
			setup: func() {
				mockRepo.EXPECT().DeleteFaculty(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name: "Failed - Delete Faculty (Record Not Found)",
			fields: fields{
				fr: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				id:  mockFacultyID,
			},
			wantErr: true,
			setup: func() {
				mockRepo.EXPECT().DeleteFaculty(gomock.Any(), gomock.Any()).Return(gorm.ErrRecordNotFound)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			fr := tt.fields.fr
			if err := fr.DeleteFaculty(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("facultyRepository.DeleteFaculty() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
