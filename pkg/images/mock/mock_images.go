package mock_images

import (
	models "e-signature/modules/v1/utilities/signatures/models"
	models0 "e-signature/modules/v1/utilities/user/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockImages is a mock of Images interface.
type MockImages struct {
	ctrl     *gomock.Controller
	recorder *MockImagesMockRecorder
}

// MockImagesMockRecorder is the mock recorder for MockImages.
type MockImagesMockRecorder struct {
	mock *MockImages
}

// NewMockImages creates a new mock instance.
func NewMockImages(ctrl *gomock.Controller) *MockImages {
	mock := &MockImages{ctrl: ctrl}
	mock.recorder = &MockImagesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImages) EXPECT() *MockImagesMockRecorder {
	return m.recorder
}

// CreateImageSignature mocks base method.
func (m *MockImages) CreateImageSignature(input models.AddSignature) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateImageSignature", input)
	ret0, _ := ret[0].(string)
	return ret0
}

// CreateImageSignature indicates an expected call of CreateImageSignature.
func (mr *MockImagesMockRecorder) CreateImageSignature(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateImageSignature", reflect.TypeOf((*MockImages)(nil).CreateImageSignature), input)
}

// CreateImgSignatureData mocks base method.
func (m *MockImages) CreateImgSignatureData(input models.AddSignature, name, font string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateImgSignatureData", input, name, font)
	ret0, _ := ret[0].(string)
	return ret0
}

// CreateImgSignatureData indicates an expected call of CreateImgSignatureData.
func (mr *MockImagesMockRecorder) CreateImgSignatureData(input, name, font interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateImgSignatureData", reflect.TypeOf((*MockImages)(nil).CreateImgSignatureData), input, name, font)
}

// CreateLatinSignatures mocks base method.
func (m *MockImages) CreateLatinSignatures(user models0.User, id, font string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLatinSignatures", user, id, font)
	ret0, _ := ret[0].(string)
	return ret0
}

// CreateLatinSignatures indicates an expected call of CreateLatinSignatures.
func (mr *MockImagesMockRecorder) CreateLatinSignatures(user, id, font interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLatinSignatures", reflect.TypeOf((*MockImages)(nil).CreateLatinSignatures), user, id, font)
}

// CreateLatinSignaturesData mocks base method.
func (m *MockImages) CreateLatinSignaturesData(user models0.User, latin, idn, font string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLatinSignaturesData", user, latin, idn, font)
	ret0, _ := ret[0].(string)
	return ret0
}

// CreateLatinSignaturesData indicates an expected call of CreateLatinSignaturesData.
func (mr *MockImagesMockRecorder) CreateLatinSignaturesData(user, latin, idn, font interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLatinSignaturesData", reflect.TypeOf((*MockImages)(nil).CreateLatinSignaturesData), user, latin, idn, font)
}

// ResizeImages mocks base method.
func (m *MockImages) ResizeImages(mysign models.MySignatures, input models.SignDocuments) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResizeImages", mysign, input)
	ret0, _ := ret[0].(string)
	return ret0
}

// ResizeImages indicates an expected call of ResizeImages.
func (mr *MockImagesMockRecorder) ResizeImages(mysign, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResizeImages", reflect.TypeOf((*MockImages)(nil).ResizeImages), mysign, input)
}
