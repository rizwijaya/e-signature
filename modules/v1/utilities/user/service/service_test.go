package service

import (
	"e-signature/modules/v1/utilities/user/models"
	m_repo "e-signature/modules/v1/utilities/user/repository/mock"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
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

func Test_service_CheckEmailExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	test := []struct {
		name     string
		email    string
		repoTest func(repo *m_repo.MockRepository)
	}{
		{
			name:  "Check Email Exist Service Case 1: Email Exist Success",
			email: "smartsign@rizwijaya.com",
			repoTest: func(repo *m_repo.MockRepository) {
				repo.EXPECT().CheckEmailExist("smartsign@rizwijaya.com").Return(models.ProfileDB{
					Email: "smartsign@rizwijaya.com",
				}, nil).Times(1)
			},
		},
		{
			name:  "Check Email Exist Service Case 2: Email Not Exist Success",
			email: "smart@rizwijaya.com",
			repoTest: func(repo *m_repo.MockRepository) {
				repo.EXPECT().CheckEmailExist("smart@rizwijaya.com").Return(models.ProfileDB{}, nil).Times(1)
			},
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			repo := m_repo.NewMockRepository(ctrl)
			tt.repoTest(repo)
			s := NewService(repo)
			_, err := s.CheckEmailExist(tt.email)
			assert.NoError(t, err)
		})
	}
}

func Test_service_GetBalance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	test := []struct {
		name     string
		user     models.ProfileDB
		password string
		balance  string
		erors    error
		repoTest func(repo *m_repo.MockRepository)
	}{
		{
			name:     "Get Balance Service Case 1: Get Balance Success",
			user:     models.ProfileDB{},
			password: "password",
			balance:  "100",
			erors:    nil,
			repoTest: func(repo *m_repo.MockRepository) {
				repo.EXPECT().GetBalance(models.ProfileDB{}, "password").Return("100", nil).Times(1)
			},
		},
		{
			name:     "Get Balance Service Case 2: Get Balance Failed",
			user:     models.ProfileDB{},
			password: "password",
			balance:  "",
			erors:    errors.New("Failed Get Balance from Blockchain"),
			repoTest: func(repo *m_repo.MockRepository) {
				repo.EXPECT().GetBalance(models.ProfileDB{}, "password").Return("", errors.New("Failed Get Balance from Blockchain")).Times(1)
			},
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			repo := m_repo.NewMockRepository(ctrl)
			tt.repoTest(repo)
			s := NewService(repo)
			balance, err := s.GetBalance(tt.user, tt.password)
			assert.Equal(t, tt.erors, err)
			assert.Equal(t, tt.balance, balance)
		})
	}
}

func Test_service_TransferBalance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	test := []struct {
		name     string
		user     models.ProfileDB
		erors    error
		repoTest func(repo *m_repo.MockRepository)
	}{
		{
			name:  "Transfer Balance Service Case 1: Transfer Balance Success",
			user:  models.ProfileDB{},
			erors: nil,
			repoTest: func(repo *m_repo.MockRepository) {
				repo.EXPECT().TransferBalance(models.ProfileDB{}).Return(nil).Times(1)
			},
		},
		{
			name:  "Transfer Balance Service Case 2: Transfer Balance Failed",
			user:  models.ProfileDB{},
			erors: errors.New("Failed Transfer Balance in Blockchain"),
			repoTest: func(repo *m_repo.MockRepository) {
				repo.EXPECT().TransferBalance(models.ProfileDB{}).Return(errors.New("Failed Transfer Balance in Blockchain")).Times(1)
			},
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			repo := m_repo.NewMockRepository(ctrl)
			tt.repoTest(repo)
			s := NewService(repo)
			err := s.TransferBalance(tt.user)
			assert.Equal(t, tt.erors, err)
		})
	}
}

