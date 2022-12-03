package repository

import (
	"e-signature/modules/v1/utilities/user/models"
	m_blockchain "e-signature/pkg/blockchain/mock"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/tkuchiki/faketime"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestInit(t *testing.T) {
	gin.SetMode(gin.TestMode)
	os.Chdir("../../../../../")
}

func TestNewRepository(t *testing.T) {
	t.Run("New Repository Case 1: Success New Repository", func(t *testing.T) {
		repo := NewRepository(&mongo.Database{}, nil)
		assert.NotNil(t, repo)
	})
}

func Test_repository_GeneratePublicKey(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id := primitive.NewObjectID()

	test := []struct {
		name    string
		input   models.User
		output  models.User
		err     error
		testing func(db *mongo.Database, blockchain *m_blockchain.MockBlockchain)
	}{
		{
			name: "Generate Public Key Case 1: Success Generate Public Key",
			input: models.User{
				Id:       id,
				Name:     "Rizqi Wijaya",
				Password: "12345678910",
			},
			output: models.User{
				Id:        id,
				Name:      "Rizqi Wijaya",
				Password:  "12345678910",
				Publickey: "0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a",
			},
			err: nil,
			testing: func(db *mongo.Database, blockchain *m_blockchain.MockBlockchain) {
				blockchain.EXPECT().GeneratePublicKey(models.User{
					Id:       id,
					Name:     "Rizqi Wijaya",
					Password: "12345678910",
				}).Return(models.User{
					Id:        id,
					Name:      "Rizqi Wijaya",
					Password:  "12345678910",
					Publickey: "0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a",
				}, nil).Times(1)
			},
		},
		{
			name: "Generate Public Key Case 2: Error Failed Generate Public Key",
			input: models.User{
				Id:       id,
				Password: "1234567",
			},
			output: models.User{
				Id:       id,
				Password: "1234567",
			},
			err: errors.New("Error Generate Public Key"),
			testing: func(db *mongo.Database, blockchain *m_blockchain.MockBlockchain) {
				blockchain.EXPECT().GeneratePublicKey(models.User{
					Id:       id,
					Password: "1234567",
				}).Return(models.User{
					Id:       id,
					Password: "1234567",
				}, errors.New("Error Generate Public Key")).Times(1)
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			blockchain := m_blockchain.NewMockBlockchain(ctrl)
			db := &mongo.Database{}
			//collection := mtest.NewMockCollection(ctrl)
			repo := NewRepository(db, blockchain)
			tt.testing(db, blockchain)
			output, err := repo.GeneratePublicKey(tt.input)
			assert.Equal(t, tt.output, output)
			assert.Equal(t, tt.err, err)
		})
	}
}

func Test_repository_Register(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id := primitive.NewObjectID()

	var err error
	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()
	location, err := time.LoadLocation("Asia/Jakarta")
	assert.NoError(t, err)
	times := time.Now().In(location)

	test := []struct {
		nameTest string
		input    models.User
		response primitive.D
		err      error
	}{
		{
			nameTest: "Register Case 1: Success Register",
			input: models.User{
				Id:             id,
				Idsignature:    "admin",
				Name:           "Rizqi Wijaya",
				Email:          "smartsign@rizwijaya.com",
				Phone:          "081234567890",
				Identity_card:  "jskjdsa903ejwkldjlaskdsa",
				Password:       "DAJSLDA79DAS9UDWQJEJWOEKkeSDADA",
				Publickey:      "xjisud98di46o8hw9nei2yeiakjn3asm2dkd2121qwelkh67an8da",
				Role:           2,
				Dateregistered: times.String(),
			},
			response: mtest.CreateSuccessResponse(),
			err:      nil,
		},
		{
			nameTest: "Register Case 2: Failed Register Duplicate Key Error",
			input: models.User{
				Id:             id,
				Idsignature:    "admin",
				Name:           "Rizqi Wijaya",
				Email:          "smartsign@rizwijaya.com",
				Phone:          "081234567890",
				Identity_card:  "jskjdsa903ejwkldjlaskdsa",
				Password:       "DAJSLDA79DAS9UDWQJEJWOEKkeSDADA",
				Publickey:      "xjisud98di46o8hw9nei2yeiakjn3asm2dkd2121qwelkh67an8da",
				Role:           2,
				Dateregistered: times.String(),
			},
			response: mtest.CreateWriteErrorsResponse(mtest.WriteError{
				Index:   1,
				Code:    11000,
				Message: "Duplicate Key Error",
			}),
			err: errors.New("Duplicate Key Error"),
		},
	}
	for _, tt := range test {
		mt.Run(tt.nameTest, func(mt *mtest.T) {
			mt.AddMockResponses(tt.response)
			blockchain := m_blockchain.NewMockBlockchain(ctrl)
			repo := NewRepository(mt.DB, blockchain)
			_, err := repo.Register(tt.input)
			if err != nil {
				if tt.err.Error() == "Duplicate Key Error" {
					assert.True(t, mongo.IsDuplicateKeyError(err))
				}
			} else {
				assert.Equal(t, tt.err, err)
			}
		})
	}
}

