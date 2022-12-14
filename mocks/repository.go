// Code generated by MockGen. DO NOT EDIT.
// Source: ../repository/interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/guntoroyk/go-user-api/entity"
)

// MockUserRepoItf is a mock of UserRepoItf interface.
type MockUserRepoItf struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepoItfMockRecorder
}

// MockUserRepoItfMockRecorder is the mock recorder for MockUserRepoItf.
type MockUserRepoItfMockRecorder struct {
	mock *MockUserRepoItf
}

// NewMockUserRepoItf creates a new mock instance.
func NewMockUserRepoItf(ctrl *gomock.Controller) *MockUserRepoItf {
	mock := &MockUserRepoItf{ctrl: ctrl}
	mock.recorder = &MockUserRepoItfMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepoItf) EXPECT() *MockUserRepoItfMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserRepoItf) CreateUser(user *entity.User) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepoItfMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepoItf)(nil).CreateUser), user)
}

// DeleteUser mocks base method.
func (m *MockUserRepoItf) DeleteUser(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserRepoItfMockRecorder) DeleteUser(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserRepoItf)(nil).DeleteUser), id)
}

// GetUser mocks base method.
func (m *MockUserRepoItf) GetUser(filter *entity.UserFilter) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", filter)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserRepoItfMockRecorder) GetUser(filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserRepoItf)(nil).GetUser), filter)
}

// GetUsers mocks base method.
func (m *MockUserRepoItf) GetUsers() ([]*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers")
	ret0, _ := ret[0].([]*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockUserRepoItfMockRecorder) GetUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUserRepoItf)(nil).GetUsers))
}

// UpdateUser mocks base method.
func (m *MockUserRepoItf) UpdateUser(user *entity.User) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", user)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserRepoItfMockRecorder) UpdateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserRepoItf)(nil).UpdateUser), user)
}
