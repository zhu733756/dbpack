// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cectc/dbpack/pkg/proto (interfaces: DBManager)

// Package testdata is a generated GoMock package.
package testdata

import (
	reflect "reflect"

	proto "github.com/cectc/dbpack/pkg/proto"
	gomock "github.com/golang/mock/gomock"
)

// MockDBManager is a mock of DBManager interface.
type MockDBManager struct {
	ctrl     *gomock.Controller
	recorder *MockDBManagerMockRecorder
}

// MockDBManagerMockRecorder is the mock recorder for MockDBManager.
type MockDBManagerMockRecorder struct {
	mock *MockDBManager
}

// NewMockDBManager creates a new mock instance.
func NewMockDBManager(ctrl *gomock.Controller) *MockDBManager {
	mock := &MockDBManager{ctrl: ctrl}
	mock.recorder = &MockDBManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBManager) EXPECT() *MockDBManagerMockRecorder {
	return m.recorder
}

// GetDB mocks base method.
func (m *MockDBManager) GetDB(arg0 string) proto.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDB", arg0)
	ret0, _ := ret[0].(proto.DB)
	return ret0
}

// GetDB indicates an expected call of GetDB.
func (mr *MockDBManagerMockRecorder) GetDB(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDB", reflect.TypeOf((*MockDBManager)(nil).GetDB), arg0)
}
