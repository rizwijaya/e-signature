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

	//var err error
	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()
	// location, err := time.LoadLocation("Asia/Jakarta")
	// assert.NoError(t, err)
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
			idsignature: "admin",
			output:      models.ProfileDB{},
			response:    mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, primitive.D{}),
			err:         nil,
		},
		{
			nameTest:    "Check User Exist Case 3: Error Failed Decoded Data User",
			idsignature: "admin",
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
