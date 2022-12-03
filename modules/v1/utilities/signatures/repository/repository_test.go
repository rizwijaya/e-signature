package repository

import (
	modelsUser "e-signature/modules/v1/utilities/user/models"
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

func Test_repository_LogTransactions(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type input struct {
		address string
		tx_hash string
		nonce   string
		desc    string
		prices  string
	}
	test := []struct {
		nameTest string
		input    input
		response primitive.D
		err      error
	}{
		{
			nameTest: "Log Transations Case 1: Success Insert Log Transactions",
			input: input{
				address: "0xDBf0b1bBb2b7b2bBb2bBb2bBb2bBb2bBb2bBb2bBb",
				tx_hash: "0x5f44e265dbf57984ffb9a833ba9cde9c51a6bec419c44f8e40b64a9ee7033c83",
				nonce:   "2",
				desc:    "Membuat Dokumen sample.pdf untuk tanda tangan",
				prices:  "300000",
			},
			response: mtest.CreateSuccessResponse(),
			err:      nil,
		},
		{
			nameTest: "Log Transations Case 2: Error Failed Insert Log Transactions Duplicate Key",
			input: input{
				address: "0xDBf0b1bBb2b7b2bBb2bBb2bBb2bBb2bBb2bBb2bBb",
				tx_hash: "0x5f44e265dbf57984ffb9a833ba9cde9c51a6bec419c44f8e40b64a9ee7033c83",
				nonce:   "2",
				desc:    "Membuat Dokumen sample.pdf untuk tanda tangan",
				prices:  "300000",
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
			err := repo.LogTransactions(tt.input.address, tt.input.tx_hash, tt.input.nonce, tt.input.desc, tt.input.prices)
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

func Test_repository_DefaultSignatures(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()

	test := []struct {
		nameTest string
		input    modelsUser.User
		id       string
		response primitive.D
		err      error
	}{
		{
			nameTest: "Default Signatures Case 1: Success Insert Default Signatures",
			input: modelsUser.User{
				Idsignature: "rizwijaya",
			},
			id:       "5fcbf9c1b2b7b2bBb2bBb2bBb2bBb2bBb2bBb2bBb",
			response: mtest.CreateSuccessResponse(),
			err:      nil,
		},
		{
			nameTest: "Default Signatures Case 2: Error Failed Insert Default Signatures Duplicate Key",
			input: modelsUser.User{
				Idsignature: "admin",
			},
			id: "f9c1b2b72bBbbBb2bBb225bfcbbBb2bBb2bBb2bBb",
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
			err := repo.DefaultSignatures(tt.input, tt.id)
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
