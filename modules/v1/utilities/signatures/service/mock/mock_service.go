package mock_service

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

// AddToBlockhain mocks base method.
func (m *MockService) AddToBlockhain(input models.SignDocuments) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToBlockhain", input)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToBlockhain indicates an expected call of AddToBlockhain.
func (mr *MockServiceMockRecorder) AddToBlockhain(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToBlockhain", reflect.TypeOf((*MockService)(nil).AddToBlockhain), input)
}

// AddUserDocs mocks base method.
func (m *MockService) AddUserDocs(input models.SignDocuments) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUserDocs", input)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUserDocs indicates an expected call of AddUserDocs.
func (mr *MockServiceMockRecorder) AddUserDocs(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUserDocs", reflect.TypeOf((*MockService)(nil).AddUserDocs), input)
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

// CheckSignature mocks base method.
func (m *MockService) CheckSignature(hash, publickey string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckSignature", hash, publickey)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckSignature indicates an expected call of CheckSignature.
func (mr *MockServiceMockRecorder) CheckSignature(hash, publickey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSignature", reflect.TypeOf((*MockService)(nil).CheckSignature), hash, publickey)
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

// DocumentSigned mocks base method.
func (m *MockService) DocumentSigned(sign models.SignDocs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DocumentSigned", sign)
	ret0, _ := ret[0].(error)
	return ret0
}

// DocumentSigned indicates an expected call of DocumentSigned.
func (mr *MockServiceMockRecorder) DocumentSigned(sign interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DocumentSigned", reflect.TypeOf((*MockService)(nil).DocumentSigned), sign)
}

// GenerateHashDocument mocks base method.
func (m *MockService) GenerateHashDocument(input string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateHashDocument", input)
	ret0, _ := ret[0].(string)
	return ret0
}

// GenerateHashDocument indicates an expected call of GenerateHashDocument.
func (mr *MockServiceMockRecorder) GenerateHashDocument(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateHashDocument", reflect.TypeOf((*MockService)(nil).GenerateHashDocument), input)
}

// GetDocument mocks base method.
func (m *MockService) GetDocument(hash, publickey string) models.DocumentBlockchain {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDocument", hash, publickey)
	ret0, _ := ret[0].(models.DocumentBlockchain)
	return ret0
}

// GetDocument indicates an expected call of GetDocument.
func (mr *MockServiceMockRecorder) GetDocument(hash, publickey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDocument", reflect.TypeOf((*MockService)(nil).GetDocument), hash, publickey)
}

// GetDocumentAllSign mocks base method.
func (m *MockService) GetDocumentAllSign(hash string) (models.DocumentAllSign, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDocumentAllSign", hash)
	ret0, _ := ret[0].(models.DocumentAllSign)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetDocumentAllSign indicates an expected call of GetDocumentAllSign.
func (mr *MockServiceMockRecorder) GetDocumentAllSign(hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDocumentAllSign", reflect.TypeOf((*MockService)(nil).GetDocumentAllSign), hash)
}

// GetDocumentNoSigners mocks base method.
func (m *MockService) GetDocumentNoSigners(hash string) models.DocumentBlockchain {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDocumentNoSigners", hash)
	ret0, _ := ret[0].(models.DocumentBlockchain)
	return ret0
}

// GetDocumentNoSigners indicates an expected call of GetDocumentNoSigners.
func (mr *MockServiceMockRecorder) GetDocumentNoSigners(hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDocumentNoSigners", reflect.TypeOf((*MockService)(nil).GetDocumentNoSigners), hash)
}

// GetListDocument mocks base method.
func (m *MockService) GetListDocument(publickey string) []models.ListDocument {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListDocument", publickey)
	ret0, _ := ret[0].([]models.ListDocument)
	return ret0
}

// GetListDocument indicates an expected call of GetListDocument.
func (mr *MockServiceMockRecorder) GetListDocument(publickey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListDocument", reflect.TypeOf((*MockService)(nil).GetListDocument), publickey)
}

// GetMySignature mocks base method.
func (m *MockService) GetMySignature(sign, id, name string) models.MySignatures {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMySignature", sign, id, name)
	ret0, _ := ret[0].(models.MySignatures)
	return ret0
}

// GetMySignature indicates an expected call of GetMySignature.
func (mr *MockServiceMockRecorder) GetMySignature(sign, id, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMySignature", reflect.TypeOf((*MockService)(nil).GetMySignature), sign, id, name)
}

// GetTransactions mocks base method.
func (m *MockService) GetTransactions() []models.Transac {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactions")
	ret0, _ := ret[0].([]models.Transac)
	return ret0
}

// GetTransactions indicates an expected call of GetTransactions.
func (mr *MockServiceMockRecorder) GetTransactions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactions", reflect.TypeOf((*MockService)(nil).GetTransactions))
}

// InvitePeople mocks base method.
func (m *MockService) InvitePeople(email string, input models.SignDocuments, users models0.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvitePeople", email, input, users)
	ret0, _ := ret[0].(error)
	return ret0
}

// InvitePeople indicates an expected call of InvitePeople.
func (mr *MockServiceMockRecorder) InvitePeople(email, input, users interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvitePeople", reflect.TypeOf((*MockService)(nil).InvitePeople), email, input, users)
}

// ResizeImages mocks base method.
func (m *MockService) ResizeImages(mysign models.MySignatures, input models.SignDocuments) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResizeImages", mysign, input)
	ret0, _ := ret[0].(string)
	return ret0
}

// ResizeImages indicates an expected call of ResizeImages.
func (mr *MockServiceMockRecorder) ResizeImages(mysign, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResizeImages", reflect.TypeOf((*MockService)(nil).ResizeImages), mysign, input)
}

// SignDocuments mocks base method.
func (m *MockService) SignDocuments(imgpath string, input models.SignDocuments) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignDocuments", imgpath, input)
	ret0, _ := ret[0].(string)
	return ret0
}

// SignDocuments indicates an expected call of SignDocuments.
func (mr *MockServiceMockRecorder) SignDocuments(imgpath, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignDocuments", reflect.TypeOf((*MockService)(nil).SignDocuments), imgpath, input)
}

// TimeFormating mocks base method.
func (m *MockService) TimeFormating(times string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TimeFormating", times)
	ret0, _ := ret[0].(string)
	return ret0
}

// TimeFormating indicates an expected call of TimeFormating.
func (mr *MockServiceMockRecorder) TimeFormating(times interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TimeFormating", reflect.TypeOf((*MockService)(nil).TimeFormating), times)
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
