package mock_document

import (
	models "e-signature/modules/v1/utilities/signatures/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	creator "github.com/unidoc/unipdf/v3/creator"
	model "github.com/unidoc/unipdf/v3/model"
)

// MockDocuments is a mock of Documents interface.
type MockDocuments struct {
	ctrl     *gomock.Controller
	recorder *MockDocumentsMockRecorder
}

// MockDocumentsMockRecorder is the mock recorder for MockDocuments.
type MockDocumentsMockRecorder struct {
	mock *MockDocuments
}

// NewMockDocuments creates a new mock instance.
func NewMockDocuments(ctrl *gomock.Controller) *MockDocuments {
	mock := &MockDocuments{ctrl: ctrl}
	mock.recorder = &MockDocumentsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDocuments) EXPECT() *MockDocumentsMockRecorder {
	return m.recorder
}

// CalcImagePos mocks base method.
func (m *MockDocuments) CalcImagePos(img *creator.Image, page *model.PdfPage, input models.SignDocuments) *creator.Image {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalcImagePos", img, page, input)
	ret0, _ := ret[0].(*creator.Image)
	return ret0
}

// CalcImagePos indicates an expected call of CalcImagePos.
func (mr *MockDocumentsMockRecorder) CalcImagePos(img, page, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalcImagePos", reflect.TypeOf((*MockDocuments)(nil).CalcImagePos), img, page, input)
}

// SignDocuments mocks base method.
func (m *MockDocuments) SignDocuments(imgpath string, input models.SignDocuments) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignDocuments", imgpath, input)
	ret0, _ := ret[0].(string)
	return ret0
}

// SignDocuments indicates an expected call of SignDocuments.
func (mr *MockDocumentsMockRecorder) SignDocuments(imgpath, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignDocuments", reflect.TypeOf((*MockDocuments)(nil).SignDocuments), imgpath, input)
}