func Test_service_GetPublicKey(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	test := []struct {
		name        string
		email       []string
		idsignature []string
		addr        []common.Address
		repoTest    func(repo *m_repo.MockRepository)
	}{
		{
			name:        "Get Public Key Service Case 1: Get Public Key Success",
			email:       []string{"rizqi@smartsign.com", "admin@smartsign.com"},
			idsignature: []string{"rizqi", "admin"},
			addr:        []common.Address{common.HexToAddress("0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a"), common.HexToAddress("0x3227fc42acAF0C6Ba14A42f8dd518eDfe72cd21D")},
			repoTest: func(repo *m_repo.MockRepository) {
				email := []string{"rizqi@smartsign.com", "admin@smartsign.com"}
				idsignature := []string{"rizqi", "admin"}
				addr := []string{"1uFwqkgHIDZ0oNIYzQyDkSBYPfxf2jdTfq7kLEivwhI68cWQW6jeHD7TnWL6dI4rXYIZWuNfCAhuzkGCcLQHfuJmMTXeMw", "o8Q0o0UAIhmdoIhLD9Y6gqFuSkeBWTFmkM5BsJOc4J9o3gIxqzZHFax/pAW8Fs/hg/qbALGOxyPi0bGdBeCQC15obooKQg"}
				for i := range email {
					repo.EXPECT().GetUserByEmail(email[i]).Return(models.User{
						Publickey:   addr[i],
						Idsignature: idsignature[i],
					}, nil)
				}
			},
		},
		{
			name:        "Get Public Key Service Case 2: Get Public Key Failed",
			email:       []string{"oke@smartsign.com", "failed@smartsign.com"},
			idsignature: nil,
			addr:        nil,
			repoTest: func(repo *m_repo.MockRepository) {
				email := []string{"oke@smartsign.com", "failed@smartsign.com"}
				for i := range email {
					repo.EXPECT().GetUserByEmail(email[i]).Return(models.User{}, errors.New("Failed Get Public Key"))
				}
			},
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			repo := m_repo.NewMockRepository(ctrl)
			tt.repoTest(repo)
			s := NewService(repo)
			addr, idsignature := s.GetPublicKey(tt.email)
			log.Println(addr)
			assert.Equal(t, tt.addr, addr)
			assert.Equal(t, tt.idsignature, idsignature)
		})
	}
}

func Test_service_GetCardDashboard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	test := []struct {
		name     string
		sign_id  string
		card     models.CardDashboard
		repoTest func(repo *m_repo.MockRepository)
	}{
		{
			name:    "Get Card Dashboard Service Case 1: Get Card Dashboard Success",
			sign_id: "rizwijaya",
			card: models.CardDashboard{
				TotalRequest:     9,
				TotalUser:        3,
				TotalTx:          5,
				TotalRequestUser: 12,
			},
			repoTest: func(repo *m_repo.MockRepository) {
				repo.EXPECT().GetTotal("signedDocuments").Return(9).Times(1)
				repo.EXPECT().GetTotal("users").Return(3).Times(1)
				repo.EXPECT().GetTotal("transactions").Return(5).Times(1)
				repo.EXPECT().GetTotalRequestUser("rizwijaya").Return(12).Times(1)
			},
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			repo := m_repo.NewMockRepository(ctrl)
			tt.repoTest(repo)
			s := NewService(repo)
			card := s.GetCardDashboard(tt.sign_id)
			assert.Equal(t, tt.card, card)
		})
	}
}

