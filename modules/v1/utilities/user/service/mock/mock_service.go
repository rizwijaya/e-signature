package mock_service

import (
	models "e-signature/modules/v1/utilities/user/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	shell "github.com/ipfs/go-ipfs-api"
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

// CheckEmailExist mocks base method.
func (m *MockService) CheckEmailExist(email string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckEmailExist", email)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckEmailExist indicates an expected call of CheckEmailExist.
func (mr *MockServiceMockRecorder) CheckEmailExist(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckEmailExist", reflect.TypeOf((*MockService)(nil).CheckEmailExist), email)
}

// CheckUserExist mocks base method.
func (m *MockService) CheckUserExist(idsignature string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserExist", idsignature)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUserExist indicates an expected call of CheckUserExist.
func (mr *MockServiceMockRecorder) CheckUserExist(idsignature interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserExist", reflect.TypeOf((*MockService)(nil).CheckUserExist), idsignature)
}

// ConnectIPFS mocks base method.
func (m *MockService) ConnectIPFS() *shell.Shell {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectIPFS")
	ret0, _ := ret[0].(*shell.Shell)
	return ret0
}

// ConnectIPFS indicates an expected call of ConnectIPFS.
func (mr *MockServiceMockRecorder) ConnectIPFS() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectIPFS", reflect.TypeOf((*MockService)(nil).ConnectIPFS))
}

// CreateAccount mocks base method.
func (m *MockService) CreateAccount(user models.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockServiceMockRecorder) CreateAccount(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockService)(nil).CreateAccount), user)
}

// CreateKey mocks base method.
func (m *MockService) CreateKey(key string) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateKey", key)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// CreateKey indicates an expected call of CreateKey.
func (mr *MockServiceMockRecorder) CreateKey(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateKey", reflect.TypeOf((*MockService)(nil).CreateKey), key)
}

// Decrypt mocks base method.
func (m *MockService) Decrypt(data []byte, passphrase string) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decrypt", data, passphrase)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Decrypt indicates an expected call of Decrypt.
func (mr *MockServiceMockRecorder) Decrypt(data, passphrase interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decrypt", reflect.TypeOf((*MockService)(nil).Decrypt), data, passphrase)
}

// DecryptFile mocks base method.
func (m *MockService) DecryptFile(filename, passphrase string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecryptFile", filename, passphrase)
	ret0, _ := ret[0].(error)
	return ret0
}

// DecryptFile indicates an expected call of DecryptFile.
func (mr *MockServiceMockRecorder) DecryptFile(filename, passphrase interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecryptFile", reflect.TypeOf((*MockService)(nil).DecryptFile), filename, passphrase)
}

// Encrypt mocks base method.
func (m *MockService) Encrypt(data []byte, passphrase string) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encrypt", data, passphrase)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Encrypt indicates an expected call of Encrypt.
func (mr *MockServiceMockRecorder) Encrypt(data, passphrase interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encrypt", reflect.TypeOf((*MockService)(nil).Encrypt), data, passphrase)
}

// EncryptFile mocks base method.
func (m *MockService) EncryptFile(filename, passphrase string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EncryptFile", filename, passphrase)
	ret0, _ := ret[0].(error)
	return ret0
}

// EncryptFile indicates an expected call of EncryptFile.
func (mr *MockServiceMockRecorder) EncryptFile(filename, passphrase interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EncryptFile", reflect.TypeOf((*MockService)(nil).EncryptFile), filename, passphrase)
}

// GetBalance mocks base method.
func (m *MockService) GetBalance(user models.ProfileDB, pw string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", user, pw)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockServiceMockRecorder) GetBalance(user, pw interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockService)(nil).GetBalance), user, pw)
}

// GetFileIPFS mocks base method.
func (m *MockService) GetFileIPFS(hash, output string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileIPFS", hash, output)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileIPFS indicates an expected call of GetFileIPFS.
func (mr *MockServiceMockRecorder) GetFileIPFS(hash, output interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileIPFS", reflect.TypeOf((*MockService)(nil).GetFileIPFS), hash, output)
}

// Login mocks base method.
func (m *MockService) Login(input models.LoginInput) (models.ProfileDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", input)
	ret0, _ := ret[0].(models.ProfileDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockServiceMockRecorder) Login(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockService)(nil).Login), input)
}

// TransferBalance mocks base method.
func (m *MockService) TransferBalance(user models.ProfileDB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferBalance", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// TransferBalance indicates an expected call of TransferBalance.
func (mr *MockServiceMockRecorder) TransferBalance(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferBalance", reflect.TypeOf((*MockService)(nil).TransferBalance), user)
}

// UploadIPFS mocks base method.
func (m *MockService) UploadIPFS(path string) (error, string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadIPFS", path)
	ret0, _ := ret[0].(error)
	ret1, _ := ret[1].(string)
	return ret0, ret1
}

// UploadIPFS indicates an expected call of UploadIPFS.
func (mr *MockServiceMockRecorder) UploadIPFS(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadIPFS", reflect.TypeOf((*MockService)(nil).UploadIPFS), path)
}
