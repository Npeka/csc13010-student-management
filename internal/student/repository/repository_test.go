package repository

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/student/dtos"
	"github.com/csc13010-student-management/internal/student/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGetStudents(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIStudentRepository(ctrl)
	ctx := context.Background()
	students := []*models.Student{
		{
			ID:        1,
			StudentID: "22127180",
			FullName:  "Nguyen Phuc Khang",
			BirthDate: time.Date(2004, 8, 27, 0, 0, 0, 0, time.UTC),
			GenderID:  1,
			FacultyID: 1,
			CourseID:  1,
			ProgramID: 1,
			Address:   "HCM",
			Email:     "npkhang287@gmail.com",
			Phone:     "0123456789",
			StatusID:  1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			StudentID: "22127108",
			FullName:  "Huynh Yen Ngoc",
			BirthDate: time.Date(2004, 10, 19, 0, 0, 0, 0, time.UTC),
			GenderID:  2,
			FacultyID: 1,
			CourseID:  1,
			ProgramID: 1,
			Address:   "HCM",
			Email:     "huynhyenngoc@gmail.com",
			Phone:     "0123456789",
			StatusID:  1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	mockRepo.EXPECT().GetStudents(ctx).Return(students, nil)

	result, err := mockRepo.GetStudents(ctx)
	assert.NoError(t, err)
	assert.Equal(t, students, result)
}

func TestCreateStudent(t *testing.T) {
	t.Parallel()

	// Tạo mock database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Chuyển sqlmock thành Gorm DB
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	assert.NoError(t, err)

	// Tạo repository với mock DB
	studentRepo := NewStudentRepository(gormDB)

	student := models.Student{
		StudentID: "22127180",
		FullName:  "Nguyen Phuc Khang",
		BirthDate: time.Now(),
		GenderID:  1,
		FacultyID: 1,
		CourseID:  1,
		ProgramID: 1,
		Address:   "Ho Chi Minh City",
		Email:     "npkhang287@gmail.com",
		Phone:     "0123456789",
		StatusID:  1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Mock SQL query khi gọi CreateStudent
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "students"`).
		WithArgs(
			student.StudentID,
			student.FullName,
			student.BirthDate,
			student.GenderID,
			student.FacultyID,
			student.CourseID,
			student.ProgramID,
			student.Address,
			student.Email,
			student.Phone,
			student.StatusID,
			student.CreatedAt,
			student.UpdatedAt).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	mock.ExpectQuery(`INSERT INTO "audit_logs"`).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	mock.ExpectCommit()

	err = studentRepo.CreateStudent(context.Background(), &student)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
	assert.NotEqual(t, 0, student.ID)
}

func TestUpdateStudent(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIStudentRepository(ctrl)
	ctx := context.Background()
	student := &models.Student{
		FullName:  "Huynh Ngoc",
		BirthDate: time.Now(),
		GenderID:  1,
		FacultyID: 1,
		CourseID:  1,
		ProgramID: 1,
		Address:   "123 Main St",
		Email:     "john.doe@example.com",
		Phone:     "1234567890",
		StatusID:  1,
	}

	mockRepo.EXPECT().UpdateStudent(ctx, student).Return(nil)

	err := mockRepo.UpdateStudent(ctx, student)
	assert.NoError(t, err)
}

func TestDeleteStudent(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIStudentRepository(ctrl)
	ctx := context.Background()
	studentID := "22127180"

	mockRepo.EXPECT().DeleteStudent(ctx, studentID).Return(nil)

	err := mockRepo.DeleteStudent(ctx, studentID)
	assert.NoError(t, err)
}

func TestGetOptions(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIStudentRepository(ctrl)
	ctx := context.Background()
	options := &dtos.OptionDTO{
		Genders: []*dtos.Option{
			{ID: 1, Name: "Male"},
			{ID: 2, Name: "Female"},
		},
		Faculties: []*dtos.Option{
			{ID: 1, Name: "Engineering"},
			{ID: 2, Name: "Science"},
		},
		Courses: []*dtos.Option{
			{ID: 1, Name: "Computer Science"},
			{ID: 2, Name: "Mathematics"},
		},
		Programs: []*dtos.Option{
			{ID: 1, Name: "Undergraduate"},
			{ID: 2, Name: "Postgraduate"},
		},
		Statuses: []*dtos.Option{
			{ID: 1, Name: "Active"},
			{ID: 2, Name: "Inactive"},
		},
	}

	mockRepo.EXPECT().GetOptions(ctx).Return(options, nil)

	result, err := mockRepo.GetOptions(ctx)
	assert.NoError(t, err)
	assert.Equal(t, options, result)
}
