// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mattermost/mattermost-plugin-channel-export/server/pluginapi (interfaces: System)

// Package mock_pluginapi is a generated GoMock package.
package mock_pluginapi

import (
	gomock "github.com/golang/mock/gomock"
	model "github.com/mattermost/mattermost-server/v6/model"
	reflect "reflect"
)

// MockSystem is a mock of System interface
type MockSystem struct {
	ctrl     *gomock.Controller
	recorder *MockSystemMockRecorder
}

// MockSystemMockRecorder is the mock recorder for MockSystem
type MockSystemMockRecorder struct {
	mock *MockSystem
}

// NewMockSystem creates a new mock instance
func NewMockSystem(ctrl *gomock.Controller) *MockSystem {
	mock := &MockSystem{ctrl: ctrl}
	mock.recorder = &MockSystemMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSystem) EXPECT() *MockSystemMockRecorder {
	return m.recorder
}

// GetLicense mocks base method
func (m *MockSystem) GetLicense() *model.License {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLicense")
	ret0, _ := ret[0].(*model.License)
	return ret0
}

// GetLicense indicates an expected call of GetLicense
func (mr *MockSystemMockRecorder) GetLicense() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLicense", reflect.TypeOf((*MockSystem)(nil).GetLicense))
}
