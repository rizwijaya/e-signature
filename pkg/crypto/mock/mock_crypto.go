package mock_crypto

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCrypto is a mock of Crypto interface.
type MockCrypto struct {
	ctrl     *gomock.Controller
	recorder *MockCryptoMockRecorder
}

// MockCryptoMockRecorder is the mock recorder for MockCrypto.
type MockCryptoMockRecorder struct {
	mock *MockCrypto
}

// NewMockCrypto creates a new mock instance.
func NewMockCrypto(ctrl *gomock.Controller) *MockCrypto {
	mock := &MockCrypto{ctrl: ctrl}
	mock.recorder = &MockCryptoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCrypto) EXPECT() *MockCryptoMockRecorder {
	return m.recorder
}

// Compare mocks base method.
func (m *MockCrypto) Compare(hash, pw string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Compare", hash, pw)
	ret0, _ := ret[0].(error)
	return ret0
}

// Compare indicates an expected call of Compare.
func (mr *MockCryptoMockRecorder) Compare(hash, pw interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Compare", reflect.TypeOf((*MockCrypto)(nil).Compare), hash, pw)
}

// CreateKey mocks base method.
func (m *MockCrypto) CreateKey(key string) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateKey", key)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// CreateKey indicates an expected call of CreateKey.
func (mr *MockCryptoMockRecorder) CreateKey(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateKey", reflect.TypeOf((*MockCrypto)(nil).CreateKey), key)
}

// Decrypt mocks base method.
func (m *MockCrypto) Decrypt(data []byte, passphrase string) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decrypt", data, passphrase)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Decrypt indicates an expected call of Decrypt.
func (mr *MockCryptoMockRecorder) Decrypt(data, passphrase interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decrypt", reflect.TypeOf((*MockCrypto)(nil).Decrypt), data, passphrase)
}

// DecryptFile mocks base method.
func (m *MockCrypto) DecryptFile(filename, passphrase string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecryptFile", filename, passphrase)
	ret0, _ := ret[0].(error)
	return ret0
}

// DecryptFile indicates an expected call of DecryptFile.
func (mr *MockCryptoMockRecorder) DecryptFile(filename, passphrase interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecryptFile", reflect.TypeOf((*MockCrypto)(nil).DecryptFile), filename, passphrase)
}

// Encrypt mocks base method.
func (m *MockCrypto) Encrypt(data []byte, passphrase string) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encrypt", data, passphrase)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Encrypt indicates an expected call of Encrypt.
func (mr *MockCryptoMockRecorder) Encrypt(data, passphrase interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encrypt", reflect.TypeOf((*MockCrypto)(nil).Encrypt), data, passphrase)
}

// EncryptFile mocks base method.
func (m *MockCrypto) EncryptFile(filename, passphrase string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EncryptFile", filename, passphrase)
	ret0, _ := ret[0].(error)
	return ret0
}

// EncryptFile indicates an expected call of EncryptFile.
func (mr *MockCryptoMockRecorder) EncryptFile(filename, passphrase interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EncryptFile", reflect.TypeOf((*MockCrypto)(nil).EncryptFile), filename, passphrase)
}

// GenerateHash mocks base method.
func (m *MockCrypto) GenerateHash(pw string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateHash", pw)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateHash indicates an expected call of GenerateHash.
func (mr *MockCryptoMockRecorder) GenerateHash(pw interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateHash", reflect.TypeOf((*MockCrypto)(nil).GenerateHash), pw)
}