func Test_repository_CheckUserExist(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id := primitive.NewObjectID()

	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()
	times := time.Now()

	test := []struct {
		nameTest    string
		idsignature string
		output      models.ProfileDB
		response    primitive.D
		err         error
	}{
		{
			nameTest:    "Check User Exist Case 1: Success User Exist",
			idsignature: "admin",
			output: models.ProfileDB{
				Id:            id,
				Idsignature:   "admin",
				Name:          "Rizqi Wijaya",
				Email:         "smartsign@rizwijaya.com",
				Phone:         "081234567890",
				Identity_card: "jskjdsa903ejwkldjlaskdsa",
				Password:      "DAJSLDA79DAS9UDWQJEJWOEKkeSDADA",
				PublicKey:     "dsaknkdlak8ywq8jlasm4342wasdas234214wdsa",
				Role_id:       2,
				Date_created:  times,
			},
			response: mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, primitive.D{
				{Key: "_id", Value: id},
				{Key: "idsignature", Value: "admin"},
				{Key: "name", Value: "Rizqi Wijaya"},
				{Key: "email", Value: "smartsign@rizwijaya.com"},
				{Key: "phone", Value: "081234567890"},
				{Key: "identity_card", Value: "jskjdsa903ejwkldjlaskdsa"},
				{Key: "password", Value: "DAJSLDA79DAS9UDWQJEJWOEKkeSDADA"},
				{Key: "public_key", Value: "dsaknkdlak8ywq8jlasm4342wasdas234214wdsa"},
				{Key: "role", Value: 2},
				{Key: "date_created", Value: times},
			}),
			err: nil,
		},
		{
			nameTest:    "Check User Exist Case 2: Success User Not Exist",
			idsignature: "admin23",
			output:      models.ProfileDB{},
			response:    mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, primitive.D{}),
			err:         nil,
		},
		{
			nameTest:    "Check User Exist Case 3: Error Failed Decoded Data User",
			idsignature: "adminbrow",
			output:      models.ProfileDB{},
			response: mtest.CreateCommandErrorResponse(mtest.CommandError{
				Code:    1,
				Message: "Failed decoded",
			}),
			err: errors.New("Failed decoded"),
		},
	}
	for _, tt := range test {
		mt.Run(tt.nameTest, func(mt *mtest.T) {
			mt.AddMockResponses(tt.response)
			blockchain := m_blockchain.NewMockBlockchain(ctrl)
			repo := NewRepository(mt.DB, blockchain)
			user, err := repo.CheckUserExist(tt.idsignature)
			if err != nil {
				assert.Equal(t, err.Error(), tt.err.Error())
			} else {
				assert.Equal(t, err, tt.err)
			}
			assert.Equal(t, tt.output, user)
		})
	}
}

