package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/student/dtos"
	"github.com/csc13010-student-management/internal/student/mocks"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetStudents(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLoggerTest()
	mockRepo := mocks.NewMockIStudentRepository(ctrl)
	stUsecase := NewStudentUsecase(mockRepo, logger)

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

	students, err := stUsecase.GetStudents(ctx)

	assert.NoError(t, err)
	assert.Equal(t, students, students)
}

func TestCreateStudent(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLoggerTest()
	mockRepo := mocks.NewMockIStudentRepository(ctrl)
	stUsecase := NewStudentUsecase(mockRepo, logger)

	fixedTime := time.Date(2025, 2, 26, 9, 20, 39, 0, time.UTC)
	ctx := context.Background()
	student := models.Student{
		StudentID: "22127180",
		FullName:  "Nguyen Phuc Khang",
		BirthDate: fixedTime,
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

	mockRepo.EXPECT().CreateStudent(ctx, gomock.Eq(&student)).Return(nil)

	err := stUsecase.CreateStudent(ctx, &student)

	assert.NoError(t, err)
}

func TestDeleteStudent(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLoggerTest()
	mockRepo := mocks.NewMockIStudentRepository(ctrl)
	stUsecase := NewStudentUsecase(mockRepo, logger)

	ctx := context.Background()
	studentID := "22127180"

	mockRepo.EXPECT().DeleteStudent(ctx, studentID).Return(nil)

	err := stUsecase.DeleteStudent(ctx, studentID)

	assert.NoError(t, err)
}

func TestUpdateStudent(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLoggerTest()
	mockRepo := mocks.NewMockIStudentRepository(ctrl)
	stUsecase := NewStudentUsecase(mockRepo, logger)

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

	err := stUsecase.UpdateStudent(ctx, student)

	assert.NoError(t, err)
}

func TestGetOptions(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLoggerTest()
	mockRepo := mocks.NewMockIStudentRepository(ctrl)
	stUsecase := NewStudentUsecase(mockRepo, logger)

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

	result, err := stUsecase.GetOptions(ctx)

	assert.NoError(t, err)
	assert.Equal(t, options, result)
}
