// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/leminhviett/Go_micro_practice-HTTP_api_server/service (interfaces: AccountService)

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dto "github.com/leminhviett/Go_micro_practice-HTTP_api_server/dto"
	errs "github.com/leminhviett/Go_micro_practice-utils_lib/errs"
)

// MockAccountService is a mock of AccountService interface.
type MockAccountService struct {
	ctrl     *gomock.Controller
	recorder *MockAccountServiceMockRecorder
}

// MockAccountServiceMockRecorder is the mock recorder for MockAccountService.
type MockAccountServiceMockRecorder struct {
	mock *MockAccountService
}

// NewMockAccountService creates a new mock instance.
func NewMockAccountService(ctrl *gomock.Controller) *MockAccountService {
	mock := &MockAccountService{ctrl: ctrl}
	mock.recorder = &MockAccountServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountService) EXPECT() *MockAccountServiceMockRecorder {
	return m.recorder
}

// FindAll mocks base method.
func (m *MockAccountService) FindAll() ([]*dto.AccountResponseDTO, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]*dto.AccountResponseDTO)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockAccountServiceMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockAccountService)(nil).FindAll))
}

// FindById mocks base method.
func (m *MockAccountService) FindById(arg0 string) (*dto.AccountResponseDTO, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", arg0)
	ret0, _ := ret[0].(*dto.AccountResponseDTO)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockAccountServiceMockRecorder) FindById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockAccountService)(nil).FindById), arg0)
}

// InsertNewAccount mocks base method.
func (m *MockAccountService) InsertNewAccount(arg0 *dto.AccountRequestDTO) (*dto.AccountResponseDTO, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertNewAccount", arg0)
	ret0, _ := ret[0].(*dto.AccountResponseDTO)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// InsertNewAccount indicates an expected call of InsertNewAccount.
func (mr *MockAccountServiceMockRecorder) InsertNewAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertNewAccount", reflect.TypeOf((*MockAccountService)(nil).InsertNewAccount), arg0)
}
