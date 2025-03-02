package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/student"
	"github.com/csc13010-student-management/internal/student/dtos"
	"github.com/csc13010-student-management/internal/student/mocks"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewStudentHandlers(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUc := mocks.NewMockIStudentUsecase(ctrl)
	mockLogger := logger.NewLoggerTest()

	type args struct {
		su student.IStudentUsecase
		lg *logger.LoggerZap
	}
	tests := []struct {
		name string
		args args
		want student.IStudentHandlers
	}{
		{
			name: "Success - Create Student Handlers",
			args: args{
				su: mockUc,
				lg: mockLogger,
			},
			want: NewStudentHandlers(mockUc, mockLogger),
		},
		{
			name: "Failed - Create Student Handlers",
			args: args{
				su: nil,
				lg: nil,
			},
			want: NewStudentHandlers(nil, nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStudentHandlers(tt.args.su, tt.args.lg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStudentHandlers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_studentHandlers_GetStudents(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUc := mocks.NewMockIStudentUsecase(ctrl)
	mockLogger := logger.NewLoggerTest()
	mockStudents := []*models.Student{
		{
			StudentID: "22127180",
			FullName:  "Nguyen Phuc Khang",
			BirthDate: "2004-8-27",
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

	sh := &studentHandlers{
		su: mockUc,
		lg: mockLogger,
	}

	r := gin.Default()
	r.GET("/students", sh.GetStudents())

	tests := []struct {
		name           string
		mockBehavior   func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Success - Get Students",
			mockBehavior: func() {
				mockUc.EXPECT().GetStudents(gomock.Any()).Return(mockStudents, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: func() string {
				bytes, _ := json.Marshal(mockStudents)
				return string(bytes)
			}(),
		},
		{
			name: "Failed - Internal Server Error",
			mockBehavior: func() {
				mockUc.EXPECT().GetStudents(gomock.Any()).Return(nil, errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"database error"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			req, _ := http.NewRequest(http.MethodGet, "/students", nil)
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}

func Test_studentHandlers_GetStudentByStudentID(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUc := mocks.NewMockIStudentUsecase(ctrl)
	mockLogger := logger.NewLoggerTest()
	mockStudent := &models.Student{
		StudentID: "22127180",
		FullName:  "Nguyen Phuc Khang",
		BirthDate: "2004-8-27",
		GenderID:  1,
		FacultyID: 1,
		CourseID:  1,
		ProgramID: 2,
		Address:   "HCM",
		Email:     "npkhang287@student.university.edu.vn",
		Phone:     "0789123456",
		StatusID:  1,
	}

	sh := &studentHandlers{
		su: mockUc,
		lg: mockLogger,
	}

	r := gin.Default()
	r.GET("/students/:student_id", sh.GetStudentByStudentID())

	tests := []struct {
		name           string
		studentID      string
		mockBehavior   func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:      "Success - Get Student By ID",
			studentID: "22127180",
			mockBehavior: func() {
				mockUc.EXPECT().GetStudentByStudentID(gomock.Any(), "22127180").Return(mockStudent, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: func() string {
				bytes, _ := json.Marshal(mockStudent)
				return string(bytes)
			}(),
		},
		{
			name:      "Failed - Student Not Found",
			studentID: "99999999",
			mockBehavior: func() {
				mockUc.EXPECT().GetStudentByStudentID(gomock.Any(), "99999999").Return(nil, errors.New("student not found"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"student not found"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			req, _ := http.NewRequest(http.MethodGet, "/students/"+tt.studentID, nil)
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}

func Test_studentHandlers_GetFullInfoStudentByStudentID(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUc := mocks.NewMockIStudentUsecase(ctrl)
	mockLogger := logger.NewLoggerTest()
	mockStudent := &dtos.StudentDTO{
		StudentID: "22127180",
		FullName:  "Nguyen Phuc Khang",
		BirthDate: "2004-8-27",
		Gender:    "Male",
		Faculty:   1,
		Course:    1,
		Program:   2,
		Address:   "HCM",
		Email:     "npkhang287@student.university.edu.vn",
		Phone:     "0789123456",
		Status:    1,
	}

	sh := &studentHandlers{
		su: mockUc,
		lg: mockLogger,
	}

	r := gin.Default()
	r.GET("/students/full/:student_id", sh.GetFullInfoStudentByStudentID())

	tests := []struct {
		name           string
		studentID      string
		mockBehavior   func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:      "Success - Get Full Info Student By ID",
			studentID: "22127180",
			mockBehavior: func() {
				mockUc.EXPECT().GetFullInfoStudentByStudentID(gomock.Any(), "22127180").Return(mockStudent, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: func() string {
				bytes, _ := json.Marshal(mockStudent)
				return string(bytes)
			}(),
		},
		{
			name:      "Failed - Student Not Found",
			studentID: "99999999",
			mockBehavior: func() {
				mockUc.EXPECT().GetFullInfoStudentByStudentID(gomock.Any(), "99999999").Return(nil, errors.New("student not found"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"student not found"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			req, _ := http.NewRequest(http.MethodGet, "/students/full/"+tt.studentID, nil)
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}

func Test_studentHandlers_CreateStudent(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUc := mocks.NewMockIStudentUsecase(ctrl)
	mockLogger := logger.NewLoggerTest()

	// Define mock student data
	mockStudent := &models.Student{
		StudentID: "22127180",
		FullName:  "Nguyen Phuc Khang",
		BirthDate: "2004-8-27",
		GenderID:  1,
		FacultyID: 1,
		CourseID:  1,
		ProgramID: 2,
		Address:   "HCM",
		Email:     "npkhang287@student.university.edu.vn",
		Phone:     "0789123456",
		StatusID:  1,
	}

	sh := &studentHandlers{
		su: mockUc,
		lg: mockLogger,
	}

	r := gin.Default()
	r.POST("/students", sh.CreateStudent())

	tests := []struct {
		name           string
		inputBody      string
		mockBehavior   func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Success - Create Student",
			inputBody: func() string {
				bytes, _ := json.Marshal(mockStudent)
				return string(bytes)
			}(),
			mockBehavior: func() {
				mockUc.EXPECT().CreateStudent(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedStatus: http.StatusCreated,
			expectedBody: func() string {
				bytes, _ := json.Marshal(mockStudent)
				return string(bytes)
			}(),
		},
		{
			name:           "Failed - Invalid Input",
			inputBody:      `invalid data`,
			mockBehavior:   func() {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"invalid character 'i' looking for beginning of value"}`,
		},
		{
			name: "Failed - Internal Server Error",
			inputBody: func() string {
				bytes, _ := json.Marshal(mockStudent)
				return string(bytes)
			}(),
			mockBehavior: func() {
				mockUc.EXPECT().CreateStudent(gomock.Any(), gomock.Any()).Return(errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"database error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call mock behavior setup
			tt.mockBehavior()

			req, _ := http.NewRequest(http.MethodPost, "/students", bytes.NewBufferString(tt.inputBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}

func Test_studentHandlers_UpdateStudent(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUc := mocks.NewMockIStudentUsecase(ctrl)
	mockLogger := logger.NewLoggerTest()
	mockStudent := &models.Student{
		StudentID: "22127180",
		FullName:  "Nguyen Phuc Khang",
		BirthDate: "2004-8-27",
		GenderID:  1,
		FacultyID: 1,
		CourseID:  1,
		ProgramID: 2,
		Address:   "HCM",
		Email:     "npkhang287@student.university.edu.vn",
		Phone:     "0789123456",
		StatusID:  1,
	}

	sh := &studentHandlers{
		su: mockUc,
		lg: mockLogger,
	}

	r := gin.Default()
	r.PATCH("/students/:student_id", sh.UpdateStudent())

	tests := []struct {
		name           string
		studentID      string
		inputBody      string
		mockBehavior   func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:      "Success - Update Student",
			studentID: "22127180",
			inputBody: func() string {
				bytes, _ := json.Marshal(mockStudent)
				return string(bytes)
			}(),
			mockBehavior: func() {
				mockUc.EXPECT().UpdateStudent(gomock.Any(), mockStudent).Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: func() string {
				bytes, _ := json.Marshal(mockStudent)
				return string(bytes)
			}(),
		},
		{
			name:           "Failed - Invalid Student ID",
			studentID:      "invalid_id",
			inputBody:      `{"invalid": "data"}`,
			mockBehavior:   func() {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Invalid student ID"}`,
		},
		{
			name:      "Failed - Internal Server Error",
			studentID: "22127180",
			inputBody: func() string {
				bytes, _ := json.Marshal(mockStudent)
				return string(bytes)
			}(),
			mockBehavior: func() {
				mockUc.EXPECT().UpdateStudent(gomock.Any(), mockStudent).Return(errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"database error"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			req, _ := http.NewRequest(http.MethodPatch, "/students/"+tt.studentID, bytes.NewBufferString(tt.inputBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}

func Test_studentHandlers_DeleteStudent(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUc := mocks.NewMockIStudentUsecase(ctrl)
	mockLogger := logger.NewLoggerTest()

	sh := &studentHandlers{
		su: mockUc,
		lg: mockLogger,
	}

	r := gin.Default()
	r.DELETE("/students/:student_id", sh.DeleteStudent())

	tests := []struct {
		name           string
		studentID      string
		mockBehavior   func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:      "Success - Delete Student",
			studentID: "22127180",
			mockBehavior: func() {
				mockUc.EXPECT().DeleteStudent(gomock.Any(), "22127180").Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"Student deleted successfully"}`,
		},
		{
			name:      "Failed - Student Not Found",
			studentID: "99999999",
			mockBehavior: func() {
				mockUc.EXPECT().DeleteStudent(gomock.Any(), "99999999").Return(errors.New("student not found"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"student not found"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			req, _ := http.NewRequest(http.MethodDelete, "/students/"+tt.studentID, nil)
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}

func Test_studentHandlers_GetOptions(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUc := mocks.NewMockIStudentUsecase(ctrl)
	mockLogger := logger.NewLoggerTest()
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
			{ID: 1, Name: "Computer Science"},
			{ID: 2, Name: "Mechanical Engineering"},
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

	sh := &studentHandlers{
		su: mockUc,
		lg: mockLogger,
	}

	r := gin.Default()
	r.GET("/options", sh.GetOptions())

	tests := []struct {
		name           string
		mockBehavior   func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Success - Get Options",
			mockBehavior: func() {
				mockUc.EXPECT().GetOptions(gomock.Any()).Return(mockOptions, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: func() string {
				bytes, _ := json.Marshal(mockOptions)
				return string(bytes)
			}(),
		},
		{
			name: "Failed - Internal Server Error",
			mockBehavior: func() {
				mockUc.EXPECT().GetOptions(gomock.Any()).Return(nil, errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"database error"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			req, _ := http.NewRequest(http.MethodGet, "/options", nil)
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}
