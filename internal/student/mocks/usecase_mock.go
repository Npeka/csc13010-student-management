// Code generated by MockGen. DO NOT EDIT.
// Source: internal/student/usecase.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	models "github.com/csc13010-student-management/internal/models"
	dtos "github.com/csc13010-student-management/internal/student/dtos"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockIStudentUsecase is a mock of IStudentUsecase interface.
type MockIStudentUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockIStudentUsecaseMockRecorder
}

// MockIStudentUsecaseMockRecorder is the mock recorder for MockIStudentUsecase.
type MockIStudentUsecaseMockRecorder struct {
	mock *MockIStudentUsecase
}

// NewMockIStudentUsecase creates a new mock instance.
func NewMockIStudentUsecase(ctrl *gomock.Controller) *MockIStudentUsecase {
	mock := &MockIStudentUsecase{ctrl: ctrl}
	mock.recorder = &MockIStudentUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStudentUsecase) EXPECT() *MockIStudentUsecaseMockRecorder {
	return m.recorder
}

// BatchUpdateUserIDs mocks base method.
func (m *MockIStudentUsecase) BatchUpdateUserIDs(ctx context.Context, userIDs map[string]uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchUpdateUserIDs", ctx, userIDs)
	ret0, _ := ret[0].(error)
	return ret0
}

// BatchUpdateUserIDs indicates an expected call of BatchUpdateUserIDs.
func (mr *MockIStudentUsecaseMockRecorder) BatchUpdateUserIDs(ctx, userIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchUpdateUserIDs", reflect.TypeOf((*MockIStudentUsecase)(nil).BatchUpdateUserIDs), ctx, userIDs)
}

// CreateStudent mocks base method.
func (m *MockIStudentUsecase) CreateStudent(ctx context.Context, student *models.Student) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStudent", ctx, student)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateStudent indicates an expected call of CreateStudent.
func (mr *MockIStudentUsecaseMockRecorder) CreateStudent(ctx, student interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStudent", reflect.TypeOf((*MockIStudentUsecase)(nil).CreateStudent), ctx, student)
}

// DeleteStudent mocks base method.
func (m *MockIStudentUsecase) DeleteStudent(ctx context.Context, student_id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStudent", ctx, student_id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStudent indicates an expected call of DeleteStudent.
func (mr *MockIStudentUsecaseMockRecorder) DeleteStudent(ctx, student_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStudent", reflect.TypeOf((*MockIStudentUsecase)(nil).DeleteStudent), ctx, student_id)
}

// GetFullInfoStudentByStudentID mocks base method.
func (m *MockIStudentUsecase) GetFullInfoStudentByStudentID(ctx context.Context, student_id string) (*dtos.StudentResponseDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFullInfoStudentByStudentID", ctx, student_id)
	ret0, _ := ret[0].(*dtos.StudentResponseDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFullInfoStudentByStudentID indicates an expected call of GetFullInfoStudentByStudentID.
func (mr *MockIStudentUsecaseMockRecorder) GetFullInfoStudentByStudentID(ctx, student_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFullInfoStudentByStudentID", reflect.TypeOf((*MockIStudentUsecase)(nil).GetFullInfoStudentByStudentID), ctx, student_id)
}

// GetOptions mocks base method.
func (m *MockIStudentUsecase) GetOptions(ctx context.Context) (*dtos.OptionDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOptions", ctx)
	ret0, _ := ret[0].(*dtos.OptionDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOptions indicates an expected call of GetOptions.
func (mr *MockIStudentUsecaseMockRecorder) GetOptions(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOptions", reflect.TypeOf((*MockIStudentUsecase)(nil).GetOptions), ctx)
}

// GetStudentByStudentID mocks base method.
func (m *MockIStudentUsecase) GetStudentByStudentID(ctx context.Context, student_id string) (*models.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStudentByStudentID", ctx, student_id)
	ret0, _ := ret[0].(*models.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudentByStudentID indicates an expected call of GetStudentByStudentID.
func (mr *MockIStudentUsecaseMockRecorder) GetStudentByStudentID(ctx, student_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudentByStudentID", reflect.TypeOf((*MockIStudentUsecase)(nil).GetStudentByStudentID), ctx, student_id)
}

// GetStudents mocks base method.
func (m *MockIStudentUsecase) GetStudents(ctx context.Context) ([]*dtos.StudentResponseDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStudents", ctx)
	ret0, _ := ret[0].([]*dtos.StudentResponseDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudents indicates an expected call of GetStudents.
func (mr *MockIStudentUsecaseMockRecorder) GetStudents(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudents", reflect.TypeOf((*MockIStudentUsecase)(nil).GetStudents), ctx)
}

// UpdateStudent mocks base method.
func (m *MockIStudentUsecase) UpdateStudent(ctx context.Context, student *models.Student) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStudent", ctx, student)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStudent indicates an expected call of UpdateStudent.
func (mr *MockIStudentUsecaseMockRecorder) UpdateStudent(ctx, student interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStudent", reflect.TypeOf((*MockIStudentUsecase)(nil).UpdateStudent), ctx, student)
}

// UpdateUserIDByUsername mocks base method.
func (m *MockIStudentUsecase) UpdateUserIDByUsername(ctx context.Context, student_id string, user_id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserIDByUsername", ctx, student_id, user_id)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserIDByUsername indicates an expected call of UpdateUserIDByUsername.
func (mr *MockIStudentUsecaseMockRecorder) UpdateUserIDByUsername(ctx, student_id, user_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserIDByUsername", reflect.TypeOf((*MockIStudentUsecase)(nil).UpdateUserIDByUsername), ctx, student_id, user_id)
}
