// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/usecase/user.go
//
// Generated by this command:
//
//	mockgen -source=./internal/usecase/user.go -destination=./internal/usecase/mock/user.go -package mock_usecase
//

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"
	models "sharecycle/internal/models"

	gomock "go.uber.org/mock/gomock"
)

// MockUserInputPort is a mock of UserInputPort interface.
type MockUserInputPort struct {
	ctrl     *gomock.Controller
	recorder *MockUserInputPortMockRecorder
}

// MockUserInputPortMockRecorder is the mock recorder for MockUserInputPort.
type MockUserInputPortMockRecorder struct {
	mock *MockUserInputPort
}

// NewMockUserInputPort creates a new mock instance.
func NewMockUserInputPort(ctrl *gomock.Controller) *MockUserInputPort {
	mock := &MockUserInputPort{ctrl: ctrl}
	mock.recorder = &MockUserInputPortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserInputPort) EXPECT() *MockUserInputPortMockRecorder {
	return m.recorder
}

// GetUserByID mocks base method.
func (m *MockUserInputPort) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, id)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockUserInputPortMockRecorder) GetUserByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUserInputPort)(nil).GetUserByID), ctx, id)
}
