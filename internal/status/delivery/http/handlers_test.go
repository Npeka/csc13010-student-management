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
	"github.com/csc13010-student-management/internal/status"
	"github.com/csc13010-student-management/internal/status/mocks"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewStatusHandlers(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUc := mocks.NewMockIStatusUsecase(ctrl)
	mockLogger := logger.NewLoggerTest()

	type args struct {
		su status.IStatusUsecase
		lg *logger.LoggerZap
	}
	tests := []struct {
		name string
		args args
		want status.IStatusHandlers
	}{
		{
			name: "Success - Create Status Handlers",
			args: args{
				su: mockUc,
				lg: mockLogger,
			},
			want: NewStatusHandlers(mockUc, mockLogger),
		},
		{
			name: "Failed - Create Status Handlers",
			args: args{
				su: nil,
				lg: nil,
			},
			want: NewStatusHandlers(nil, nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStatusHandlers(tt.args.su, tt.args.lg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStatusHandlers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_statusHandlers_GetStatuses(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUc := mocks.NewMockIStatusUsecase(ctrl)
	mockLogger := logger.NewLoggerTest()
	mockStatuses := []*models.Status{
		{ID: 1, Name: "Status 1"},
		{ID: 2, Name: "Status 2"},
	}

	sh := &statusHandlers{
		su: mockUc,
		lg: mockLogger,
	}

	r := gin.Default()
	r.GET("/statuses", sh.GetStatuses())

	tests := []struct {
		name           string
		mockBehavior   func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Success - Get Statuses",
			mockBehavior: func() {
				mockUc.EXPECT().GetStatuses(gomock.Any()).Return(mockStatuses, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: func() string {
				bytes, _ := json.Marshal(mockStatuses)
				return string(bytes)
			}(),
		},
		{
			name: "Failed - Internal Server Error",
			mockBehavior: func() {
				mockUc.EXPECT().GetStatuses(gomock.Any()).Return(nil, errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"database error"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			req, _ := http.NewRequest(http.MethodGet, "/statuses", nil)
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}

func Test_statusHandlers_CreateStatus(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUc := mocks.NewMockIStatusUsecase(ctrl)
	mockLogger := logger.NewLoggerTest()

	sh := &statusHandlers{
		su: mockUc,
		lg: mockLogger,
	}

	r := gin.Default()
	r.POST("/statuses", sh.CreateStatus())

	tests := []struct {
		name           string
		inputBody      string
		mockBehavior   func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:      "Success - Create Status",
			inputBody: `{"name": "Status 1"}`,
			mockBehavior: func() {
				mockUc.EXPECT().CreateStatus(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"id":0,"name":"Status 1"}`,
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
			inputBody: `{"name": "Status 1"}`,
			mockBehavior: func() {
				mockUc.EXPECT().CreateStatus(gomock.Any(), gomock.Any()).Return(errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"database error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			req, _ := http.NewRequest(http.MethodPost, "/statuses", strings.NewReader(tt.inputBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}

func Test_statusHandlers_DeleteStatus(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUc := mocks.NewMockIStatusUsecase(ctrl)
	mockLogger := logger.NewLoggerTest()

	sh := &statusHandlers{
		su: mockUc,
		lg: mockLogger,
	}

	r := gin.Default()
	r.DELETE("/statuses/:id", sh.DeleteStatus())

	tests := []struct {
		name           string
		statusID       string
		mockBehavior   func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:     "Success - Delete Status",
			statusID: "1",
			mockBehavior: func() {
				mockUc.EXPECT().DeleteStatus(gomock.Any(), 1).Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"Status deleted successfully"}`,
		},
		{
			name:           "Failed - Bad Request",
			statusID:       "invalid",
			mockBehavior:   func() {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Invalid status ID"}`,
		},
		{
			name:     "Failed - Internal Server Error",
			statusID: "1",
			mockBehavior: func() {
				mockUc.EXPECT().DeleteStatus(gomock.Any(), 1).Return(errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"database error"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			req, _ := http.NewRequest(http.MethodDelete, "/statuses/"+tt.statusID, nil)
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}

func Test_statusHandlers_UpdateStatus(t *testing.T) {
	type fields struct {
		su status.IStatusUsecase
		lg *logger.LoggerZap
	}
	tests := []struct {
		name   string
		fields fields
		want   gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sh := &statusHandlers{
				su: tt.fields.su,
				lg: tt.fields.lg,
			}
			if got := sh.UpdateStatus(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("statusHandlers.UpdateStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
