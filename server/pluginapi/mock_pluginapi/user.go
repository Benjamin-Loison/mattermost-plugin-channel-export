// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mattermost/mattermost-plugin-channel-export/server/pluginapi (interfaces: User)

// Package mock_pluginapi is a generated GoMock package.
package mock_pluginapi

import (
	gomock "github.com/golang/mock/gomock"
	model "github.com/mattermost/mattermost-server/v6/model"
	reflect "reflect"
)

// MockUser is a mock of User interface
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockUser) Get(arg0 string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockUserMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUser)(nil).Get), arg0)
}

// HasPermissionTo mocks base method
func (m *MockUser) HasPermissionTo(arg0 string, arg1 *model.Permission) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasPermissionTo", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasPermissionTo indicates an expected call of HasPermissionTo
func (mr *MockUserMockRecorder) HasPermissionTo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasPermissionTo", reflect.TypeOf((*MockUser)(nil).HasPermissionTo), arg0, arg1)
}

// HasPermissionToChannel mocks base method
func (m *MockUser) HasPermissionToChannel(arg0, arg1 string, arg2 *model.Permission) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasPermissionToChannel", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasPermissionToChannel indicates an expected call of HasPermissionToChannel
func (mr *MockUserMockRecorder) HasPermissionToChannel(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasPermissionToChannel", reflect.TypeOf((*MockUser)(nil).HasPermissionToChannel), arg0, arg1, arg2)
}
