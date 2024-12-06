// Code generated by MockGen. DO NOT EDIT.
// Source: internal/controller/controller.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	model "backend/internal/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockControllerInterface is a mock of ControllerInterface interface.
type MockControllerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockControllerInterfaceMockRecorder
}

// MockControllerInterfaceMockRecorder is the mock recorder for MockControllerInterface.
type MockControllerInterfaceMockRecorder struct {
	mock *MockControllerInterface
}

// NewMockControllerInterface creates a new mock instance.
func NewMockControllerInterface(ctrl *gomock.Controller) *MockControllerInterface {
	mock := &MockControllerInterface{ctrl: ctrl}
	mock.recorder = &MockControllerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockControllerInterface) EXPECT() *MockControllerInterfaceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockControllerInterface) CreateUser(user model.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockControllerInterfaceMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockControllerInterface)(nil).CreateUser), user)
}

// DeleteUser mocks base method.
func (m *MockControllerInterface) DeleteUser(ID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockControllerInterfaceMockRecorder) DeleteUser(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockControllerInterface)(nil).DeleteUser), ID)
}

// GetUser mocks base method.
func (m *MockControllerInterface) GetUser(ID int) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ID)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockControllerInterfaceMockRecorder) GetUser(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockControllerInterface)(nil).GetUser), ID)
}

// GetUsers mocks base method.
func (m *MockControllerInterface) GetUsers() ([]model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers")
	ret0, _ := ret[0].([]model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockControllerInterfaceMockRecorder) GetUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockControllerInterface)(nil).GetUsers))
}

// UpdateUser mocks base method.
func (m *MockControllerInterface) UpdateUser(ID int, user model.User) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ID, user)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockControllerInterfaceMockRecorder) UpdateUser(ID, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockControllerInterface)(nil).UpdateUser), ID, user)
}
