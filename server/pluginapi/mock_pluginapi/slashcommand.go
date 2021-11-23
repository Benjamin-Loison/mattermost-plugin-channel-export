// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mattermost/mattermost-plugin-channel-export/server/pluginapi (interfaces: SlashCommand)

// Package mock_pluginapi is a generated GoMock package.
package mock_pluginapi

import (
	gomock "github.com/golang/mock/gomock"
	model "github.com/mattermost/mattermost-server/v6/model"
	reflect "reflect"
)

// MockSlashCommand is a mock of SlashCommand interface
type MockSlashCommand struct {
	ctrl     *gomock.Controller
	recorder *MockSlashCommandMockRecorder
}

// MockSlashCommandMockRecorder is the mock recorder for MockSlashCommand
type MockSlashCommandMockRecorder struct {
	mock *MockSlashCommand
}

// NewMockSlashCommand creates a new mock instance
func NewMockSlashCommand(ctrl *gomock.Controller) *MockSlashCommand {
	mock := &MockSlashCommand{ctrl: ctrl}
	mock.recorder = &MockSlashCommandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSlashCommand) EXPECT() *MockSlashCommandMockRecorder {
	return m.recorder
}

// Register mocks base method
func (m *MockSlashCommand) Register(arg0 *model.Command) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockSlashCommandMockRecorder) Register(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockSlashCommand)(nil).Register), arg0)
}
