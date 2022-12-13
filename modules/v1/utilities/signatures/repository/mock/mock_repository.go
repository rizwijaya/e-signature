package mock_repository

import (
	models "e-signature/modules/v1/utilities/signatures/models"
	models0 "e-signature/modules/v1/utilities/user/models"
	big "math/big"
	reflect "reflect"

	common "github.com/ethereum/go-ethereum/common"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// AddToBlockhain mocks base method.
func (m *MockRepository) AddToBlockhain(input models.SignDocuments, times *big.Int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToBlockhain", input, times)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToBlockhain indicates an expected call of AddToBlockhain.
func (mr *MockRepositoryMockRecorder) AddToBlockhain(input, times interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToBlockhain", reflect.TypeOf((*MockRepository)(nil).AddToBlockhain), input, times)
}

// AddUserDocs mocks base method.
func (m *MockRepository) AddUserDocs(input models.SignDocuments) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUserDocs", input)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUserDocs indicates an expected call of AddUserDocs.
func (mr *MockRepositoryMockRecorder) AddUserDocs(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUserDocs", reflect.TypeOf((*MockRepository)(nil).AddUserDocs), input)
}

// ChangeSignature mocks base method.
func (m *MockRepository) ChangeSignature(sign_type, sign string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeSignature", sign_type, sign)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeSignature indicates an expected call of ChangeSignature.
func (mr *MockRepositoryMockRecorder) ChangeSignature(sign_type, sign interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeSignature", reflect.TypeOf((*MockRepository)(nil).ChangeSignature), sign_type, sign)
}

// CheckSignature mocks base method.
func (m *MockRepository) CheckSignature(hash, publickey string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckSignature", hash, publickey)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckSignature indicates an expected call of CheckSignature.
func (mr *MockRepositoryMockRecorder) CheckSignature(hash, publickey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSignature", reflect.TypeOf((*MockRepository)(nil).CheckSignature), hash, publickey)
}

// DefaultSignatures mocks base method.
func (m *MockRepository) DefaultSignatures(user models0.User, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultSignatures", user, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DefaultSignatures indicates an expected call of DefaultSignatures.
func (mr *MockRepositoryMockRecorder) DefaultSignatures(user, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultSignatures", reflect.TypeOf((*MockRepository)(nil).DefaultSignatures), user, id)
}

// DocumentSigned mocks base method.
func (m *MockRepository) DocumentSigned(sign models.SignDocs, timeSign *big.Int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DocumentSigned", sign, timeSign)
	ret0, _ := ret[0].(error)
	return ret0
}

// DocumentSigned indicates an expected call of DocumentSigned.
func (mr *MockRepositoryMockRecorder) DocumentSigned(sign, timeSign interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DocumentSigned", reflect.TypeOf((*MockRepository)(nil).DocumentSigned), sign, timeSign)
}

// GetDocument mocks base method.
func (m *MockRepository) GetDocument(hash, publickey string) models.DocumentBlockchain {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDocument", hash, publickey)
	ret0, _ := ret[0].(models.DocumentBlockchain)
	return ret0
}

// GetDocument indicates an expected call of GetDocument.
func (mr *MockRepositoryMockRecorder) GetDocument(hash, publickey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDocument", reflect.TypeOf((*MockRepository)(nil).GetDocument), hash, publickey)
}

// GetHashOriginal mocks base method.
func (m *MockRepository) GetHashOriginal(hash, publickey string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHashOriginal", hash, publickey)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetHashOriginal indicates an expected call of GetHashOriginal.
func (mr *MockRepositoryMockRecorder) GetHashOriginal(hash, publickey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHashOriginal", reflect.TypeOf((*MockRepository)(nil).GetHashOriginal), hash, publickey)
}

// GetListSign mocks base method.
func (m *MockRepository) GetListSign(hash string) []common.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListSign", hash)
	ret0, _ := ret[0].([]common.Address)
	return ret0
}

// GetListSign indicates an expected call of GetListSign.
func (mr *MockRepositoryMockRecorder) GetListSign(hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListSign", reflect.TypeOf((*MockRepository)(nil).GetListSign), hash)
}

// GetMySignature mocks base method.
func (m *MockRepository) GetMySignature(sign string) (models.Signatures, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMySignature", sign)
	ret0, _ := ret[0].(models.Signatures)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMySignature indicates an expected call of GetMySignature.
func (mr *MockRepositoryMockRecorder) GetMySignature(sign interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMySignature", reflect.TypeOf((*MockRepository)(nil).GetMySignature), sign)
}

// GetSigners mocks base method.
func (m *MockRepository) GetSigners(hash, publickey string) models.Signers {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSigners", hash, publickey)
	ret0, _ := ret[0].(models.Signers)
	return ret0
}

// GetSigners indicates an expected call of GetSigners.
func (mr *MockRepositoryMockRecorder) GetSigners(hash, publickey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSigners", reflect.TypeOf((*MockRepository)(nil).GetSigners), hash, publickey)
}

// GetTransactions mocks base method.
func (m *MockRepository) GetTransactions() []models.Transac {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactions")
	ret0, _ := ret[0].([]models.Transac)
	return ret0
}

// GetTransactions indicates an expected call of GetTransactions.
func (mr *MockRepositoryMockRecorder) GetTransactions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactions", reflect.TypeOf((*MockRepository)(nil).GetTransactions))
}

// GetUserByIdSignatures mocks base method.
func (m *MockRepository) GetUserByIdSignatures(idsignature string) models0.ProfileDB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByIdSignatures", idsignature)
	ret0, _ := ret[0].(models0.ProfileDB)
	return ret0
}

// GetUserByIdSignatures indicates an expected call of GetUserByIdSignatures.
func (mr *MockRepositoryMockRecorder) GetUserByIdSignatures(idsignature interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByIdSignatures", reflect.TypeOf((*MockRepository)(nil).GetUserByIdSignatures), idsignature)
}

// ListDocumentNoSign mocks base method.
func (m *MockRepository) ListDocumentNoSign(publickey string) []models.ListDocument {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDocumentNoSign", publickey)
	ret0, _ := ret[0].([]models.ListDocument)
	return ret0
}

// ListDocumentNoSign indicates an expected call of ListDocumentNoSign.
func (mr *MockRepositoryMockRecorder) ListDocumentNoSign(publickey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDocumentNoSign", reflect.TypeOf((*MockRepository)(nil).ListDocumentNoSign), publickey)
}

// LogTransactions mocks base method.
func (m *MockRepository) LogTransactions(address, tx_hash, nonce, desc, prices string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogTransactions", address, tx_hash, nonce, desc, prices)
	ret0, _ := ret[0].(error)
	return ret0
}

// LogTransactions indicates an expected call of LogTransactions.
func (mr *MockRepositoryMockRecorder) LogTransactions(address, tx_hash, nonce, desc, prices interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogTransactions", reflect.TypeOf((*MockRepository)(nil).LogTransactions), address, tx_hash, nonce, desc, prices)
}

// UpdateMySignatures mocks base method.
func (m *MockRepository) UpdateMySignatures(signature, signaturedata, sign string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMySignatures", signature, signaturedata, sign)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMySignatures indicates an expected call of UpdateMySignatures.
func (mr *MockRepositoryMockRecorder) UpdateMySignatures(signature, signaturedata, sign interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMySignatures", reflect.TypeOf((*MockRepository)(nil).UpdateMySignatures), signature, signaturedata, sign)
}

// VerifyDoc mocks base method.
func (m *MockRepository) VerifyDoc(hash string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyDoc", hash)
	ret0, _ := ret[0].(bool)
	return ret0
}

// VerifyDoc indicates an expected call of VerifyDoc.
func (mr *MockRepositoryMockRecorder) VerifyDoc(hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyDoc", reflect.TypeOf((*MockRepository)(nil).VerifyDoc), hash)
}
