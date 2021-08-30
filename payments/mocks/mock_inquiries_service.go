// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ngavinsir/golangtraining/payments (interfaces: InquiriesService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	golangtraining "github.com/ngavinsir/golangtraining"
)

// MockInquiriesService is a mock of InquiriesService interface.
type MockInquiriesService struct {
	ctrl     *gomock.Controller
	recorder *MockInquiriesServiceMockRecorder
}

// MockInquiriesServiceMockRecorder is the mock recorder for MockInquiriesService.
type MockInquiriesServiceMockRecorder struct {
	mock *MockInquiriesService
}

// NewMockInquiriesService creates a new mock instance.
func NewMockInquiriesService(ctrl *gomock.Controller) *MockInquiriesService {
	mock := &MockInquiriesService{ctrl: ctrl}
	mock.recorder = &MockInquiriesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInquiriesService) EXPECT() *MockInquiriesServiceMockRecorder {
	return m.recorder
}

// GetByTransactionID mocks base method.
func (m *MockInquiriesService) GetByTransactionID(arg0 context.Context, arg1 string) (golangtraining.Inquiry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByTransactionID", arg0, arg1)
	ret0, _ := ret[0].(golangtraining.Inquiry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByTransactionID indicates an expected call of GetByTransactionID.
func (mr *MockInquiriesServiceMockRecorder) GetByTransactionID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByTransactionID", reflect.TypeOf((*MockInquiriesService)(nil).GetByTransactionID), arg0, arg1)
}
