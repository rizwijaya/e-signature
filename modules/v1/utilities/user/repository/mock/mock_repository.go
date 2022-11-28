package mock_repository

import (
	models "e-signature/modules/v1/utilities/user/models"
	os "os"
	reflect "reflect"

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

// CheckEmailExist mocks base method.
func (m *MockRepository) CheckEmailExist(email string) (models.ProfileDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckEmailExist", email)
	ret0, _ := ret[0].(models.ProfileDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckEmailExist indicates an expected call of CheckEmailExist.
func (mr *MockRepositoryMockRecorder) CheckEmailExist(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckEmailExist", reflect.TypeOf((*MockRepository)(nil).CheckEmailExist), email)
}

// CheckUserExist mocks base method.
func (m *MockRepository) CheckUserExist(idsignature string) (models.ProfileDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserExist", idsignature)
	ret0, _ := ret[0].(models.ProfileDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUserExist indicates an expected call of CheckUserExist.
func (mr *MockRepositoryMockRecorder) CheckUserExist(idsignature interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserExist", reflect.TypeOf((*MockRepository)(nil).CheckUserExist), idsignature)
}

// GeneratePublicKey mocks base method.
func (m *MockRepository) GeneratePublicKey(user models.User) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeneratePublicKey", user)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GeneratePublicKey indicates an expected call of GeneratePublicKey.
func (mr *MockRepositoryMockRecorder) GeneratePublicKey(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeneratePublicKey", reflect.TypeOf((*MockRepository)(nil).GeneratePublicKey), user)
}

// GetBalance mocks base method.
func (m *MockRepository) GetBalance(user models.ProfileDB, pw string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", user, pw)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockRepositoryMockRecorder) GetBalance(user, pw interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockRepository)(nil).GetBalance), user, pw)
}

// GetLogUser mocks base method.
func (m *MockRepository) GetLogUser(idsignature string) ([]models.UserLog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogUser", idsignature)
	ret0, _ := ret[0].([]models.UserLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogUser indicates an expected call of GetLogUser.
func (mr *MockRepositoryMockRecorder) GetLogUser(idsignature interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogUser", reflect.TypeOf((*MockRepository)(nil).GetLogUser), idsignature)
}

// GetPrivateKey mocks base method.
func (m *MockRepository) GetPrivateKey(user models.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrivateKey", user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPrivateKey indicates an expected call of GetPrivateKey.
func (mr *MockRepositoryMockRecorder) GetPrivateKey(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrivateKey", reflect.TypeOf((*MockRepository)(nil).GetPrivateKey), user)
}

// GetTotal mocks base method.
func (m *MockRepository) GetTotal(db string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotal", db)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetTotal indicates an expected call of GetTotal.
func (mr *MockRepositoryMockRecorder) GetTotal(db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotal", reflect.TypeOf((*MockRepository)(nil).GetTotal), db)
}

// GetTotalRequestUser mocks base method.
func (m *MockRepository) GetTotalRequestUser(sign_id string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalRequestUser", sign_id)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetTotalRequestUser indicates an expected call of GetTotalRequestUser.
func (mr *MockRepositoryMockRecorder) GetTotalRequestUser(sign_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalRequestUser", reflect.TypeOf((*MockRepository)(nil).GetTotalRequestUser), sign_id)
}

// GetUserByEmail mocks base method.
func (m *MockRepository) GetUserByEmail(email string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", email)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockRepositoryMockRecorder) GetUserByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockRepository)(nil).GetUserByEmail), email)
}

// Logging mocks base method.
func (m *MockRepository) Logging(logg models.UserLog) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logging", logg)
	ret0, _ := ret[0].(error)
	return ret0
}

// Logging indicates an expected call of Logging.
func (mr *MockRepositoryMockRecorder) Logging(logg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logging", reflect.TypeOf((*MockRepository)(nil).Logging), logg)
}

// Register mocks base method.
func (m *MockRepository) Register(user models.User) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", user)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockRepositoryMockRecorder) Register(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockRepository)(nil).Register), user)
}

// SearchFile mocks base method.
func (m *MockRepository) SearchFile(path string, info os.FileInfo, err error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchFile", path, info, err)
	ret0, _ := ret[0].(error)
	return ret0
}

// SearchFile indicates an expected call of SearchFile.
func (mr *MockRepositoryMockRecorder) SearchFile(path, info, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchFile", reflect.TypeOf((*MockRepository)(nil).SearchFile), path, info, err)
}

// TransferBalance mocks base method.
func (m *MockRepository) TransferBalance(user models.ProfileDB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferBalance", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// TransferBalance indicates an expected call of TransferBalance.
func (mr *MockRepositoryMockRecorder) TransferBalance(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferBalance", reflect.TypeOf((*MockRepository)(nil).TransferBalance), user)
}
