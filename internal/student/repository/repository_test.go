package repository

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/student"
	"github.com/csc13010-student-management/internal/student/dtos"
	"github.com/csc13010-student-management/internal/student/mocks"
	"github.com/golang/mock/gomock"
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
		{
			name: "Success - Create Student Repository",
			args: args{
				db: &gorm.DB{},
			},
			want: NewStudentRepository(&gorm.DB{}),
		},
		{
			name: "Failed - Create Student Repository",
			args: args{
				db: nil,
			},
			want: NewStudentRepository(nil),
		},
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
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIStudentRepository(ctrl)
	mockStudents := []*models.Student{
		{
			StudentID: "22127180",
			FullName:  "Nguyen Phuc Khang",
			BirthDate: time.Date(2004, 8, 27, 0, 0, 0, 0, time.UTC),
			GenderID:  1,
			FacultyID: 1,
			CourseID:  1,
			ProgramID: 2,
			Address:   "HCM",
			Email:     "npkhang287@student.university.edu.vn",
			Phone:     "0789123456",
			StatusID:  1,
		},
	}

	type fields struct {
		sr student.IStudentRepository
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
		setup   func()
	}{
		{
			name: "Success - Get Students",
			fields: fields{
				sr: mockRepo,
			},
			args: args{
				ctx: context.Background(),
			},
			want:    mockStudents,
			wantErr: false,
			setup: func() {
				mockRepo.EXPECT().GetStudents(gomock.Any()).Return(mockStudents, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			sr := tt.fields.sr
			got, err := sr.GetStudents(tt.args.ctx)
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

func Test_studentRepository_CreateStudent(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIStudentRepository(ctrl)
	mockStudent := &models.Student{
		StudentID: "22127180",
		FullName:  "Nguyen Phuc Khang",
		BirthDate: time.Date(2004, 8, 27, 0, 0, 0, 0, time.UTC),
		GenderID:  1,
		FacultyID: 1,
		CourseID:  1,
		ProgramID: 2,
		Address:   "HCM",
		Email:     "npkhang287@student.university.edu.vn",
		Phone:     "0789123456",
		StatusID:  1,
	}

	type fields struct {
		sr student.IStudentRepository
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			sr := tt.fields.sr
			if err := sr.CreateStudent(tt.args.ctx, tt.args.student); (err != nil) != tt.wantErr {
				t.Errorf("studentRepository.CreateStudent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_studentRepository_UpdateStudent(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIStudentRepository(ctrl)
	mockStudent := &models.Student{
		StudentID: "22127180",
		FullName:  "Nguyen Phuc Khang",
		BirthDate: time.Date(2004, 8, 27, 0, 0, 0, 0, time.UTC),
		GenderID:  1,
		FacultyID: 1,
		CourseID:  1,
		ProgramID: 2,
		Address:   "HCM",
		Email:     "npkhang287@student.university.edu.vn",
		Phone:     "0789123456",
		StatusID:  1,
	}

	type fields struct {
		sr student.IStudentRepository
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			sr := tt.fields.sr
			if err := sr.UpdateStudent(tt.args.ctx, tt.args.student); (err != nil) != tt.wantErr {
				t.Errorf("studentRepository.UpdateStudent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_studentRepository_DeleteStudent(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIStudentRepository(ctrl)

	type fields struct {
		sr student.IStudentRepository
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			sr := tt.fields.sr
			if err := sr.DeleteStudent(tt.args.ctx, tt.args.studentID); (err != nil) != tt.wantErr {
				t.Errorf("studentRepository.DeleteStudent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_studentRepository_GetFullInfoStudentByStudentID(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIStudentRepository(ctrl)
	mockStudentDTO := &dtos.StudentDTO{
		StudentID: "22127180",
		FullName:  "Nguyen Phuc Khang",
		BirthDate: time.Date(2004, 8, 27, 0, 0, 0, 0, time.UTC),
		Gender:    "Male",
		Faculty:   1,
		Course:    1,
		Program:   2,
		Address:   "HCM",
		Email:     "npkhang287@student.university.edu.vn",
		Phone:     "0789123456",
		Status:    1,
	}

	type fields struct {
		sr student.IStudentRepository
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
		setup   func()
	}{
		{
			name: "Success - Get Full Info Student By Student ID",
			fields: fields{
				sr: mockRepo,
			},
			args: args{
				ctx:       context.Background(),
				studentID: "22127180",
			},
			want:    mockStudentDTO,
			wantErr: false,
			setup: func() {
				mockRepo.EXPECT().GetFullInfoStudentByStudentID(gomock.Any(), "22127180").Return(mockStudentDTO, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			sr := tt.fields.sr
			got, err := sr.GetFullInfoStudentByStudentID(tt.args.ctx, tt.args.studentID)
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

func Test_studentRepository_GetOptions(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIStudentRepository(ctrl)
	mockOptions := &dtos.OptionDTO{
		Genders: []*dtos.Option{
			{ID: 1, Name: "Male"},
			{ID: 2, Name: "Female"},
		},
		Faculties: []*dtos.Option{
			{ID: 1, Name: "Engineering"},
			{ID: 2, Name: "Science"},
		},
		Courses: []*dtos.Option{
			{ID: 1, Name: "Course1"},
			{ID: 2, Name: "Course2"},
		},
		Programs: []*dtos.Option{
			{ID: 1, Name: "Program1"},
			{ID: 2, Name: "Program2"},
		},
		Statuses: []*dtos.Option{
			{ID: 1, Name: "Active"},
			{ID: 2, Name: "Inactive"},
		},
	}

	type fields struct {
		sr student.IStudentRepository
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
		setup   func()
	}{
		{
			name: "Success - Get Options",
			fields: fields{
				sr: mockRepo,
			},
			args: args{
				ctx: context.Background(),
			},
			want:    mockOptions,
			wantErr: false,
			setup: func() {
				mockRepo.EXPECT().GetOptions(gomock.Any()).Return(mockOptions, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			sr := tt.fields.sr
			got, err := sr.GetOptions(tt.args.ctx)
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

func Test_studentRepository_UpdateUserIDByUsername(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx       context.Context
		studentID string
		userID    uint
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
			if err := s.UpdateUserIDByUsername(tt.args.ctx, tt.args.studentID, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("studentRepository.UpdateUserIDByUsername() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
