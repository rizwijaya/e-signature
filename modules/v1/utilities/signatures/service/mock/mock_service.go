package mock

import (
	models "e-signature/modules/v1/utilities/signatures/models"
	models0 "e-signature/modules/v1/utilities/user/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// ChangeSignatures mocks base method.
func (m *MockService) ChangeSignatures(sign_type, idsignature string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeSignatures", sign_type, idsignature)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeSignatures indicates an expected call of ChangeSignatures.
func (mr *MockServiceMockRecorder) ChangeSignatures(sign_type, idsignature interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeSignatures", reflect.TypeOf((*MockService)(nil).ChangeSignatures), sign_type, idsignature)
}

// CreateImgSignature mocks base method.
func (m *MockService) CreateImgSignature(input models.AddSignature) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateImgSignature", input)
	ret0, _ := ret[0].(string)
	return ret0
}

// CreateImgSignature indicates an expected call of CreateImgSignature.
func (mr *MockServiceMockRecorder) CreateImgSignature(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateImgSignature", reflect.TypeOf((*MockService)(nil).CreateImgSignature), input)
}

// CreateImgSignatureData mocks base method.
func (m *MockService) CreateImgSignatureData(input models.AddSignature, name string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateImgSignatureData", input, name)
	ret0, _ := ret[0].(string)
	return ret0
}

// CreateImgSignatureData indicates an expected call of CreateImgSignatureData.
func (mr *MockServiceMockRecorder) CreateImgSignatureData(input, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateImgSignatureData", reflect.TypeOf((*MockService)(nil).CreateImgSignatureData), input, name)
}

// CreateLatinSignatures mocks base method.
func (m *MockService) CreateLatinSignatures(user models0.User, id string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLatinSignatures", user, id)
	ret0, _ := ret[0].(string)
	return ret0
}

// CreateLatinSignatures indicates an expected call of CreateLatinSignatures.
func (mr *MockServiceMockRecorder) CreateLatinSignatures(user, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLatinSignatures", reflect.TypeOf((*MockService)(nil).CreateLatinSignatures), user, id)
}

// CreateLatinSignaturesData mocks base method.
func (m *MockService) CreateLatinSignaturesData(user models0.User, latin, idn string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLatinSignaturesData", user, latin, idn)
	ret0, _ := ret[0].(string)
	return ret0
}

// CreateLatinSignaturesData indicates an expected call of CreateLatinSignaturesData.
func (mr *MockServiceMockRecorder) CreateLatinSignaturesData(user, latin, idn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLatinSignaturesData", reflect.TypeOf((*MockService)(nil).CreateLatinSignaturesData), user, latin, idn)
}

// DefaultSignatures mocks base method.
func (m *MockService) DefaultSignatures(user models0.User, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultSignatures", user, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DefaultSignatures indicates an expected call of DefaultSignatures.
func (mr *MockServiceMockRecorder) DefaultSignatures(user, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultSignatures", reflect.TypeOf((*MockService)(nil).DefaultSignatures), user, id)
}

// GetMySignature mocks base method.
func (m *MockService) GetMySignature(sign, id, name string) (models.MySignatures, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMySignature", sign, id, name)
	ret0, _ := ret[0].(models.MySignatures)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMySignature indicates an expected call of GetMySignature.
func (mr *MockServiceMockRecorder) GetMySignature(sign, id, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMySignature", reflect.TypeOf((*MockService)(nil).GetMySignature), sign, id, name)
}

// UpdateMySignatures mocks base method.
func (m *MockService) UpdateMySignatures(signature, signaturedata, sign string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMySignatures", signature, signaturedata, sign)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMySignatures indicates an expected call of UpdateMySignatures.
func (mr *MockServiceMockRecorder) UpdateMySignatures(signature, signaturedata, sign interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMySignatures", reflect.TypeOf((*MockService)(nil).UpdateMySignatures), signature, signaturedata, sign)
}
