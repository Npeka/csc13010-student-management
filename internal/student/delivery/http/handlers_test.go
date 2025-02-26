package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/student/dtos"
	"github.com/csc13010-student-management/internal/student/mocks"
	"github.com/csc13010-student-management/internal/student/usecase"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetStudents(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLoggerTest()
	mockRepo := mocks.NewMockIStudentRepository(ctrl)
	uc := usecase.NewStudentUsecase(mockRepo, logger)
	handler := NewStudentHandlers(uc, logger)

	router := gin.Default()
	router.GET("/students", handler.GetStudents())

	mockRepo.EXPECT().GetStudents(gomock.Any()).Return([]*models.Student{}, nil)

	req, _ := http.NewRequest(http.MethodGet, "/students", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestCreateStudent(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLoggerTest()
	mockRepo := mocks.NewMockIStudentRepository(ctrl)
	uc := usecase.NewStudentUsecase(mockRepo, logger)
	handler := NewStudentHandlers(uc, logger)

	router := gin.Default()
	router.POST("/students", handler.CreateStudent())

	fixedTime := time.Date(2025, 2, 26, 9, 20, 39, 0, time.UTC)
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
		CreatedAt: fixedTime,
		UpdatedAt: fixedTime,
	}

	mockRepo.EXPECT().CreateStudent(gomock.Any(), &student).Return(nil)

	jsonValue, _ := json.Marshal(student)
	req, _ := http.NewRequest(http.MethodPost, "/students", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestUpdateStudent(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLoggerTest()
	mockRepo := mocks.NewMockIStudentRepository(ctrl)
	uc := usecase.NewStudentUsecase(mockRepo, logger)
	handler := NewStudentHandlers(uc, logger)

	router := gin.Default()
	router.PATCH("/students/:student_id", handler.UpdateStudent())

	fixedTime := time.Date(2025, 2, 26, 9, 20, 39, 0, time.UTC)
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
	}
	mockRepo.EXPECT().UpdateStudent(gomock.Any(), &student).Return(nil)

	jsonValue, _ := json.Marshal(student)
	req, _ := http.NewRequest(http.MethodPatch, "/students/22127180", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestDeleteStudent(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLoggerTest()
	mockRepo := mocks.NewMockIStudentRepository(ctrl)
	uc := usecase.NewStudentUsecase(mockRepo, logger)
	handler := NewStudentHandlers(uc, logger)

	router := gin.Default()
	router.DELETE("/students/:student_id", handler.DeleteStudent())

	mockRepo.EXPECT().DeleteStudent(gomock.Any(), "22127180").Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, "/students/22127180", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetOptions(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger.NewLoggerTest()
	mockRepo := mocks.NewMockIStudentRepository(ctrl)
	uc := usecase.NewStudentUsecase(mockRepo, logger)
	handler := NewStudentHandlers(uc, logger)

	router := gin.Default()
	router.GET("/students/options", handler.GetOptions())

	mockRepo.EXPECT().GetOptions(gomock.Any()).Return(&dtos.OptionDTO{}, nil)

	req, _ := http.NewRequest(http.MethodGet, "/students/options", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}
