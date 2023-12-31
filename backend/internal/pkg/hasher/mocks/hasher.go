// Code generated by MockGen. DO NOT EDIT.
// Source: hasher.go

// Package mock_hasher is a generated GoMock package.
package mock_hasher

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHasher is a mock of Hasher interface.
type MockHasher struct {
	ctrl     *gomock.Controller
	recorder *MockHasherMockRecorder
}

// MockHasherMockRecorder is the mock recorder for MockHasher.
type MockHasherMockRecorder struct {
	mock *MockHasher
}

// NewMockHasher creates a new mock instance.
func NewMockHasher(ctrl *gomock.Controller) *MockHasher {
	mock := &MockHasher{ctrl: ctrl}
	mock.recorder = &MockHasherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHasher) EXPECT() *MockHasherMockRecorder {
	return m.recorder
}

// CheckUnhashedValue mocks base method.
func (m *MockHasher) CheckUnhashedValue(hashedString, unhashedString string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUnhashedValue", hashedString, unhashedString)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckUnhashedValue indicates an expected call of CheckUnhashedValue.
func (mr *MockHasherMockRecorder) CheckUnhashedValue(hashedString, unhashedString interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUnhashedValue", reflect.TypeOf((*MockHasher)(nil).CheckUnhashedValue), hashedString, unhashedString)
}

// GetHash mocks base method.
func (m *MockHasher) GetHash(stringToHash string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHash", stringToHash)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHash indicates an expected call of GetHash.
func (mr *MockHasherMockRecorder) GetHash(stringToHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHash", reflect.TypeOf((*MockHasher)(nil).GetHash), stringToHash)
}