func Test_repository_CheckEmailExist(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id := primitive.NewObjectID()

	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()
	times := time.Now()

	test := []struct {
		nameTest string
		email    string
		output   models.ProfileDB
		response primitive.D
		err      error
	}{
		{
			nameTest: "Check Email Exist Case 1: Success User Exist",
			email:    "smartsign@rizwijaya.com",
			output: models.ProfileDB{
				Id:            id,
				Idsignature:   "admin",
				Name:          "Rizqi Wijaya",
				Email:         "smartsign@rizwijaya.com",
				Phone:         "081234567890",
				Identity_card: "jskjdsa903ejwkldjlaskdsa",
				Password:      "DAJSLDA79DAS9UDWQJEJWOEKkeSDADA",
				PublicKey:     "dsaknkdlak8ywq8jlasm4342wasdas234214wdsa",
				Role_id:       2,
				Date_created:  times,
			},
			response: mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, primitive.D{
				{Key: "_id", Value: id},
				{Key: "idsignature", Value: "admin"},
				{Key: "name", Value: "Rizqi Wijaya"},
				{Key: "email", Value: "smartsign@rizwijaya.com"},
				{Key: "phone", Value: "081234567890"},
				{Key: "identity_card", Value: "jskjdsa903ejwkldjlaskdsa"},
				{Key: "password", Value: "DAJSLDA79DAS9UDWQJEJWOEKkeSDADA"},
				{Key: "public_key", Value: "dsaknkdlak8ywq8jlasm4342wasdas234214wdsa"},
				{Key: "role", Value: 2},
				{Key: "date_created", Value: times},
			}),
			err: nil,
		},
		{
			nameTest: "Check Email Exist Case 2: Success User Not Exist",
			email:    "admin@smart.com",
			output:   models.ProfileDB{},
			response: mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, primitive.D{}),
			err:      nil,
		},
		{
			nameTest: "Check Email Exist Case 3: Error Failed Decoded Data User",
			email:    "sembada@smartsign.com",
			output:   models.ProfileDB{},
			response: mtest.CreateCommandErrorResponse(mtest.CommandError{
				Code:    1,
				Message: "Failed decoded",
			}),
			err: errors.New("Failed decoded"),
		},
	}
	for _, tt := range test {
		mt.Run(tt.nameTest, func(mt *mtest.T) {
			mt.AddMockResponses(tt.response)
			blockchain := m_blockchain.NewMockBlockchain(ctrl)
			repo := NewRepository(mt.DB, blockchain)
			user, err := repo.CheckEmailExist(tt.email)
			if err != nil {
				assert.Equal(t, err.Error(), tt.err.Error())
			} else {
				assert.Equal(t, err, tt.err)
			}
			assert.Equal(t, tt.output, user)
		})
	}
}

func Test_repository_GetUserByEmail(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id := primitive.NewObjectID()

	test := []struct {
		nameTest string
		email    string
		output   models.User
		response primitive.D
		err      error
	}{
		{
			nameTest: "Get User By Email Case 1: Success User Exist",
			email:    "smartsign@rizwijaya.com",
			output: models.User{
				Id:             id,
				Idsignature:    "admin",
				Name:           "Rizqi Wijaya",
				Password:       "DAJSLDA79DAS9UDWQJEJWOEKkeSDADA",
				PasswordHash:   "",
				Role:           2,
				Publickey:      "dsaknkdlak8ywq8jlasm4342wasdas234214wdsa",
				Identity_card:  "jskjdsa903ejwkldjlaskdsa",
				Email:          "smartsign@rizwijaya.com",
				Phone:          "081234567890",
				Dateregistered: "",
			},
			response: mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, primitive.D{
				{Key: "_id", Value: id},
				{Key: "idsignature", Value: "admin"},
				{Key: "name", Value: "Rizqi Wijaya"},
				{Key: "email", Value: "smartsign@rizwijaya.com"},
				{Key: "phone", Value: "081234567890"},
				{Key: "identity_card", Value: "jskjdsa903ejwkldjlaskdsa"},
				{Key: "password", Value: "DAJSLDA79DAS9UDWQJEJWOEKkeSDADA"},
				{Key: "public_key", Value: "dsaknkdlak8ywq8jlasm4342wasdas234214wdsa"},
				{Key: "role", Value: 2},
			}),
			err: nil,
		},
		{
			nameTest: "Get User By Email Case 2: Success User Not Exist",
			email:    "admin@smart.com",
			output:   models.User{},
			response: mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, primitive.D{}),
			err:      nil,
		},
		{
			nameTest: "Get User By Email Case 3: Error Failed Decoded Data User",
			email:    "sembada@smartsign.com",
			output:   models.User{},
			response: mtest.CreateCommandErrorResponse(mtest.CommandError{
				Code:    1,
				Message: "Failed decoded",
			}),
			err: errors.New("Failed decoded"),
		},
	}
	for _, tt := range test {
		mt.Run(tt.nameTest, func(mt *mtest.T) {
			mt.AddMockResponses(tt.response)
			blockchain := m_blockchain.NewMockBlockchain(ctrl)
			repo := NewRepository(mt.DB, blockchain)
			user, err := repo.GetUserByEmail(tt.email)
			if err != nil {
				assert.Equal(t, err.Error(), tt.err.Error())
			} else {
				assert.Equal(t, err, tt.err)
			}
			assert.Equal(t, tt.output, user)
		})
	}
}

