package mock_service

import (
	models "e-signature/modules/v1/utilities/user/models"
	reflect "reflect"

	common "github.com/ethereum/go-ethereum/common"
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

// GetCardDashboard mocks base method.
func (m *MockService) GetCardDashboard(sign_id string) models.CardDashboard {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCardDashboard", sign_id)
	ret0, _ := ret[0].(models.CardDashboard)
	return ret0
}

// GetCardDashboard indicates an expected call of GetCardDashboard.
func (mr *MockServiceMockRecorder) GetCardDashboard(sign_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCardDashboard", reflect.TypeOf((*MockService)(nil).GetCardDashboard), sign_id)
}

// GetFileIPFS mocks base method.
func (m *MockService) GetFileIPFS(hash, output, directory string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileIPFS", hash, output, directory)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileIPFS indicates an expected call of GetFileIPFS.
func (mr *MockServiceMockRecorder) GetFileIPFS(hash, output, directory interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileIPFS", reflect.TypeOf((*MockService)(nil).GetFileIPFS), hash, output, directory)
}

// GetLogUser mocks base method.
func (m *MockService) GetLogUser(idsignature string) ([]models.UserLog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogUser", idsignature)
	ret0, _ := ret[0].([]models.UserLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogUser indicates an expected call of GetLogUser.
func (mr *MockServiceMockRecorder) GetLogUser(idsignature interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogUser", reflect.TypeOf((*MockService)(nil).GetLogUser), idsignature)
}

// GetPublicKey mocks base method.
func (m *MockService) GetPublicKey(email []string) ([]common.Address, []string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicKey", email)
	ret0, _ := ret[0].([]common.Address)
	ret1, _ := ret[1].([]string)
	return ret0, ret1
}

// GetPublicKey indicates an expected call of GetPublicKey.
func (mr *MockServiceMockRecorder) GetPublicKey(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicKey", reflect.TypeOf((*MockService)(nil).GetPublicKey), email)
}

// GetUserByEmail mocks base method.
func (m *MockService) GetUserByEmail(email string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", email)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockServiceMockRecorder) GetUserByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockService)(nil).GetUserByEmail), email)
}

// Logging mocks base method.
func (m *MockService) Logging(action, idsignature, ip, user_agent string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logging", action, idsignature, ip, user_agent)
	ret0, _ := ret[0].(error)
	return ret0
}

// Logging indicates an expected call of Logging.
func (mr *MockServiceMockRecorder) Logging(action, idsignature, ip, user_agent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logging", reflect.TypeOf((*MockService)(nil).Logging), action, idsignature, ip, user_agent)
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
func (m *MockService) UploadIPFS(path string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadIPFS", path)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadIPFS indicates an expected call of UploadIPFS.
func (mr *MockServiceMockRecorder) UploadIPFS(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadIPFS", reflect.TypeOf((*MockService)(nil).UploadIPFS), path)
}
