package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/program"
	"github.com/csc13010-student-management/internal/program/mocks"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewProgramHandlers(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUc := mocks.NewMockIProgramUsecase(ctrl)
	mockLogger := logger.NewLoggerTest()

	type args struct {
		pu program.IProgramUsecase
		lg *logger.LoggerZap
	}
	tests := []struct {
		name string
		args args
		want program.IProgramHandlers
	}{
		{
			name: "Success - Create Program Handlers",
			args: args{
				pu: mockUc,
				lg: mockLogger,
			},
			want: NewProgramHandlers(mockUc, mockLogger),
		},
		{
			name: "Failed - Create Program Handlers",
			args: args{
				pu: nil,
				lg: nil,
			},
			want: NewProgramHandlers(nil, nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProgramHandlers(tt.args.pu, tt.args.lg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProgramHandlers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_programHandlers_GetPrograms(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUc := mocks.NewMockIProgramUsecase(ctrl)
	mockLogger := logger.NewLoggerTest()
	mockPrograms := []*models.Program{
		{ID: 1, Name: "Program 1"},
		{ID: 2, Name: "Program 2"},
	}

	ph := &programHandlers{
		pu: mockUc,
		lg: mockLogger,
	}

	r := gin.Default()
	r.GET("/programs", ph.GetPrograms())

	tests := []struct {
		name           string
		mockBehavior   func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Success - Get Programs",
			mockBehavior: func() {
				mockUc.EXPECT().GetPrograms(gomock.Any()).Return(mockPrograms, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: func() string {
				bytes, _ := json.Marshal(mockPrograms)
				return string(bytes)
			}(),
		},
		{
			name: "Failed - Internal Server Error",
			mockBehavior: func() {
				mockUc.EXPECT().GetPrograms(gomock.Any()).Return(nil, errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"database error"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			req, _ := http.NewRequest(http.MethodGet, "/programs", nil)
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}

func Test_programHandlers_CreateProgram(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUc := mocks.NewMockIProgramUsecase(ctrl)
	mockLogger := logger.NewLoggerTest()

	ph := &programHandlers{
		pu: mockUc,
		lg: mockLogger,
	}

	r := gin.Default()
	r.POST("/programs", ph.CreateProgram())

	tests := []struct {
		name           string
		inputBody      string
		mockBehavior   func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:      "Success - Create Program",
			inputBody: `{"name": "Program 1"}`,
			mockBehavior: func() {
				mockUc.EXPECT().CreateProgram(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"id":0,"name":"Program 1"}`,
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
			inputBody: `{"name": "Program 1"}`,
			mockBehavior: func() {
				mockUc.EXPECT().CreateProgram(gomock.Any(), gomock.Any()).Return(errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"database error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			req, _ := http.NewRequest(http.MethodPost, "/programs", strings.NewReader(tt.inputBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}

func Test_programHandlers_DeleteProgram(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUc := mocks.NewMockIProgramUsecase(ctrl)
	mockLogger := logger.NewLoggerTest()

	ph := &programHandlers{
		pu: mockUc,
		lg: mockLogger,
	}

	r := gin.Default()
	r.DELETE("/programs/:id", ph.DeleteProgram())

	tests := []struct {
		name           string
		programID      string
		mockBehavior   func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:      "Success - Delete Program",
			programID: "1",
			mockBehavior: func() {
				mockUc.EXPECT().DeleteProgram(gomock.Any(), 1).Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"Program deleted successfully"}`,
		},
		{
			name:           "Failed - Bad Request",
			programID:      "invalid",
			mockBehavior:   func() {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Invalid program ID"}`,
		},
		{
			name:      "Failed - Internal Server Error",
			programID: "1",
			mockBehavior: func() {
				mockUc.EXPECT().DeleteProgram(gomock.Any(), 1).Return(errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"database error"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			req, _ := http.NewRequest(http.MethodDelete, "/programs/"+tt.programID, nil)
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}