func Test_repository_GetTotal(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	//id := primitive.NewObjectID()

	test := []struct {
		nameTest string
		db       string
		output   int
		response primitive.D
		err      error
	}{
		{
			nameTest: "Get Total Data Case 1: Success Get Total SignedDocuments",
			db:       "signedDocuments",
			output:   430,
			response: mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, primitive.D{
				{Key: "n", Value: 430},
			}),
			err: nil,
		},
		{
			nameTest: "Get Total Data Case 2: Success Get Total Users",
			db:       "users",
			output:   123,
			response: mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, primitive.D{
				{Key: "n", Value: 123},
			}),
			err: nil,
		},
		{
			nameTest: "Get Total Data Case 3: Success Get Total Transactions",
			db:       "transactions",
			output:   743,
			response: mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, primitive.D{
				{Key: "n", Value: 743},
			}),
			err: nil,
		},
		{
			nameTest: "Get Total Data Case 4: Error Failed Get Total Data",
			db:       "signss",
			output:   0,
			response: mtest.CreateCommandErrorResponse(mtest.CommandError{
				Code:    1,
				Message: "Failed get total data",
			}),
		},
	}
	for _, tt := range test {
		mt.Run(tt.nameTest, func(mt *mtest.T) {
			mt.AddMockResponses(tt.response)
			blockchain := m_blockchain.NewMockBlockchain(ctrl)
			repo := NewRepository(mt.DB, blockchain)
			total := repo.GetTotal(tt.db)
			assert.Equal(t, tt.output, total)
		})
	}
}

func Test_repository_GetTotalRequestUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	//id := primitive.NewObjectID()

	test := []struct {
		nameTest string
		sign_id  string
		output   int
		response primitive.D
		err      error
	}{
		{
			nameTest: "Get Total Request User Case 1: Success Get Total Request User",
			sign_id:  "riwijaya",
			output:   430,
			response: mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, primitive.D{
				{Key: "n", Value: 430},
			}),
			err: nil,
		},
		{
			nameTest: "Get Total Request User Case 2: Error Failed Get Total Request User",
			sign_id:  "adminss",
			output:   0,
			response: mtest.CreateCommandErrorResponse(mtest.CommandError{
				Code:    1,
				Message: "Failed get total data",
			}),
		},
	}
	for _, tt := range test {
		mt.Run(tt.nameTest, func(mt *mtest.T) {
			mt.AddMockResponses(tt.response)
			blockchain := m_blockchain.NewMockBlockchain(ctrl)
			repo := NewRepository(mt.DB, blockchain)
			total := repo.GetTotalRequestUser(tt.sign_id)
			assert.Equal(t, tt.output, total)
		})
	}
}

func Test_repository_Logging(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id := primitive.NewObjectID()

	var err error
	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()
	location, err := time.LoadLocation("Asia/Jakarta")
	assert.NoError(t, err)
	times := time.Now().In(location)

	test := []struct {
		nameTest string
		input    models.UserLog
		response primitive.D
		err      error
	}{
		{
			nameTest: "Logging Case 1: Success Insert User Log",
			input: models.UserLog{
				Id:          id,
				Idsignature: "admin",
				User_agent:  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36",
				Ip_address:  "127.0.0.1",
				Action:      "Mengakses halaman dashboard",
				Date_access: times,
			},
			response: mtest.CreateSuccessResponse(),
			err:      nil,
		},
		{
			nameTest: "Logging Case 2: Error Failed Insert User Log",
			input: models.UserLog{
				Id:          id,
				Idsignature: "admin",
				User_agent:  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36",
				Ip_address:  "127.0.0.1",
				Action:      "Mengakses halaman dashboard",
				Date_access: times,
			},
			response: mtest.CreateWriteErrorsResponse(mtest.WriteError{
				Index:   1,
				Code:    11000,
				Message: "Duplicate Key Error",
			}),
			err: errors.New("Duplicate Key Error"),
		},
	}
	for _, tt := range test {
		mt.Run(tt.nameTest, func(mt *mtest.T) {
			mt.AddMockResponses(tt.response)
			blockchain := m_blockchain.NewMockBlockchain(ctrl)
			repo := NewRepository(mt.DB, blockchain)
			err := repo.Logging(tt.input)
			if err != nil {
				if tt.err.Error() == "Duplicate Key Error" {
					assert.True(t, mongo.IsDuplicateKeyError(err))
				}
			} else {
				assert.Equal(t, tt.err, err)
			}
		})
	}
}