func Test_service_Logging(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()
	location, err := time.LoadLocation("Asia/Jakarta")
	assert.NoError(t, err)
	times := time.Now().In(location)
	test := []struct {
		name     string
		logg     models.UserLog
		wanterr  bool
		repoTest func(repo *m_repo.MockRepository)
	}{
		{
			name: "Logging Service Case 1: Write Log Login User Access Success",
			logg: models.UserLog{
				Idsignature:     "rizwijaya",
				User_agent:      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
				Ip_address:      "127.0.0.1",
				Action:          "Mengakses Halaman Login",
				Date_access:     times,
				Date_access_wib: "Minggu, 27 Nop 2022 | 18:30 WIB",
			},
			wanterr: false,
			repoTest: func(repo *m_repo.MockRepository) {
				logg := models.UserLog{
					Idsignature:     "rizwijaya",
					User_agent:      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
					Ip_address:      "127.0.0.1",
					Action:          "Mengakses Halaman Login",
					Date_access:     times,
					Date_access_wib: "Minggu, 27 Nop 2022 | 18:30 WIB",
				}
				repo.EXPECT().Logging(logg).Return(nil).Times(1)
			},
		},
		{
			name: "Logging Service Case 2: Write Log Dashboard Access Success",
			logg: models.UserLog{
				Idsignature:     "rizwijaya",
				User_agent:      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
				Ip_address:      "127.0.0.1",
				Action:          "Mengakses Halaman Dashboard",
				Date_access:     times,
				Date_access_wib: "Minggu, 27 Nop 2022 | 18:30 WIB",
			},
			wanterr: false,
			repoTest: func(repo *m_repo.MockRepository) {
				logg := models.UserLog{
					Idsignature:     "rizwijaya",
					User_agent:      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
					Ip_address:      "127.0.0.1",
					Action:          "Mengakses Halaman Dashboard",
					Date_access:     times,
					Date_access_wib: "Minggu, 27 Nop 2022 | 18:30 WIB",
				}
				repo.EXPECT().Logging(logg).Return(nil).Times(1)
			},
		},
		{
			name: "Logging Service Case 3: Write Log Failed",
			logg: models.UserLog{
				Idsignature:     "rizwijaya",
				User_agent:      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
				Ip_address:      "127.0.0.1",
				Action:          "Mengakses Halaman Dashboard",
				Date_access:     times,
				Date_access_wib: "Minggu, 27 Nop 2022 | 18:30 WIB",
			},
			wanterr: true,
			repoTest: func(repo *m_repo.MockRepository) {
				logg := models.UserLog{
					Idsignature:     "rizwijaya",
					User_agent:      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
					Ip_address:      "127.0.0.1",
					Action:          "Mengakses Halaman Dashboard",
					Date_access:     times,
					Date_access_wib: "Minggu, 27 Nop 2022 | 18:30 WIB",
				}
				repo.EXPECT().Logging(logg).Return(errors.New("Error Writing Log Failed")).Times(1)
			},
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			repo := m_repo.NewMockRepository(ctrl)
			tt.repoTest(repo)
			s := NewService(repo)
			err := s.Logging(tt.logg.Action, tt.logg.Idsignature, tt.logg.Ip_address, tt.logg.User_agent)
			if tt.wanterr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_service_GetLogUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()

	location, err := time.LoadLocation("Asia/Jakarta")
	assert.NoError(t, err)
	times := time.Now().In(location)
	id := primitive.NewObjectID()
	id2 := primitive.NewObjectID()

	test := []struct {
		name        string
		idsignature string
		logg        []models.UserLog
		wanterr     bool
		repoTest    func(repo *m_repo.MockRepository)
	}{
		{
			name:        "Get Log User Case 1: Get Log User Success",
			idsignature: "rizwijaya",
			logg: []models.UserLog{
				{
					Id:              id,
					Idsignature:     "rizwijaya",
					User_agent:      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
					Ip_address:      "127.0.0.1",
					Action:          "Mengakses Halaman Dashboard",
					Date_access:     times,
					Date_access_wib: "Minggu, 27 Nop 2022 | 18:30 WIB",
				},
				{
					Id:              id2,
					Idsignature:     "admin",
					User_agent:      "Mozilla/5.0 (Android 11; Mobile; rv:68.0) Gecko/68.0 Firefox/68.0",
					Ip_address:      "127.0.0.1",
					Action:          "Mengakses Halaman Login",
					Date_access:     times,
					Date_access_wib: "Minggu, 27 Nop 2022 | 18:30 WIB",
				},
			},
			wanterr: false,
			repoTest: func(repo *m_repo.MockRepository) {
				logg := []models.UserLog{
					{
						Id:              id,
						Idsignature:     "rizwijaya",
						User_agent:      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
						Ip_address:      "127.0.0.1",
						Action:          "Mengakses Halaman Dashboard",
						Date_access:     times,
						Date_access_wib: "Minggu, 27 Nop 2022 | 18:30 WIB",
					},
					{
						Id:              id2,
						Idsignature:     "admin",
						User_agent:      "Mozilla/5.0 (Android 11; Mobile; rv:68.0) Gecko/68.0 Firefox/68.0",
						Ip_address:      "127.0.0.1",
						Action:          "Mengakses Halaman Login",
						Date_access:     times,
						Date_access_wib: "Minggu, 27 Nop 2022 | 18:30 WIB",
					},
				}
				repo.EXPECT().GetLogUser("rizwijaya").Return(logg, nil).Times(1)
			},
		},
		{
			name:        "Get Log User Case 2: Error Get Log User Failed",
			idsignature: "rizwijaya",
			logg:        []models.UserLog{},
			wanterr:     true,
			repoTest: func(repo *m_repo.MockRepository) {
				repo.EXPECT().GetLogUser("rizwijaya").Return([]models.UserLog{}, errors.New("Failed to get user log")).Times(1)
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			repo := m_repo.NewMockRepository(ctrl)
			s := &service{
				repository: repo,
			}
			tt.repoTest(repo)
			userlog, err := s.GetLogUser(tt.idsignature)
			if tt.wanterr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, userlog, tt.logg)
			}
		})
	}
}
