// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	models "backend/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockClientRepository is a mock of ClientRepository interface.
type MockClientRepository struct {
	ctrl     *gomock.Controller
	recorder *MockClientRepositoryMockRecorder
}

// MockClientRepositoryMockRecorder is the mock recorder for MockClientRepository.
type MockClientRepositoryMockRecorder struct {
	mock *MockClientRepository
}

// NewMockClientRepository creates a new mock instance.
func NewMockClientRepository(ctrl *gomock.Controller) *MockClientRepository {
	mock := &MockClientRepository{ctrl: ctrl}
	mock.recorder = &MockClientRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientRepository) EXPECT() *MockClientRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockClientRepository) Create(client *models.Client) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", client)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockClientRepositoryMockRecorder) Create(client interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockClientRepository)(nil).Create), client)
}

// Delete mocks base method.
func (m *MockClientRepository) Delete(id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockClientRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClientRepository)(nil).Delete), id)
}

// GetAllClient mocks base method.
func (m *MockClientRepository) GetAllClient() ([]models.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllClient")
	ret0, _ := ret[0].([]models.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllClient indicates an expected call of GetAllClient.
func (mr *MockClientRepositoryMockRecorder) GetAllClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllClient", reflect.TypeOf((*MockClientRepository)(nil).GetAllClient))
}

// GetClientById mocks base method.
func (m *MockClientRepository) GetClientById(id uint64) (*models.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClientById", id)
	ret0, _ := ret[0].(*models.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClientById indicates an expected call of GetClientById.
func (mr *MockClientRepositoryMockRecorder) GetClientById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClientById", reflect.TypeOf((*MockClientRepository)(nil).GetClientById), id)
}

// GetClientByLogin mocks base method.
func (m *MockClientRepository) GetClientByLogin(login string) (*models.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClientByLogin", login)
	ret0, _ := ret[0].(*models.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClientByLogin indicates an expected call of GetClientByLogin.
func (mr *MockClientRepositoryMockRecorder) GetClientByLogin(login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClientByLogin", reflect.TypeOf((*MockClientRepository)(nil).GetClientByLogin), login)
}

// SetRole mocks base method.
func (m *MockClientRepository) SetRole() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetRole")
	ret0, _ := ret[0].(error)
	return ret0
}

// SetRole indicates an expected call of SetRole.
func (mr *MockClientRepositoryMockRecorder) SetRole() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRole", reflect.TypeOf((*MockClientRepository)(nil).SetRole))
}