func Test_repository_GetLogUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id := primitive.NewObjectID()
	id2 := primitive.NewObjectID()
	id3 := primitive.NewObjectID()
	id4 := primitive.NewObjectID()

	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()
	times := time.Now()

	test := []struct {
		nameTest    string
		idsignature string
		output      []models.UserLog
		response1   primitive.D
		response2   primitive.D
		response3   primitive.D
		response4   primitive.D
		err         error
	}{
		{
			nameTest:    "Get Log User Case 1: Success Get Log User",
			idsignature: "admin",
			output: []models.UserLog{
				{
					Id:          id,
					Idsignature: "admin",
					User_agent:  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36",
					Ip_address:  "127.0.0.1",
					Action:      "Mengakses halaman dashboard",
					Date_access: times,
				},
				{
					Id:          id2,
					Idsignature: "admin",
					User_agent:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1)",
					Ip_address:  "127.0.0.1",
					Action:      "Mengakses halaman my signature",
					Date_access: times,
				},
				{
					Id:          id3,
					Idsignature: "admin",
					User_agent:  "Chrome/95.0.4638.69 Safari/537.36",
					Ip_address:  "127.0.0.1",
					Action:      "Mengakses halaman tanda tangan sekarang",
					Date_access: times,
				},
				{
					Id:          id4,
					Idsignature: "admin",
					User_agent:  "Mozilla/6.0 (Android 10; Mobile; rv:68.0) Gecko/68.0 Firefox/68.0",
					Ip_address:  "127.0.0.1",
					Action:      "Mengakses halaman riwayat tanda tangan",
					Date_access: times,
				},
			},
			response1: mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, primitive.D{
				{Key: "_id", Value: id},
				{Key: "idsignature", Value: "admin"},
				{Key: "user_agent", Value: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36"},
				{Key: "ip_address", Value: "127.0.0.1"},
				{Key: "action", Value: "Mengakses halaman dashboard"},
				{Key: "date_accessed", Value: times},
			}),
			response2: mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, primitive.D{
				{Key: "_id", Value: id2},
				{Key: "idsignature", Value: "admin"},
				{Key: "user_agent", Value: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1)"},
				{Key: "ip_address", Value: "127.0.0.1"},
				{Key: "action", Value: "Mengakses halaman my signature"},
				{Key: "date_accessed", Value: times},
			}),
			response3: mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, primitive.D{
				{Key: "_id", Value: id3},
				{Key: "idsignature", Value: "admin"},
				{Key: "user_agent", Value: "Chrome/95.0.4638.69 Safari/537.36"},
				{Key: "ip_address", Value: "127.0.0.1"},
				{Key: "action", Value: "Mengakses halaman tanda tangan sekarang"},
				{Key: "date_accessed", Value: times},
			}),
			response4: mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch, primitive.D{
				{Key: "_id", Value: id4},
				{Key: "idsignature", Value: "admin"},
				{Key: "user_agent", Value: "Mozilla/6.0 (Android 10; Mobile; rv:68.0) Gecko/68.0 Firefox/68.0"},
				{Key: "ip_address", Value: "127.0.0.1"},
				{Key: "action", Value: "Mengakses halaman riwayat tanda tangan"},
				{Key: "date_accessed", Value: times},
			}),
			err: nil,
		},
		{
			nameTest:    "Get Log User Case 2: Error Failed Get Log User",
			idsignature: "admin",
			output:      []models.UserLog(nil),
			response1: mtest.CreateCommandErrorResponse(mtest.CommandError{
				Code:    50,
				Message: "Database Timeout",
			}),
			err: errors.New("Database Timeout"),
		},
		{
			nameTest:    "Get Log User Case 3: No Data Log User",
			idsignature: "admin",
			output: []models.UserLog{
				{
					Id: primitive.ObjectID{},
				},
			},
			response1: mtest.CreateCursorResponse(0, "foo.bar", mtest.FirstBatch, primitive.D{}),
			err:       nil,
		},
	}
	for _, tt := range test {
		mt.Run(tt.nameTest, func(mt *mtest.T) {
			mt.AddMockResponses(tt.response1, tt.response2, tt.response3, tt.response4)

			blockchain := m_blockchain.NewMockBlockchain(ctrl)
			repo := NewRepository(mt.DB, blockchain)
			log, err := repo.GetLogUser(tt.idsignature)
			if err != nil {
				assert.Equal(t, tt.err.Error(), err.Error())
			} else {
				assert.Equal(t, err, tt.err)
			}
			assert.Equal(t, tt.output, log)
		})
	}
}
