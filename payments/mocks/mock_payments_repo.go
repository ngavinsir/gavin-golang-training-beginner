// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ngavinsir/golangtraining/payments (interfaces: PaymentsRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	golangtraining "github.com/ngavinsir/golangtraining"
)

// MockPaymentsRepository is a mock of PaymentsRepository interface.
type MockPaymentsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentsRepositoryMockRecorder
}

// MockPaymentsRepositoryMockRecorder is the mock recorder for MockPaymentsRepository.
type MockPaymentsRepositoryMockRecorder struct {
	mock *MockPaymentsRepository
}

// NewMockPaymentsRepository creates a new mock instance.
func NewMockPaymentsRepository(ctrl *gomock.Controller) *MockPaymentsRepository {
	mock := &MockPaymentsRepository{ctrl: ctrl}
	mock.recorder = &MockPaymentsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaymentsRepository) EXPECT() *MockPaymentsRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPaymentsRepository) Create(arg0 context.Context, arg1 *golangtraining.Payment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockPaymentsRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPaymentsRepository)(nil).Create), arg0, arg1)
}
