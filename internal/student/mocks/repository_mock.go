// Code generated by MockGen. DO NOT EDIT.
// Source: internal/student/repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	models "github.com/csc13010-student-management/internal/models"
	dtos "github.com/csc13010-student-management/internal/student/dtos"
	gomock "github.com/golang/mock/gomock"
)

// MockIStudentRepository is a mock of IStudentRepository interface.
type MockIStudentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIStudentRepositoryMockRecorder
}

// MockIStudentRepositoryMockRecorder is the mock recorder for MockIStudentRepository.
type MockIStudentRepositoryMockRecorder struct {
	mock *MockIStudentRepository
}

// NewMockIStudentRepository creates a new mock instance.
func NewMockIStudentRepository(ctrl *gomock.Controller) *MockIStudentRepository {
	mock := &MockIStudentRepository{ctrl: ctrl}
	mock.recorder = &MockIStudentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStudentRepository) EXPECT() *MockIStudentRepositoryMockRecorder {
	return m.recorder
}

// BatchInsertStudents mocks base method.
func (m *MockIStudentRepository) BatchInsertStudents(ctx context.Context, students []models.Student) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchInsertStudents", ctx, students)
	ret0, _ := ret[0].(error)
	return ret0
}

// BatchInsertStudents indicates an expected call of BatchInsertStudents.
func (mr *MockIStudentRepositoryMockRecorder) BatchInsertStudents(ctx, students interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchInsertStudents", reflect.TypeOf((*MockIStudentRepository)(nil).BatchInsertStudents), ctx, students)
}

// CreateStudent mocks base method.
func (m *MockIStudentRepository) CreateStudent(ctx context.Context, student *models.Student) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStudent", ctx, student)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateStudent indicates an expected call of CreateStudent.
func (mr *MockIStudentRepositoryMockRecorder) CreateStudent(ctx, student interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStudent", reflect.TypeOf((*MockIStudentRepository)(nil).CreateStudent), ctx, student)
}

// DeleteStudent mocks base method.
func (m *MockIStudentRepository) DeleteStudent(ctx context.Context, student_id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStudent", ctx, student_id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStudent indicates an expected call of DeleteStudent.
func (mr *MockIStudentRepositoryMockRecorder) DeleteStudent(ctx, student_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStudent", reflect.TypeOf((*MockIStudentRepository)(nil).DeleteStudent), ctx, student_id)
}

// GetAllStudents mocks base method.
func (m *MockIStudentRepository) GetAllStudents(ctx context.Context, students *[]models.Student) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllStudents", ctx, students)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetAllStudents indicates an expected call of GetAllStudents.
func (mr *MockIStudentRepositoryMockRecorder) GetAllStudents(ctx, students interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllStudents", reflect.TypeOf((*MockIStudentRepository)(nil).GetAllStudents), ctx, students)
}

// GetOptions mocks base method.
func (m *MockIStudentRepository) GetOptions(ctx context.Context) (*dtos.OptionDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOptions", ctx)
	ret0, _ := ret[0].(*dtos.OptionDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOptions indicates an expected call of GetOptions.
func (mr *MockIStudentRepositoryMockRecorder) GetOptions(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOptions", reflect.TypeOf((*MockIStudentRepository)(nil).GetOptions), ctx)
}

// GetStudentByStudentID mocks base method.
func (m *MockIStudentRepository) GetStudentByStudentID(ctx context.Context, student_id string) (*models.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStudentByStudentID", ctx, student_id)
	ret0, _ := ret[0].(*models.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudentByStudentID indicates an expected call of GetStudentByStudentID.
func (mr *MockIStudentRepositoryMockRecorder) GetStudentByStudentID(ctx, student_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudentByStudentID", reflect.TypeOf((*MockIStudentRepository)(nil).GetStudentByStudentID), ctx, student_id)
}

// GetStudents mocks base method.
func (m *MockIStudentRepository) GetStudents(ctx context.Context) ([]*models.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStudents", ctx)
	ret0, _ := ret[0].([]*models.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudents indicates an expected call of GetStudents.
func (mr *MockIStudentRepositoryMockRecorder) GetStudents(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudents", reflect.TypeOf((*MockIStudentRepository)(nil).GetStudents), ctx)
}

// UpdateStudent mocks base method.
func (m *MockIStudentRepository) UpdateStudent(ctx context.Context, student *models.Student) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStudent", ctx, student)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStudent indicates an expected call of UpdateStudent.
func (mr *MockIStudentRepositoryMockRecorder) UpdateStudent(ctx, student interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStudent", reflect.TypeOf((*MockIStudentRepository)(nil).UpdateStudent), ctx, student)
}
