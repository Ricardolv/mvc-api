package mocks

import (
	reflect "reflect"

	"github.com/Ricardolv/mvc-api/src/config/rest_err"
	"github.com/Ricardolv/mvc-api/src/model"
	gomock "go.uber.org/mock/gomock"
)

// MockUserDomainService is a mock of UserDomainService interface.
type MockUserDomainService struct {
	ctrl     *gomock.Controller
	recorder *MockUserDomainServiceMockRecorder
}

// MockUserDomainServiceMockRecorder is the mock recorder for MockUserDomainService.
type MockUserDomainServiceMockRecorder struct {
	mock *MockUserDomainService
}

// NewMockUserDomainService creates a new mock instance.
func NewMockUserDomainService(ctrl *gomock.Controller) *MockUserDomainService {
	mock := &MockUserDomainService{ctrl: ctrl}
	mock.recorder = &MockUserDomainServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserDomainService) EXPECT() *MockUserDomainServiceMockRecorder {
	return m.recorder
}

// CreateUserServices mocks base method.
func (m *MockUserDomainService) CreateService(arg0 model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateServices", arg0)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(*rest_err.RestErr)
	return ret0, ret1
}

// CreateUserServices indicates an expected call of CreateUserServices.
func (mr *MockUserDomainServiceMockRecorder) CreateService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateServices", reflect.TypeOf((*MockUserDomainService)(nil).CreateService), arg0)
}

// DeleteUser mocks base method.
func (m *MockUserDomainService) Delete(arg0 string) *rest_err.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(*rest_err.RestErr)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserDomainServiceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserDomainService)(nil).Delete), arg0)
}

// FindUserByEmailServices mocks base method.
func (m *MockUserDomainService) FindByEmailService(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmailServices", email)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(*rest_err.RestErr)
	return ret0, ret1
}

// FindUserByEmailServices indicates an expected call of FindUserByEmailServices.
func (mr *MockUserDomainServiceMockRecorder) FindByEmailService(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmailServices", reflect.TypeOf((*MockUserDomainService)(nil).FindByEmailService), email)
}

// FindUserByIDServices mocks base method.
func (m *MockUserDomainService) FindByIDService(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByIDServices", id)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(*rest_err.RestErr)
	return ret0, ret1
}

// FindUserByIDServices indicates an expected call of FindUserByIDServices.
func (mr *MockUserDomainServiceMockRecorder) FindByIDService(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIDServices", reflect.TypeOf((*MockUserDomainService)(nil).FindByIDService), id)
}

// LoginUserServices mocks base method.
func (m *MockUserDomainService) LoginUserService(userDomain model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginUserService", userDomain)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(*rest_err.RestErr)
	return ret0, ret1, ret2
}

// LoginUserServices indicates an expected call of LoginUserServices.
func (mr *MockUserDomainServiceMockRecorder) LoginUserService(userDomain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginUserServices", reflect.TypeOf((*MockUserDomainService)(nil).LoginUserService), userDomain)
}

// UpdateUser mocks base method.
func (m *MockUserDomainService) Update(arg0 string, arg1 model.UserDomainInterface) *rest_err.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*rest_err.RestErr)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserDomainServiceMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserDomainService)(nil).Update), arg0, arg1)
}
