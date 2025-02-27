// Code generated by MockGen. DO NOT EDIT.
// Source: internal/program/usecase.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	models "github.com/csc13010-student-management/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockIProgramUsecase is a mock of IProgramUsecase interface.
type MockIProgramUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockIProgramUsecaseMockRecorder
}

// MockIProgramUsecaseMockRecorder is the mock recorder for MockIProgramUsecase.
type MockIProgramUsecaseMockRecorder struct {
	mock *MockIProgramUsecase
}

// NewMockIProgramUsecase creates a new mock instance.
func NewMockIProgramUsecase(ctrl *gomock.Controller) *MockIProgramUsecase {
	mock := &MockIProgramUsecase{ctrl: ctrl}
	mock.recorder = &MockIProgramUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProgramUsecase) EXPECT() *MockIProgramUsecaseMockRecorder {
	return m.recorder
}

// CreateProgram mocks base method.
func (m *MockIProgramUsecase) CreateProgram(ctx context.Context, program *models.Program) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProgram", ctx, program)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProgram indicates an expected call of CreateProgram.
func (mr *MockIProgramUsecaseMockRecorder) CreateProgram(ctx, program interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProgram", reflect.TypeOf((*MockIProgramUsecase)(nil).CreateProgram), ctx, program)
}

// DeleteProgram mocks base method.
func (m *MockIProgramUsecase) DeleteProgram(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProgram", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProgram indicates an expected call of DeleteProgram.
func (mr *MockIProgramUsecaseMockRecorder) DeleteProgram(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProgram", reflect.TypeOf((*MockIProgramUsecase)(nil).DeleteProgram), ctx, id)
}

// GetPrograms mocks base method.
func (m *MockIProgramUsecase) GetPrograms(ctx context.Context) ([]*models.Program, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrograms", ctx)
	ret0, _ := ret[0].([]*models.Program)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPrograms indicates an expected call of GetPrograms.
func (mr *MockIProgramUsecaseMockRecorder) GetPrograms(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrograms", reflect.TypeOf((*MockIProgramUsecase)(nil).GetPrograms), ctx)
}
