package mocks

import (
	reflect "reflect"

	"github.com/Ricardolv/mvc-api/src/config/rest_err"
	gomock "go.uber.org/mock/gomock"
)

// MockUserDomainInterface is a mock of UserDomainInterface interface.
type MockUserDomainInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserDomainInterfaceMockRecorder
}

// MockUserDomainInterfaceMockRecorder is the mock recorder for MockUserDomainInterface.
type MockUserDomainInterfaceMockRecorder struct {
	mock *MockUserDomainInterface
}

// NewMockUserDomainInterface creates a new mock instance.
func NewMockUserDomainInterface(ctrl *gomock.Controller) *MockUserDomainInterface {
	mock := &MockUserDomainInterface{ctrl: ctrl}
	mock.recorder = &MockUserDomainInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserDomainInterface) EXPECT() *MockUserDomainInterfaceMockRecorder {
	return m.recorder
}

// EncryptPassword mocks base method.
func (m *MockUserDomainInterface) EncryptPassword() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EncryptPassword")
}

// EncryptPassword indicates an expected call of EncryptPassword.
func (mr *MockUserDomainInterfaceMockRecorder) EncryptPassword() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EncryptPassword", reflect.TypeOf((*MockUserDomainInterface)(nil).EncryptPassword))
}

// GenerateToken mocks base method.
func (m *MockUserDomainInterface) GenerateToken() (string, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*rest_err.RestErr)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockUserDomainInterfaceMockRecorder) GenerateToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockUserDomainInterface)(nil).GenerateToken))
}

// GetAge mocks base method.
func (m *MockUserDomainInterface) GetAge() int8 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAge")
	ret0, _ := ret[0].(int8)
	return ret0
}

// GetAge indicates an expected call of GetAge.
func (mr *MockUserDomainInterfaceMockRecorder) GetAge() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAge", reflect.TypeOf((*MockUserDomainInterface)(nil).GetAge))
}

// GetEmail mocks base method.
func (m *MockUserDomainInterface) GetEmail() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmail")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetEmail indicates an expected call of GetEmail.
func (mr *MockUserDomainInterfaceMockRecorder) GetEmail() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmail", reflect.TypeOf((*MockUserDomainInterface)(nil).GetEmail))
}

// GetID mocks base method.
func (m *MockUserDomainInterface) GetID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetID indicates an expected call of GetID.
func (mr *MockUserDomainInterfaceMockRecorder) GetID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockUserDomainInterface)(nil).GetID))
}

// GetName mocks base method.
func (m *MockUserDomainInterface) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockUserDomainInterfaceMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockUserDomainInterface)(nil).GetName))
}

// GetPassword mocks base method.
func (m *MockUserDomainInterface) GetPassword() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPassword")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPassword indicates an expected call of GetPassword.
func (mr *MockUserDomainInterfaceMockRecorder) GetPassword() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPassword", reflect.TypeOf((*MockUserDomainInterface)(nil).GetPassword))
}

// SetID mocks base method.
func (m *MockUserDomainInterface) SetID(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetID", arg0)
}

// SetID indicates an expected call of SetID.
func (mr *MockUserDomainInterfaceMockRecorder) SetID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetID", reflect.TypeOf((*MockUserDomainInterface)(nil).SetID), arg0)
}
