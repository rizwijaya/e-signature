package repository

import (
	"e-signature/modules/v1/utilities/user/models"
	m_blockchain "e-signature/pkg/blockchain/mock"
	"errors"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
			repo := NewRepository(db, blockchain)
			tt.testing(db, blockchain)
			output, err := repo.GeneratePublicKey(tt.input)
			assert.Equal(t, tt.output, output)
			assert.Equal(t, tt.err, err)
		})
	}
}

// func Test_repository_Register(t *testing.T) {
// 	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
// 	defer mt.Close()

// 	mt.Run("Register Case 1: Success Register", func(mt *mtest.T) {
// 		// mt.AddMockResponses(bson.D{{"ok", 1.0}, {"n", 1}})

// 		// repo := NewRepository(&mongo.Database{}, nil)
// 		// user := models.User{}
// 		// _, err := repo.Register(user)
// 		// assert.Nil(t, err)
// 		userCollection := mt.Coll
// 		userCollection.InsertOne(nil, bson.M{
// 			"name": "test",
// 		})
// 	})

// }
