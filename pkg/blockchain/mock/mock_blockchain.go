package mock_blockchain

import (
	models "e-signature/modules/v1/utilities/signatures/models"
	models0 "e-signature/modules/v1/utilities/user/models"
	big "math/big"
	reflect "reflect"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	common "github.com/ethereum/go-ethereum/common"
	types "github.com/ethereum/go-ethereum/core/types"
	gomock "github.com/golang/mock/gomock"
)

// MockBlockchain is a mock of Blockchain interface.
type MockBlockchain struct {
	ctrl     *gomock.Controller
	recorder *MockBlockchainMockRecorder
}

// MockBlockchainMockRecorder is the mock recorder for MockBlockchain.
type MockBlockchainMockRecorder struct {
	mock *MockBlockchain
}

// NewMockBlockchain creates a new mock instance.
func NewMockBlockchain(ctrl *gomock.Controller) *MockBlockchain {
	mock := &MockBlockchain{ctrl: ctrl}
	mock.recorder = &MockBlockchainMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlockchain) EXPECT() *MockBlockchainMockRecorder {
	return m.recorder
}

// AddToBlockhain mocks base method.
func (m *MockBlockchain) AddToBlockhain(input models.SignDocuments, times *big.Int) (*types.Transaction, *bind.TransactOpts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToBlockhain", input, times)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(*bind.TransactOpts)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddToBlockhain indicates an expected call of AddToBlockhain.
func (mr *MockBlockchainMockRecorder) AddToBlockhain(input, times interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToBlockhain", reflect.TypeOf((*MockBlockchain)(nil).AddToBlockhain), input, times)
}

// DocumentSigned mocks base method.
func (m *MockBlockchain) DocumentSigned(sign models.SignDocs, timeSign *big.Int) (*types.Transaction, *bind.TransactOpts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DocumentSigned", sign, timeSign)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(*bind.TransactOpts)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DocumentSigned indicates an expected call of DocumentSigned.
func (mr *MockBlockchainMockRecorder) DocumentSigned(sign, timeSign interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DocumentSigned", reflect.TypeOf((*MockBlockchain)(nil).DocumentSigned), sign, timeSign)
}

// GeneratePublicKey mocks base method.
func (m *MockBlockchain) GeneratePublicKey(user models0.User) (models0.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeneratePublicKey", user)
	ret0, _ := ret[0].(models0.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GeneratePublicKey indicates an expected call of GeneratePublicKey.
func (mr *MockBlockchainMockRecorder) GeneratePublicKey(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeneratePublicKey", reflect.TypeOf((*MockBlockchain)(nil).GeneratePublicKey), user)
}

// GetBalance mocks base method.
func (m *MockBlockchain) GetBalance(user models0.ProfileDB, pw string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", user, pw)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockBlockchainMockRecorder) GetBalance(user, pw interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockBlockchain)(nil).GetBalance), user, pw)
}

// GetDocument mocks base method.
func (m *MockBlockchain) GetDocument(hash, publickey string) models.DocumentBlockchain {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDocument", hash, publickey)
	ret0, _ := ret[0].(models.DocumentBlockchain)
	return ret0
}

// GetDocument indicates an expected call of GetDocument.
func (mr *MockBlockchainMockRecorder) GetDocument(hash, publickey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDocument", reflect.TypeOf((*MockBlockchain)(nil).GetDocument), hash, publickey)
}

// GetHashOriginal mocks base method.
func (m *MockBlockchain) GetHashOriginal(hash, publickey string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHashOriginal", hash, publickey)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetHashOriginal indicates an expected call of GetHashOriginal.
func (mr *MockBlockchainMockRecorder) GetHashOriginal(hash, publickey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHashOriginal", reflect.TypeOf((*MockBlockchain)(nil).GetHashOriginal), hash, publickey)
}

// GetListSign mocks base method.
func (m *MockBlockchain) GetListSign(hash string) []common.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListSign", hash)
	ret0, _ := ret[0].([]common.Address)
	return ret0
}

// GetListSign indicates an expected call of GetListSign.
func (mr *MockBlockchainMockRecorder) GetListSign(hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListSign", reflect.TypeOf((*MockBlockchain)(nil).GetListSign), hash)
}

// GetPrivateKey mocks base method.
func (m *MockBlockchain) GetPrivateKey(user models0.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrivateKey", user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPrivateKey indicates an expected call of GetPrivateKey.
func (mr *MockBlockchainMockRecorder) GetPrivateKey(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrivateKey", reflect.TypeOf((*MockBlockchain)(nil).GetPrivateKey), user)
}

// GetSigners mocks base method.
func (m *MockBlockchain) GetSigners(hash, publickey string) models.Signers {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSigners", hash, publickey)
	ret0, _ := ret[0].(models.Signers)
	return ret0
}

// GetSigners indicates an expected call of GetSigners.
func (mr *MockBlockchainMockRecorder) GetSigners(hash, publickey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSigners", reflect.TypeOf((*MockBlockchain)(nil).GetSigners), hash, publickey)
}

// TransferBalance mocks base method.
func (m *MockBlockchain) TransferBalance(user models0.ProfileDB) (string, string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferBalance", user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// TransferBalance indicates an expected call of TransferBalance.
func (mr *MockBlockchainMockRecorder) TransferBalance(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferBalance", reflect.TypeOf((*MockBlockchain)(nil).TransferBalance), user)
}

// VerifyDoc mocks base method.
func (m *MockBlockchain) VerifyDoc(hash string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyDoc", hash)
	ret0, _ := ret[0].(bool)
	return ret0
}

// VerifyDoc indicates an expected call of VerifyDoc.
func (mr *MockBlockchainMockRecorder) VerifyDoc(hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyDoc", reflect.TypeOf((*MockBlockchain)(nil).VerifyDoc), hash)
}
