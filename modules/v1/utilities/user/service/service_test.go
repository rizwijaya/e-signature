package service

import (
	"e-signature/modules/v1/utilities/user/models"
	m_repo "e-signature/modules/v1/utilities/user/repository/mock"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/tkuchiki/faketime"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInit(t *testing.T) {
	gin.SetMode(gin.TestMode)
	os.Chdir("../../../../../")
}

func Test_service_NewService(t *testing.T) {
	t.Run("Case 1: Success New Service", func(t *testing.T) {
		serv := NewService(nil)
		assert.NotNil(t, serv)
	})
}

func Test_service_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var err error
	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()
	location, err := time.LoadLocation("Asia/Jakarta")
	assert.NoError(t, err)
	times := time.Now().In(location)

	test := []struct {
		name     string
		input    models.LoginInput
		repoTest func(repo *m_repo.MockRepository)
	}{
		{
			name: "Login Service Case 1: Failed Login User Not Found",
			input: models.LoginInput{
				IdSignature: "rizwijaya",
				Password:    "123456",
			},
			repoTest: func(repo *m_repo.MockRepository) {
				repo.EXPECT().CheckUserExist("rizwijaya").Return(models.ProfileDB{}, nil).Times(1)
			},
		},
		{
			name: "Login Service Case 2: Failed Login Password Not Valid",
			input: models.LoginInput{
				IdSignature: "rizwijaya",
				Password:    "123",
			},
			repoTest: func(repo *m_repo.MockRepository) {
				repo.EXPECT().CheckUserExist("rizwijaya").Return(models.ProfileDB{
					Id:            primitive.NewObjectID(),
					Idsignature:   "rizwijaya",
					Name:          "Rizqi Wijaya",
					Email:         "smartsign@rizwijaya.com",
					Phone:         "081234567890",
					Identity_card: "a1ddfs2s3fes4s5s2aas3sdsdaad167890",
					Password:      "123456",
					PublicKey:     "sjn729nad6a9804jd7tnfs8bqvi3iaewc9y80c",
					Role_id:       2,
					Date_created:  times,
				}, nil).Times(1)
			},
		},
		{
			name: "Login Service Case 3: Login Success",
			input: models.LoginInput{
				IdSignature: "rizwijaya",
				Password:    "rizwijaya123",
			},
			repoTest: func(repo *m_repo.MockRepository) {
				repo.EXPECT().CheckUserExist("rizwijaya").Return(models.ProfileDB{
					Id:            primitive.NewObjectID(),
					Idsignature:   "rizwijaya",
					Name:          "Rizqi Wijaya",
					Email:         "smartsign@rizwijaya.com",
					Phone:         "081234567890",
					Identity_card: "a1ddfs2s3fes4s5s2aas3sdsdaad167890",
					Password:      "$2a$04$mL8CVXcMKPOINfUgBszbgupxCC9lj0eqPvnnz/iNng/CisnSjdYdu",
					PublicKey:     "sjn729nad6a9804jd7tnfs8bqvi3iaewc9y80c",
					Role_id:       2,
					Date_created:  times,
				}, nil).Times(1)
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			repo := m_repo.NewMockRepository(ctrl)
			tt.repoTest(repo)
			s := NewService(repo)
			_, err := s.Login(tt.input)
			if err == errors.New("user not found") {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_service_UploadIPFS(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("Upload File IPFS Service Case 1: Upload File in IPFS", func(t *testing.T) {
		repo := m_repo.NewMockRepository(ctrl)
		s := &service{
			repository: repo,
		}
		cidr, _ := s.UploadIPFS("")
		assert.Equal(t, cidr, "")
	})
}

func Test_service_GetFileIPFS(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("Get File IPFS Service Case 1: Get File in IPFS", func(t *testing.T) {
		repo := m_repo.NewMockRepository(ctrl)
		s := &service{
			repository: repo,
		}
		cidr, _ := s.GetFileIPFS("", "", "")
		assert.Equal(t, cidr, "")
	})
}

// func Test_service_CreateAccount(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	//var err error
// 	// f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
// 	// defer f.Undo()
// 	// f.Do()
// 	// location, err := time.LoadLocation("Asia/Jakarta")
// 	// assert.NoError(t, err)
// 	// times := time.Now().In(location)
// 	id := primitive.NewObjectID()
// 	test := []struct {
// 		name     string
// 		id       string
// 		err      error
// 		user     models.User
// 		repoTest func(repo *m_repo.MockRepository)
// 	}{
// 		{
// 			//6380b5cbdc938c5fdf8e6bfe
// 			name: "Register Service Case 1: Failed Generated Password",
// 			id:   "",
// 			err:  errors.New("password must be more than 6 characters"),
// 			user: models.User{
// 				Id:          id,
// 				Idsignature: "rizwijaya",
// 				Password:    "",
// 			},
// 		},
// 		{
// 			//6380b5cbdc938c5fdf8e6bfe
// 			name: "Register Service Case 2: Failed Generated Public Key",
// 			id:   "",
// 			err:  nil,
// 			user: models.User{
// 				Id:          id,
// 				Idsignature: "rizwijaya",
// 				Password:    "1234567",
// 			},
// 			repoTest: func(repo *m_repo.MockRepository) {
// 				user := models.User{
// 					Id:          id,
// 					Idsignature: "rizwijaya",
// 					Password:    "1234567",
// 					// PasswordHash: "$2a$04$8vhqEvMjmX0ywxFdQV2geeLKNEgzE1WEqHNYWlGHWSJbn0cZsEjbi",
// 				}
// 				user = repo.GeneratePublicKey(user.Password)
// 				//repo.EXPECT().GeneratePublicKey(user.Password).Times(1)
// 			},
// 		},
// 	}

// 	for _, tt := range test {
// 		t.Run(tt.name, func(t *testing.T) {
// 			repo := m_repo.NewMockRepository(ctrl)
// 			if tt.repoTest != nil {
// 				tt.repoTest(repo)
// 			}
// 			s := &service{
// 				repository: repo,
// 			}
// 			id, err := s.CreateAccount(tt.user)
// 			assert.Equal(t, tt.id, id)
// 			if tt.err != nil {
// 				assert.Equal(t, tt.err, err)
// 			} else {
// 				assert.NoError(t, err)
// 			}
// 		})
// 	}
// }

func Test_service_Encrypt(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("Encrypt Service Case 1: Encryption Success", func(t *testing.T) {
		repo := m_repo.NewMockRepository(ctrl)
		s := &service{
			repository: repo,
		}
		key := s.Encrypt([]byte{0x5e, 0x88, 0x48, 0x98, 0xda, 0x28, 0x4, 0x71, 0x51, 0xd0, 0xe5, 0x6f, 0x8d, 0xc6, 0x29, 0x27, 0x73, 0x60, 0x3d, 0xd, 0x6a, 0xab, 0xbd, 0xd6, 0x2a, 0x11, 0xef, 0x72, 0x1d, 0x15, 0x42, 0xd8}, "password")
		assert.NotEqual(t, string(key), "")
		fmt.Println(string(key))
	})
}

func Test_service_Decrypt(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	test := []struct {
		name string
		data string
		// dataDecrypt string
		passphrase string
		//repoTest   func(repo *m_repo.MockRepository)
	}{
		{
			name: "Decrypt Service Case 1: Decryption Success",
			data: "32V4Ob933XJqS7yOSLdcl5ld5IEdpEmB/7cCZU6yv5N6V4FbnopPxOSt9AqaFRFmFmt5uZGF9lhcNjB4",
			// dataDecrypt: "password",
			passphrase: "password", //password
		},
		{
			name: "Decrypt Service Case 2: Decryption Failed",
			data: "sdadferwwreEasda",
			// dataDecrypt: "",
			passphrase: "password", //password
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			repo := m_repo.NewMockRepository(ctrl)
			s := &service{
				repository: repo,
			}
			s.Decrypt([]byte(tt.data), tt.passphrase)
		})
	}
}
func Test_service_EncryptFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	test := []struct {
		name       string
		fileName   string
		passphrase string
		err        bool
	}{
		{
			name:       "Encrypt File Service Case 1: Encryption Success",
			fileName:   "./public/temp/pdfsign/signed_sample_test.pdf",
			passphrase: "password",
			err:        false,
		},
		{
			name:       "Encrypt File Service Case 2: Encryption Failed",
			fileName:   "./public/temp/pdfsign/signed_test.pdf",
			passphrase: "password",
			err:        true,
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			repo := m_repo.NewMockRepository(ctrl)
			s := &service{
				repository: repo,
			}
			err := s.EncryptFile(tt.fileName, tt.passphrase)
			if !tt.err {
				assert.NoError(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func Test_service_DecryptFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	test := []struct {
		name       string
		fileName   string
		passphrase string
		err        bool
	}{
		{
			name:       "Decrypt File Service Case 1: Decryption Success",
			fileName:   "./public/temp/pdfsign/signed_sample_test.pdf",
			passphrase: "password",
			err:        false,
		},
		{
			name:       "Decrypt File Service Case 2: Decryption Failed",
			fileName:   "./public/temp/pdfsign/signed_test.pdf",
			passphrase: "password",
			err:        true,
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			repo := m_repo.NewMockRepository(ctrl)
			s := &service{
				repository: repo,
			}
			err := s.DecryptFile(tt.fileName, tt.passphrase)
			if !tt.err {
				assert.NoError(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func Test_service_CheckUserExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	test := []struct {
		name        string
		idsignature string
		repoTest    func(repo *m_repo.MockRepository)
	}{
		{
			name:        "Check User Exist Service Case 1: User Exist Success",
			idsignature: "rizwijaya",
			repoTest: func(repo *m_repo.MockRepository) {
				repo.EXPECT().CheckUserExist("rizwijaya").Return(models.ProfileDB{
					Email: "smartsign@rizwijaya.com",
				}, nil).Times(1)
			},
		},
		{
			name:        "Check User Exist Service Case 2: User Not Exist Success",
			idsignature: "smartsign",
			repoTest: func(repo *m_repo.MockRepository) {
				repo.EXPECT().CheckUserExist("smartsign").Return(models.ProfileDB{}, nil).Times(1)
			},
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			repo := m_repo.NewMockRepository(ctrl)
			tt.repoTest(repo)
			s := NewService(repo)
			_, err := s.CheckUserExist(tt.idsignature)
			assert.NoError(t, err)
		})
	}
}
