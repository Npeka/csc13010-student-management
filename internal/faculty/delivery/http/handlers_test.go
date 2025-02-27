package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/csc13010-student-management/internal/faculty"
	"github.com/csc13010-student-management/internal/faculty/mocks"
	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewFacultyHandlers(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUc := mocks.NewMockIFacultyUsecase(ctrl)
	mockLogger := logger.NewLoggerTest()

	type args struct {
		fu faculty.IFacultyUsecase
		lg *logger.LoggerZap
	}
	tests := []struct {
		name string
		args args
		want faculty.IFacultyHandlers
	}{
		// TODO: Add test cases.
		{
			name: "Success - Create Faculty Handlers",
			args: args{
				fu: mockUc,
				lg: mockLogger,
			},
			want: NewFacultyHandlers(mockUc, mockLogger),
		},
		{
			name: "Failed - Create Faculty Handlers",
			args: args{
				fu: nil,
				lg: nil,
			},
			want: NewFacultyHandlers(nil, nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFacultyHandlers(tt.args.fu, tt.args.lg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFacultyHandlers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_facultyHandlers_GetFaculties(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUc := mocks.NewMockIFacultyUsecase(ctrl)
	mockLogger := logger.NewLoggerTest()
	mockFaculties := []*models.Faculty{
		{ID: 1, Name: "Computer Science"},
		{ID: 2, Name: "Mathematics"},
	}

	fh := &facultyHandlers{
		fu: mockUc,
		lg: mockLogger,
	}

	r := gin.Default()
	r.GET("/faculties", fh.GetFaculties())

	tests := []struct {
		name           string
		mockBehavior   func()
		expectedStatus int
		expectedBody   string
	}{
		// TODO: Add test cases.
		{
			name: "Success - Get Faculties",
			mockBehavior: func() {
				mockUc.EXPECT().GetFaculties(gomock.Any()).Return(mockFaculties, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: func() string {
				bytes, _ := json.Marshal(mockFaculties)
				return string(bytes)
			}(),
		},
		{
			name: "Failed - Internal Server Error",
			mockBehavior: func() {
				mockUc.EXPECT().GetFaculties(gomock.Any()).Return(nil, errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"database error"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			req, _ := http.NewRequest(http.MethodGet, "/faculties", nil)
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}

func Test_facultyHandlers_CreateFaculty(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUc := mocks.NewMockIFacultyUsecase(ctrl)
	mockLogger := logger.NewLoggerTest()

	fh := &facultyHandlers{
		fu: mockUc,
		lg: mockLogger,
	}

	r := gin.Default()
	r.POST("/faculties", fh.CreateFaculty())

	tests := []struct {
		name           string
		inputBody      string
		mockBehavior   func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:      "Success - Create Faculty",
			inputBody: `{"name": "Computer Science"}`,
			mockBehavior: func() {
				mockUc.EXPECT().CreateFaculty(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"id":0,"name":"Computer Science"}`,
		},
		{
			name:           "Failed - Bad Request",
			inputBody:      `invalid json`,
			mockBehavior:   func() {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"invalid character 'i' looking for beginning of value"}`,
		},
		{
			name:      "Failed - Internal Server Error",
			inputBody: `{"name": "Computer Science"}`,
			mockBehavior: func() {
				mockUc.EXPECT().CreateFaculty(gomock.Any(), gomock.Any()).Return(errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"database error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			req, _ := http.NewRequest(http.MethodPost, "/faculties", strings.NewReader(tt.inputBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}

func Test_facultyHandlers_DeleteFaculty(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUc := mocks.NewMockIFacultyUsecase(ctrl)
	mockLogger := logger.NewLoggerTest()

	fh := &facultyHandlers{
		fu: mockUc,
		lg: mockLogger,
	}

	r := gin.Default()
	r.DELETE("/faculties/:faculty_id", fh.DeleteFaculty())

	tests := []struct {
		name           string
		facultyID      string
		mockBehavior   func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:      "Success - Delete Faculty",
			facultyID: "1",
			mockBehavior: func() {
				mockUc.EXPECT().DeleteFaculty(gomock.Any(), 1).Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"Faculty deleted successfully"}`,
		},
		{
			name:           "Failed - Bad Request",
			facultyID:      "invalid",
			mockBehavior:   func() {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Invalid faculty ID"}`,
		},
		{
			name:      "Failed - Internal Server Error",
			facultyID: "1",
			mockBehavior: func() {
				mockUc.EXPECT().DeleteFaculty(gomock.Any(), 1).Return(errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"database error"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			req, _ := http.NewRequest(http.MethodDelete, "/faculties/"+tt.facultyID, nil)
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}
