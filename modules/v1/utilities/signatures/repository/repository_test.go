package repository

import (
	"e-signature/modules/v1/utilities/signatures/models"
	modelsUser "e-signature/modules/v1/utilities/user/models"
	m_blockchain "e-signature/pkg/blockchain/mock"
	"errors"
	"log"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

func Test_repository_UpdateMySignatures(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()

	test := []struct {
		nameTest      string
		signature     string
		signaturedata string
		sign          string
		response      primitive.D
		err           error
	}{
		{
			nameTest:      "Update My Signatures Case 1: Success Update My Signatures",
			signature:     "default.png",
			signaturedata: "default.png",
			sign:          "rizwijaya",
			response:      mtest.CreateSuccessResponse(),
			err:           nil,
		},
		{
			nameTest:      "Update My Signatures Case 2: Error Failed Update My Signatures Data Not Found",
			signature:     "default.png",
			signaturedata: "default.png",
			sign:          "rizwijaya",
			response: mtest.CreateWriteErrorsResponse(
				mtest.WriteError{
					Index:   1,
					Code:    0,
					Message: "Data Not Found",
				},
			),
			err: errors.New("Data Not Found"),
		},
	}
	for _, tt := range test {
		mt.Run(tt.nameTest, func(mt *mtest.T) {
			mt.AddMockResponses(tt.response)
			blockchain := m_blockchain.NewMockBlockchain(ctrl)
			repo := NewRepository(mt.DB, blockchain)
			err := repo.UpdateMySignatures(tt.signature, tt.signaturedata, tt.sign)
			if err != nil {
				assert.Equal(t, "write exception: write errors: ["+tt.err.Error()+"]", err.Error())
			} else {
				assert.Equal(t, tt.err, err)
			}
		})
	}
}

func Test_repository_GetMySignature(t *testing.T) {
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
		sign     string
		output   models.Signatures
		response primitive.D
		err      error
	}{
		{
			nameTest: "Get My Signature Case 1: Success Get My Signature Data",
			sign:     "admin",
			output: models.Signatures{
				Id:                 id,
				User:               "admin",
				Signature:          "default.png",
				Signature_data:     "default.png",
				Latin:              "latin-signature.png",
				Latin_data:         "latindata-signature.png",
				Signature_selected: "signature",
				Date_update:        times,
				Date_created:       times,
			},
			response: mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, primitive.D{
				{Key: "_id", Value: id},
				{Key: "user", Value: "admin"},
				{Key: "signature", Value: "default.png"},
				{Key: "signature_data", Value: "default.png"},
				{Key: "latin", Value: "latin-signature.png"},
				{Key: "latin_data", Value: "latindata-signature.png"},
				{Key: "signature_selected", Value: "signature"},
				{Key: "date_update", Value: times},
				{Key: "date_created", Value: times},
			}),
			err: nil,
		},
		{
			nameTest: "Get My Signature Case 2: Signature Data Not Found",
			sign:     "admin23",
			output:   models.Signatures{},
			response: mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, primitive.D{}),
			err:      nil,
		},
		{
			nameTest: "Get My Signature Case 3: Error Failed Decode Data My Signature",
			sign:     "adminbrow",
			output:   models.Signatures{},
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
			signatures, err := repo.GetMySignature(tt.sign)
			if err != nil {
				assert.Equal(t, err.Error(), tt.err.Error())
			} else {
				assert.Equal(t, err, tt.err)
			}
			assert.Equal(t, tt.output, signatures)
		})
	}
}

func Test_repository_ChangeSignatures(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()

	test := []struct {
		nameTest  string
		sign_type string
		sign      string
		response  primitive.D
		err       error
	}{
		{
			nameTest:  "Change Signatures Case 1: Success Change Signatures to Signature",
			sign_type: "signature",
			sign:      "rizwijaya",
			response:  mtest.CreateSuccessResponse(),
			err:       nil,
		},
		{
			nameTest:  "Change Signatures Case 2: Success Change Signatures to Signature Data",
			sign_type: "signature_data",
			sign:      "rizwijaya",
			response:  mtest.CreateSuccessResponse(),
			err:       nil,
		},
		{
			nameTest:  "Change Signatures Case 4: Success Change Signatures to Latin",
			sign_type: "latin",
			sign:      "rizwijaya",
			response:  mtest.CreateSuccessResponse(),
			err:       nil,
		},
		{
			nameTest:  "Change Signatures Case 5: Success Change Signatures to Latin Data",
			sign_type: "latin_data",
			sign:      "rizwijaya",
			response:  mtest.CreateSuccessResponse(),
			err:       nil,
		},
		{
			nameTest:  "Change Signatures Case 5: Error Failed Change Signatures Data Not Found",
			sign_type: "signatursse",
			sign:      "admin12",
			response: mtest.CreateWriteErrorsResponse(
				mtest.WriteError{
					Index:   1,
					Code:    0,
					Message: "Data Not Found",
				},
			),
			err: errors.New("Data Not Found"),
		},
	}
	for _, tt := range test {
		mt.Run(tt.nameTest, func(mt *mtest.T) {
			mt.AddMockResponses(tt.response)
			blockchain := m_blockchain.NewMockBlockchain(ctrl)
			repo := NewRepository(mt.DB, blockchain)
			err := repo.ChangeSignature(tt.sign_type, tt.sign)
			if err != nil {
				assert.Equal(t, "write exception: write errors: ["+tt.err.Error()+"]", err.Error())
			} else {
				assert.Equal(t, tt.err, err)
			}
		})
	}
}

func Test_repository_AddToBlockhain(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()
	times := time.Now()
	timeSign := new(big.Int)
	timeFormat := times.Format("15040502012006")
	timeSign, _ = timeSign.SetString(timeFormat, 10)

	test := []struct {
		nameTest string
		input    models.SignDocuments
		times    *big.Int
		testing  func(blockchain *m_blockchain.MockBlockchain, repo Repository)
		err      error
	}{
		{
			nameTest: "Add To Blockchain Case 1: Success Insert Data To Blockchain",
			input: models.SignDocuments{
				Name:       "sample_test.pdf",
				SignPage:   1.0,
				X_coord:    1.3,
				Y_coord:    1.2,
				Height:     4.2,
				Width:      5.3,
				Invite_sts: true,
				Email:      []string{"admin@smartsign.com"},
				Creator:    "0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a",
				Creator_id: "rizwijaya",
			},
			times: timeSign,
			testing: func(blockchain *m_blockchain.MockBlockchain, repo Repository) {
				input := models.SignDocuments{
					Name:       "sample_test.pdf",
					SignPage:   1.0,
					X_coord:    1.3,
					Y_coord:    1.2,
					Height:     4.2,
					Width:      5.3,
					Invite_sts: true,
					Email:      []string{"admin@smartsign.com"},
					Creator:    "0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a",
					Creator_id: "rizwijaya",
				}
				testAddr := common.HexToAddress("b94f5374fce5edbc8e2a8697c15331677e6ebf0b")
				tx := types.NewTx(&types.AccessListTx{
					ChainID:  big.NewInt(1),
					Nonce:    3,
					To:       &testAddr,
					Value:    big.NewInt(10),
					Gas:      25000,
					GasPrice: big.NewInt(1),
					Data:     common.FromHex("0xc4d57bb9b9f95452da96c5d2ca8e3e477672c01afea8362348406a4236fb942f"),
				})
				auth := &bind.TransactOpts{
					From:     common.HexToAddress("0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a"),
					Signer:   nil,
					GasLimit: 0,
					Value:    big.NewInt(0),
					Nonce:    big.NewInt(1),
				}
				blockchain.EXPECT().AddToBlockhain(input, timeSign).Return(tx, auth, nil).Times(1)
			},
			err: nil,
		},
		{
			nameTest: "Add To Blockchain Case 2: Error Failed Insert Data To Blockchain",
			input: models.SignDocuments{
				Name:       "testing.pdf",
				SignPage:   1.0,
				X_coord:    24.3,
				Y_coord:    369.2,
				Height:     400.2,
				Width:      500.3,
				Invite_sts: true,
				Email:      []string{"admin@smartsign.com"},
				Creator:    "0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a",
				Creator_id: "admin",
			},
			times: timeSign,
			testing: func(blockchain *m_blockchain.MockBlockchain, repo Repository) {
				input := models.SignDocuments{
					Name:       "testing.pdf",
					SignPage:   1.0,
					X_coord:    24.3,
					Y_coord:    369.2,
					Height:     400.2,
					Width:      500.3,
					Invite_sts: true,
					Email:      []string{"admin@smartsign.com"},
					Creator:    "0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a",
					Creator_id: "admin",
				}
				testAddr := common.HexToAddress("b94f5374fce5edbc8e2a8697c15331677e6ebf0b")
				tx := types.NewTx(&types.AccessListTx{
					ChainID:  big.NewInt(1),
					Nonce:    4,
					To:       &testAddr,
					Value:    big.NewInt(25),
					Gas:      20,
					GasPrice: big.NewInt(30000),
					Data:     common.FromHex("0xc4d57bb9b9f95452da96c5d2ca8e3e477672c01afea8362348406a4236fb942f"),
				})
				auth := &bind.TransactOpts{
					From:     common.HexToAddress("0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a"),
					Signer:   nil,
					GasLimit: 20,
					Value:    big.NewInt(0),
					Nonce:    big.NewInt(1),
				}
				blockchain.EXPECT().AddToBlockhain(input, timeSign).Return(tx, auth, types.ErrGasFeeCapTooLow).Times(1)
			},
			err: types.ErrGasFeeCapTooLow,
		},
	}
	for _, tt := range test {
		mt.Run(tt.nameTest, func(mt *mtest.T) {
			blockchain := m_blockchain.NewMockBlockchain(ctrl)
			repo := NewRepository(mt.DB, blockchain)
			if tt.testing != nil {
				tt.testing(blockchain, repo)
			}

			err := repo.AddToBlockhain(tt.input, tt.times)
			assert.Equal(t, err, tt.err)
		})
	}
}

func Test_repository_AddUserDocs(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()

	test := []struct {
		nameTest string
		input    models.SignDocuments
		response primitive.D
		err      error
	}{
		{
			nameTest: "Add User Document Case 1: Success Insert Data Document and User",
			input: models.SignDocuments{
				Name:          "test_sample.pdf",
				Hash_original: "doasdjihqwoyrno3y981n234mbxjlwnalksdmlasdp93uejrk3e213",
				Hash:          "dsjaldajsdoasjdopwud09aud98peu21j3newndaksdnlaskjdasodjkasd",
				Judul:         "",
				Note:          "",
				Address:       []common.Address{common.HexToAddress("0xAyysae6513c99443cF32Ca8A449f5287aaD6f91a")},
			},
			response: mtest.CreateSuccessResponse(),
			err:      nil,
		},
		{
			nameTest: "Add User Document Case 2: Error Failed Insert Data Document and User Duplicate Key",
			input: models.SignDocuments{
				Name:          "test.pdf",
				Hash_original: "rnn234mbxjldjihqwoywndoasalksdo3y981mlasdp93uejrk3e213",
				Hash:          "eusdnlaskjdasodjkasddsj21j3newndakaldajsdoasjdopwud09aud98p",
				Judul:         "",
				Note:          "",
				Address:       []common.Address{common.HexToAddress("0xF32CaA449f52ae6513c8Ayys99443c87aaD6f91a"), common.HexToAddress("0xBha62e9443cF32Ca86f916513c9aA449f5287aaD")},
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
			err := repo.AddUserDocs(tt.input)
			if err != nil {
				log.Println(err)
				if tt.err.Error() == "Duplicate Key Error" {
					assert.True(t, mongo.IsDuplicateKeyError(err))
				}
			} else {
				assert.Equal(t, tt.err, err)
			}
		})
	}
}

func Test_repository_DocumentSigned(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()
	times := time.Now()
	timeSign := new(big.Int)
	timeFormat := times.Format("15040502012006")
	timeSign, _ = timeSign.SetString(timeFormat, 10)

	test := []struct {
		nameTest string
		input    models.SignDocs
		times    *big.Int
		testing  func(blockchain *m_blockchain.MockBlockchain, repo Repository)
		err      error
	}{
		{
			nameTest: "Document Signed Case 1: Success Signed Document in Blockchain",
			input: models.SignDocs{
				Hash_original: "doasdjihqwoyrno3y981n234mbxjlwnalksdmlasdp93uejrk3e213",
				Creator:       "rizwijaya",
				Hash:          "dsjaldajsdoasjdopwud09aud98peu21j3newndaksdnlaskjdasodjkasd",
				IPFS:          "aldasodjkas8peu21jjsdoasjdopwud09aud9aj3newndaksdnlaskjdd",
			},
			times: timeSign,
			testing: func(blockchain *m_blockchain.MockBlockchain, repo Repository) {
				input := models.SignDocs{
					Hash_original: "doasdjihqwoyrno3y981n234mbxjlwnalksdmlasdp93uejrk3e213",
					Creator:       "rizwijaya",
					Hash:          "dsjaldajsdoasjdopwud09aud98peu21j3newndaksdnlaskjdasodjkasd",
					IPFS:          "aldasodjkas8peu21jjsdoasjdopwud09aud9aj3newndaksdnlaskjdd",
				}
				testAddr := common.HexToAddress("b94f5374fce5edbc8e2a8697c15331677e6ebf0b")
				tx := types.NewTx(&types.AccessListTx{
					ChainID:  big.NewInt(1),
					Nonce:    3,
					To:       &testAddr,
					Value:    big.NewInt(10),
					Gas:      25000,
					GasPrice: big.NewInt(1),
					Data:     common.FromHex("0xc4d57bb9b9f95452da96c5d2ca8e3e477672c01afea8362348406a4236fb942f"),
				})
				auth := &bind.TransactOpts{
					From:     common.HexToAddress("0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a"),
					Signer:   nil,
					GasLimit: 0,
					Value:    big.NewInt(0),
					Nonce:    big.NewInt(1),
				}
				blockchain.EXPECT().DocumentSigned(input, timeSign).Return(tx, auth, nil).Times(1)
			},
			err: nil,
		},
		{
			nameTest: "Document Signed Case 2: Error Failed Signed Document in Blockchain",
			input: models.SignDocs{
				Hash_original: "doasdjihqwoyrno3y981n234mbxjlwnalksdmlasdp93uejrk3e213",
				Creator:       "rizwijaya",
				Hash:          "dsjaldajsdoasjdopwud09aud98peu21j3newndaksdnlaskjdasodjkasd",
				IPFS:          "aldasodjkas8peu21jjsdoasjdopwud09aud9aj3newndaksdnlaskjdd",
			},
			times: timeSign,
			testing: func(blockchain *m_blockchain.MockBlockchain, repo Repository) {
				input := models.SignDocs{
					Hash_original: "doasdjihqwoyrno3y981n234mbxjlwnalksdmlasdp93uejrk3e213",
					Creator:       "rizwijaya",
					Hash:          "dsjaldajsdoasjdopwud09aud98peu21j3newndaksdnlaskjdasodjkasd",
					IPFS:          "aldasodjkas8peu21jjsdoasjdopwud09aud9aj3newndaksdnlaskjdd",
				}
				testAddr := common.HexToAddress("b94f5374fce5edbc8e2a8697c15331677e6ebf0b")
				tx := types.NewTx(&types.AccessListTx{
					ChainID:  big.NewInt(1),
					Nonce:    4,
					To:       &testAddr,
					Value:    big.NewInt(25),
					Gas:      20,
					GasPrice: big.NewInt(30000),
					Data:     common.FromHex("0xc4d57bb9b9f95452da96c5d2ca8e3e477672c01afea8362348406a4236fb942f"),
				})
				auth := &bind.TransactOpts{
					From:     common.HexToAddress("0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a"),
					Signer:   nil,
					GasLimit: 20,
					Value:    big.NewInt(0),
					Nonce:    big.NewInt(1),
				}
				blockchain.EXPECT().DocumentSigned(input, timeSign).Return(tx, auth, types.ErrGasFeeCapTooLow).Times(1)
			},
			err: types.ErrGasFeeCapTooLow,
		},
	}
	for _, tt := range test {
		mt.Run(tt.nameTest, func(mt *mtest.T) {
			blockchain := m_blockchain.NewMockBlockchain(ctrl)
			repo := NewRepository(mt.DB, blockchain)
			if tt.testing != nil {
				tt.testing(blockchain, repo)
			}

			err := repo.DocumentSigned(tt.input, tt.times)
			assert.Equal(t, err, tt.err)
		})
	}
}

func Test_repository_ListDocumentNoSign(t *testing.T) {
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

	test := []struct {
		nameTest  string
		publickey string
		output    []models.ListDocument
		response1 primitive.D
		response2 primitive.D
		response3 primitive.D
		response4 primitive.D
		err       error
	}{
		{
			nameTest:  "List Document No Sign Case 1: Success Get List Document Data",
			publickey: "0x8a9c4dfe8b62e51b88291c37e0d6dc15d34dbf1e",
			output: []models.ListDocument{
				{
					Id:               id,
					Address:          "0x8a9c4dfe8b62e51b88291c37e0d6dc15d34dbf1e",
					Hash:             "eusdnlaskjdasodjkasddsj21j3newndakaldajsdoasjdopwud09aud98p",
					Hash_original:    "rnn234mbxjldjihqwoywndoasalksdo3y981mlasdp93uejrk3e213",
					Judul:            "Test Judul 1",
					Note:             "Test Note 1",
					Date_created:     time.Now(),
					Date_created_WIB: "Minggu, 27 Nop 2022 | 11:30 WIB",
				},
				{
					Id:               id2,
					Address:          "0x8a9c4dfe8b62e51b88291c37e0d6dc15d34dbf1e",
					Hash:             "wndakaldajsdoasjdopwud09aud98peus3nednlaskjdasodjkasddsj21j",
					Hash_original:    "hqwoywndoasalksdo3y981mlasdp93ue3jrkrnn234mbxjldjie213",
					Judul:            "Test Judul 2",
					Note:             "Test Note 2",
					Date_created:     time.Now(),
					Date_created_WIB: "Minggu, 27 Nop 2022 | 11:30 WIB",
				},
				{
					Id:               id3,
					Address:          "0x8a9c4dfe8b62e51b88291c37e0d6dc15d34dbf1e",
					Hash:             "dakaldajsdoasjdasddud09audjdasodjkwn98peus3nednlaskjsj21opw",
					Hash_original:    "9qwoywndoasalksdo3y1381mlasdp93ue3jrkrnn234mbxjldjie2h",
					Judul:            "Test Judul 3",
					Note:             "Test Note 3",
					Date_created:     time.Now(),
					Date_created_WIB: "Minggu, 27 Nop 2022 | 11:30 WIB",
				},
				{
					Id:               id4,
					Address:          "0x8a9c4dfe8b62e51b88291c37e0d6dc15d34dbf1e",
					Hash:             "wn98peus3nednlaskjsj21opwdjdasodjkdakaldajsdoasjdasddud09au",
					Hash_original:    "doa8lasdp93ue3jrkrnn234mbxjldjie21mhs9qwoywnalksdo3y13",
					Judul:            "Test Judul 4",
					Note:             "Test Note 4",
					Date_created:     time.Now(),
					Date_created_WIB: "Minggu, 27 Nop 2022 | 11:30 WIB",
				},
			},
			response1: mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, primitive.D{
				{Key: "_id", Value: id},
				{Key: "address", Value: "0x8a9c4dfe8b62e51b88291c37e0d6dc15d34dbf1e"},
				{Key: "hash", Value: "eusdnlaskjdasodjkasddsj21j3newndakaldajsdoasjdopwud09aud98p"},
				{Key: "hash_ori", Value: "rnn234mbxjldjihqwoywndoasalksdo3y981mlasdp93uejrk3e213"},
				{Key: "judul", Value: "Test Judul 1"},
				{Key: "note", Value: "Test Note 1"},
				{Key: "date_created", Value: time.Now()},
				{Key: "date_created_WIB", Value: "Minggu, 27 Nop 2022 | 11:30 WIB"},
			}),
			response2: mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, primitive.D{
				{Key: "_id", Value: id2},
				{Key: "address", Value: "0x8a9c4dfe8b62e51b88291c37e0d6dc15d34dbf1e"},
				{Key: "hash", Value: "wndakaldajsdoasjdopwud09aud98peus3nednlaskjdasodjkasddsj21j"},
				{Key: "hash_ori", Value: "hqwoywndoasalksdo3y981mlasdp93ue3jrkrnn234mbxjldjie213"},
				{Key: "judul", Value: "Test Judul 2"},
				{Key: "note", Value: "Test Note 2"},
				{Key: "date_created", Value: time.Now()},
				{Key: "date_created_WIB", Value: "Minggu, 27 Nop 2022 | 11:30 WIB"},
			}),
			response3: mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, primitive.D{
				{Key: "_id", Value: id3},
				{Key: "address", Value: "0x8a9c4dfe8b62e51b88291c37e0d6dc15d34dbf1e"},
				{Key: "hash", Value: "dakaldajsdoasjdasddud09audjdasodjkwn98peus3nednlaskjsj21opw"},
				{Key: "hash_ori", Value: "9qwoywndoasalksdo3y1381mlasdp93ue3jrkrnn234mbxjldjie2h"},
				{Key: "judul", Value: "Test Judul 3"},
				{Key: "note", Value: "Test Note 3"},
				{Key: "date_created", Value: time.Now()},
				{Key: "date_created_WIB", Value: "Minggu, 27 Nop 2022 | 11:30 WIB"},
			}),
			response4: mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch, primitive.D{
				{Key: "_id", Value: id4},
				{Key: "address", Value: "0x8a9c4dfe8b62e51b88291c37e0d6dc15d34dbf1e"},
				{Key: "hash", Value: "wn98peus3nednlaskjsj21opwdjdasodjkdakaldajsdoasjdasddud09au"},
				{Key: "hash_ori", Value: "doa8lasdp93ue3jrkrnn234mbxjldjie21mhs9qwoywnalksdo3y13"},
				{Key: "judul", Value: "Test Judul 4"},
				{Key: "note", Value: "Test Note 4"},
				{Key: "date_created", Value: time.Now()},
				{Key: "date_created_WIB", Value: "Minggu, 27 Nop 2022 | 11:30 WIB"},
			}),
			err: nil,
		},
		{
			nameTest:  "List Document No Sign Case 2: Error Failed Get List Document Data",
			publickey: "0x8a9c4dfe8b62e51b88291c37e0d6dc15d34dbf1e",
			output:    []models.ListDocument(nil),
			response1: mtest.CreateCommandErrorResponse(mtest.CommandError{
				Code:    50,
				Message: "Database Timeout",
			}),
			err: errors.New("Database Timeout"),
		},
		{
			nameTest:  "List Document No Sign Case 3: Success Get List Document Data Empty",
			publickey: "0x8a9c4dfe8b62e51b88291c37e0d6dc15d34dbf1e",
			output: []models.ListDocument{
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
			signers := repo.ListDocumentNoSign(tt.publickey)
			assert.Equal(t, tt.output, signers)
		})
	}
}

func Test_repository_GetDocument(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()
	times := time.Now()
	timeSign := new(big.Int)
	timeFormat := times.Format("15040502012006")
	timeSign, _ = timeSign.SetString(timeFormat, 10)

	test := []struct {
		nameTest  string
		hash      string
		publickey string
		output    models.DocumentBlockchain
		testing   func(blockchain *m_blockchain.MockBlockchain, repo Repository)
	}{
		{
			nameTest:  "Get Document Case 1: Success Get Document in Blockchain",
			hash:      "doa8lasdp93ue3jrkrnn234mbxjldjie21mhs9qwoywnalksdo3y13",
			publickey: "0x8a9c4dfe8b62e51b88291c37e0d6dc15d34dbf1e",
			output: models.DocumentBlockchain{
				Document_id:    "0x1",
				Creator:        common.HexToAddress("0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a"),
				Creator_string: "0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a",
				Creator_id:     "rizwijaya",
				Metadata:       "sample_test.pdf",
				Hash_ori:       "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b",
				Hash:           "doa8lasdp93ue3jrkrnn234mbxjldjie21mhs9qwoywnalksdo3y13",
				IPFS:           "d9sj84msl02ndm93d8df4d2u43soj3bdsds",
				State:          "2",
				Mode:           "1",
				Createdtime:    timeSign.String(),
				Completedtime:  timeSign.String(),
				Exist:          true,
				Signers: models.Signers{
					Sign_addr:     common.HexToAddress("0x8a9c4dfe8b62e51b88291c37e0d6dc15d34dbf1e"),
					Sign_id:       "1",
					Signers_id:    "rizwijaya",
					Signers_hash:  "9f1bf51bd36a4c244e82419f9d725e15d9cc537106cb54u798sc272b66cda64b",
					Signers_state: true,
					Sign_time:     timeSign.String(),
				},
			},
			testing: func(blockchain *m_blockchain.MockBlockchain, repo Repository) {
				output := models.DocumentBlockchain{
					Document_id:    "0x1",
					Creator:        common.HexToAddress("0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a"),
					Creator_string: "0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a",
					Creator_id:     "rizwijaya",
					Metadata:       "sample_test.pdf",
					Hash_ori:       "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b",
					Hash:           "doa8lasdp93ue3jrkrnn234mbxjldjie21mhs9qwoywnalksdo3y13",
					IPFS:           "d9sj84msl02ndm93d8df4d2u43soj3bdsds",
					State:          "2",
					Mode:           "1",
					Createdtime:    timeSign.String(),
					Completedtime:  timeSign.String(),
					Exist:          true,
					Signers: models.Signers{
						Sign_addr:     common.HexToAddress("0x8a9c4dfe8b62e51b88291c37e0d6dc15d34dbf1e"),
						Sign_id:       "1",
						Signers_id:    "rizwijaya",
						Signers_hash:  "9f1bf51bd36a4c244e82419f9d725e15d9cc537106cb54u798sc272b66cda64b",
						Signers_state: true,
						Sign_time:     timeSign.String(),
					},
				}
				hash := "doa8lasdp93ue3jrkrnn234mbxjldjie21mhs9qwoywnalksdo3y13"
				public_key := "0x8a9c4dfe8b62e51b88291c37e0d6dc15d34dbf1e"
				blockchain.EXPECT().GetDocument(hash, public_key).Return(output).Times(1)
			},
		},
		{
			nameTest:  "Get Document Case 2: Error Failed Get Document in Blockchain",
			hash:      "e3jrkrnn234mbxjldoywnalksdo3y13jie21mhsdoa8lasdp93u9qw",
			publickey: "0x2e51b88291c37e0d6dc15d34db8a9c4dfe8b6f1e",
			output:    models.DocumentBlockchain{},
			testing: func(blockchain *m_blockchain.MockBlockchain, repo Repository) {
				output := models.DocumentBlockchain{}
				hash := "e3jrkrnn234mbxjldoywnalksdo3y13jie21mhsdoa8lasdp93u9qw"
				public_key := "0x2e51b88291c37e0d6dc15d34db8a9c4dfe8b6f1e"
				blockchain.EXPECT().GetDocument(hash, public_key).Return(output).Times(1)
			},
		},
	}
	for _, tt := range test {
		mt.Run(tt.nameTest, func(mt *mtest.T) {
			blockchain := m_blockchain.NewMockBlockchain(ctrl)
			repo := NewRepository(mt.DB, blockchain)
			if tt.testing != nil {
				tt.testing(blockchain, repo)
			}

			output := repo.GetDocument(tt.hash, tt.publickey)
			assert.Equal(t, output, tt.output)
		})
	}
}

func Test_repository_GetSigners(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()
	times := time.Now()
	timeSign := new(big.Int)
	timeFormat := times.Format("15040502012006")
	timeSign, _ = timeSign.SetString(timeFormat, 10)

	test := []struct {
		nameTest  string
		hash      string
		publickey string
		output    models.Signers
		testing   func(blockchain *m_blockchain.MockBlockchain, repo Repository)
	}{
		{
			nameTest:  "Get Signers Case 1: Success Get Signers Data in Blockchain",
			hash:      "doa8lasdp93ue3jrkrnn234mbxjldjie21mhs9qwoywnalksdo3y13",
			publickey: "0x8a9c4dfe8b62e51b88291c37e0d6dc15d34dbf1e",
			output: models.Signers{
				Sign_addr:     common.HexToAddress("0x8a9c4dfe8b62e51b88291c37e0d6dc15d34dbf1e"),
				Sign_id:       "1",
				Signers_id:    "rizwijaya",
				Signers_hash:  "9f1bf51bd36a4c244e82419f9d725e15d9cc537106cb54u798sc272b66cda64b",
				Signers_state: true,
				Sign_time:     timeSign.String(),
			},
			testing: func(blockchain *m_blockchain.MockBlockchain, repo Repository) {
				output := models.Signers{
					Sign_addr:     common.HexToAddress("0x8a9c4dfe8b62e51b88291c37e0d6dc15d34dbf1e"),
					Sign_id:       "1",
					Signers_id:    "rizwijaya",
					Signers_hash:  "9f1bf51bd36a4c244e82419f9d725e15d9cc537106cb54u798sc272b66cda64b",
					Signers_state: true,
					Sign_time:     timeSign.String(),
				}
				hash := "doa8lasdp93ue3jrkrnn234mbxjldjie21mhs9qwoywnalksdo3y13"
				public_key := "0x8a9c4dfe8b62e51b88291c37e0d6dc15d34dbf1e"
				blockchain.EXPECT().GetSigners(hash, public_key).Return(output).Times(1)
			},
		},
		{
			nameTest:  "Get Signers Case 2: Error Failed Get Signers Data in Blockchain",
			hash:      "e3jrkrnn234mbxjldoywnalksdo3y13jie21mhsdoa8lasdp93u9qw",
			publickey: "0x2e51b88291c37e0d6dc15d34db8a9c4dfe8b6f1e",
			output:    models.Signers{},
			testing: func(blockchain *m_blockchain.MockBlockchain, repo Repository) {
				output := models.Signers{}
				hash := "e3jrkrnn234mbxjldoywnalksdo3y13jie21mhsdoa8lasdp93u9qw"
				public_key := "0x2e51b88291c37e0d6dc15d34db8a9c4dfe8b6f1e"
				blockchain.EXPECT().GetSigners(hash, public_key).Return(output).Times(1)
			},
		},
	}
	for _, tt := range test {
		mt.Run(tt.nameTest, func(mt *mtest.T) {
			blockchain := m_blockchain.NewMockBlockchain(ctrl)
			repo := NewRepository(mt.DB, blockchain)
			if tt.testing != nil {
				tt.testing(blockchain, repo)
			}

			output := repo.GetSigners(tt.hash, tt.publickey)
			assert.Equal(t, output, tt.output)
		})
	}
}

func Test_repository_GetHashOriginal(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()
	times := time.Now()
	timeSign := new(big.Int)
	timeFormat := times.Format("15040502012006")
	timeSign, _ = timeSign.SetString(timeFormat, 10)

	test := []struct {
		nameTest  string
		hash      string
		publickey string
		output    string
		testing   func(blockchain *m_blockchain.MockBlockchain, repo Repository)
	}{
		{
			nameTest:  "Get Hash Original Document Case 1: Success Get Hash Original Document in Blockchain",
			hash:      "doa8lasdp93ue3jrkrnn234mbxjldjie21mhs9qwoywnalksdo3y13",
			publickey: "0x8a9c4dfe8b62e51b88291c37e0d6dc15d34dbf1e",
			output:    "9f1bf51bd36a4c244e82419f9d725e15d9cc537106cb54u798sc272b66cda64b",
			testing: func(blockchain *m_blockchain.MockBlockchain, repo Repository) {
				output := "9f1bf51bd36a4c244e82419f9d725e15d9cc537106cb54u798sc272b66cda64b"
				hash := "doa8lasdp93ue3jrkrnn234mbxjldjie21mhs9qwoywnalksdo3y13"
				public_key := "0x8a9c4dfe8b62e51b88291c37e0d6dc15d34dbf1e"
				blockchain.EXPECT().GetHashOriginal(hash, public_key).Return(output).Times(1)
			},
		},
		{
			nameTest:  "Get Hash Original Document Case 2: Error Failed Get Hash Original Document in Blockchain",
			hash:      "e3jrkrnn234mbxjldoywnalksdo3y13jie21mhsdoa8lasdp93u9qw",
			publickey: "0x2e51b88291c37e0d6dc15d34db8a9c4dfe8b6f1e",
			output:    "",
			testing: func(blockchain *m_blockchain.MockBlockchain, repo Repository) {
				output := ""
				hash := "e3jrkrnn234mbxjldoywnalksdo3y13jie21mhsdoa8lasdp93u9qw"
				public_key := "0x2e51b88291c37e0d6dc15d34db8a9c4dfe8b6f1e"
				blockchain.EXPECT().GetHashOriginal(hash, public_key).Return(output).Times(1)
			},
		},
	}
	for _, tt := range test {
		mt.Run(tt.nameTest, func(mt *mtest.T) {
			blockchain := m_blockchain.NewMockBlockchain(ctrl)
			repo := NewRepository(mt.DB, blockchain)
			if tt.testing != nil {
				tt.testing(blockchain, repo)
			}

			output := repo.GetHashOriginal(tt.hash, tt.publickey)
			assert.Equal(t, output, tt.output)
		})
	}
}

func Test_repository_GetListSign(t *testing.T) {
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

	test := []struct {
		nameTest  string
		hash      string
		output    []models.SignersData
		response1 primitive.D
		response2 primitive.D
		response3 primitive.D
		response4 primitive.D
		err       error
	}{
		{
			nameTest: "Get List Signatures Case 1: Success Get List Signatures Data",
			hash:     "58s9t0u1v2w3x4y5z89cp4b4c3b5c0e9b7d6q7r0f2e6f1e2f3a4b5c6d4e5f6g7h8i9j0k1l2m3n4o6d7e8f9a1b2c3",
			output: []models.SignersData{
				{
					Sign_addr:    "0x5B38Da6a701c568545dCfcB03FcB875f56beddC4",
					Sign_name:    "Administrator",
					Sign_email:   "admin@localhost",
					Sign_id_db:   id.Hex(),
					Signers_id:   "admin",
					Signers_hash: "58s9t0u1v2w3x4y5z89cp4b4c3b5c0e9b7d6q7r0f2e6f1e2f3a4b5c6d4e5f6g7h8i9j0k1l2m3n4o6d7e8f9a1b2c3",
				},
				{
					Sign_addr:    "0xD3CdA913deB6f67967B99D67aCDFa1712C293601",
					Sign_name:    "Rizqi Wijaya",
					Sign_email:   "smartsign@rizwijaya.com",
					Sign_id_db:   id2.Hex(),
					Signers_id:   "rizwijaya",
					Signers_hash: "58s9t0u1v2w3x4y5z89cp4b4c3b5c0e9b7d6q7r0f2e6f1e2f3a4b5c6d4e5f6g7h8i9j0k1l2m3n4o6d7e8f9a1b2c3",
				},
				{
					Sign_addr:    "0xAyysae6513c99443cF32Ca8A449f5287aaD6f91a",
					Sign_name:    "Sembada Oke",
					Sign_email:   "sembada@rizwijaya.com",
					Sign_id_db:   id3.Hex(),
					Signers_id:   "sembada",
					Signers_hash: "58s9t0u1v2w3x4y5z89cp4b4c3b5c0e9b7d6q7r0f2e6f1e2f3a4b5c6d4e5f6g7h8i9j0k1l2m3n4o6d7e8f9a1b2c3",
				},
				{
					Sign_addr:    "0xD6fAyysae91a6513c99443cF32Ca8A449f5287aa",
					Sign_name:    "Sitaman Sutama",
					Sign_email:   "sitaman@rizwijaya.com",
					Sign_id_db:   id4.Hex(),
					Signers_id:   "sitama",
					Signers_hash: "58s9t0u1v2w3x4y5z89cp4b4c3b5c0e9b7d6q7r0f2e6f1e2f3a4b5c6d4e5f6g7h8i9j0k1l2m3n4o6d7e8f9a1b2c3",
				},
			},
			response1: mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, primitive.D{
				{Key: "_id", Value: id},
				{Key: "address", Value: "0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"},
				{Key: "name", Value: "Administrator"},
				{Key: "email", Value: "admin@localhost"},
				{Key: "idsignature", Value: "admin"},
				{Key: "hash", Value: "58s9t0u1v2w3x4y5z89cp4b4c3b5c0e9b7d6q7r0f2e6f1e2f3a4b5c6d4e5f6g7h8i9j0k1l2m3n4o6d7e8f9a1b2c3"},
			}),
			response2: mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, primitive.D{
				{Key: "_id", Value: id2},
				{Key: "address", Value: "0xD3CdA913deB6f67967B99D67aCDFa1712C293601"},
				{Key: "name", Value: "Rizqi Wijaya"},
				{Key: "email", Value: "smartsign@rizwijaya.com"},
				{Key: "idsignature", Value: "rizwijaya"},
				{Key: "hash", Value: "58s9t0u1v2w3x4y5z89cp4b4c3b5c0e9b7d6q7r0f2e6f1e2f3a4b5c6d4e5f6g7h8i9j0k1l2m3n4o6d7e8f9a1b2c3"},
			}),
			response3: mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, primitive.D{
				{Key: "_id", Value: id3},
				{Key: "address", Value: "0xAyysae6513c99443cF32Ca8A449f5287aaD6f91a"},
				{Key: "name", Value: "Sembada Oke"},
				{Key: "email", Value: "sembada@rizwijaya.com"},
				{Key: "idsignature", Value: "sembada"},
				{Key: "hash", Value: "58s9t0u1v2w3x4y5z89cp4b4c3b5c0e9b7d6q7r0f2e6f1e2f3a4b5c6d4e5f6g7h8i9j0k1l2m3n4o6d7e8f9a1b2c3"},
			}),
			response4: mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch, primitive.D{
				{Key: "_id", Value: id4},
				{Key: "address", Value: "0xD6fAyysae91a6513c99443cF32Ca8A449f5287aa"},
				{Key: "name", Value: "Sitaman Sutama"},
				{Key: "email", Value: "sitaman@rizwijaya.com"},
				{Key: "idsignature", Value: "sitama"},
				{Key: "hash", Value: "58s9t0u1v2w3x4y5z89cp4b4c3b5c0e9b7d6q7r0f2e6f1e2f3a4b5c6d4e5f6g7h8i9j0k1l2m3n4o6d7e8f9a1b2c3"},
			}),
			err: nil,
		},
		{
			nameTest: "Get List Signatures Case 2: Error Failed Get List Signatures Data",
			hash:     "6f1e2f3a4b5c6d4e5f6g7h8i9j0k7e8f9a1b2c3b7d58s9t0u1v2w3x4y5z89cp4b4c3b5c0e96q7r0f2e1l2m3n4o6d",
			output:   []models.SignersData(nil),
			response1: mtest.CreateCommandErrorResponse(mtest.CommandError{
				Code:    50,
				Message: "Database Timeout",
			}),
			err: errors.New("Database Timeout"),
		},
		{
			nameTest: "Get List Signatures Case 2: Success Get List Signatures Data with Empty Data",
			hash:     "c0e9bu1v2w3x4y5z809cp4b4c3b57d6q7r0f2e6f1e2f3a4b5c6d4e5f6g7h8i9j0k1l2m3n4o6d7e8f9a1b258s9tc3",
			output: []models.SignersData{
				{
					Sign_id_db: "",
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
			signers := repo.GetListSign(tt.hash)
			assert.Equal(t, tt.output, signers)
		})
	}
}

func Test_repository_GetUserByIdSignatures(t *testing.T) {
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
		output      modelsUser.ProfileDB
		response    primitive.D
		err         error
	}{
		{
			nameTest:    "Get User By Id Signature Case 1: Success Get User Data Exist",
			idsignature: "admin",
			output: modelsUser.ProfileDB{
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
			nameTest:    "Get User By Id Signature Case 1: Success Get User Data Not Exist",
			idsignature: "admin23",
			output:      modelsUser.ProfileDB{},
			response:    mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, primitive.D{}),
			err:         nil,
		},
		{
			nameTest:    "Get User By Id Signature Case 1: Error Failed Decoded Data User",
			idsignature: "adminbrow",
			output:      modelsUser.ProfileDB{},
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
			user := repo.GetUserByIdSignatures(tt.idsignature)
			assert.Equal(t, tt.output, user)
		})
	}
}

func Test_repository_VerifyDoc(t *testing.T) {
	t.Skip("Skip Test Verify Doc")
}

func Test_repository_GetTransactions(t *testing.T) {
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
		output      []models.Transac
		response1   primitive.D
		response2   primitive.D
		response3   primitive.D
		response4   primitive.D
		err         error
	}{
		{
			nameTest:    "Get Transactions Case 1: Success Get Transactions Data",
			idsignature: "admin",
			output: []models.Transac{
				{
					Id:               id,
					Address:          "0x5B38Da6a701c568545dCfcB03FcB875f56beddC4",
					Tx_hash:          "0x5f44e265dbf57984ffb9a833ba9cde9c51a6bec419c44f8e40b64a9ee7033c83",
					Nonce:            "3",
					Prices:           "30000",
					Description:      "Membuat Dokumen sample.pdf untuk tanda tangan",
					Date_created:     times,
					Date_created_wib: "Minggu, 27 Nop 2022 | 11:30 WIB",
				},
				{
					Id:               id2,
					Address:          "0xDB5Ad0D8c8acE3A9eB8fD530eD8F5254044c9fA0",
					Tx_hash:          "0x5c53f765f935ae06dfdbaf820f066ce01ad30acb77bcb4d6a013d9cfe68c2b5a",
					Nonce:            "4",
					Prices:           "30000",
					Description:      "Membuat Dokumen Sample_testing.pdf untuk tanda tangan",
					Date_created:     times,
					Date_created_wib: "Minggu, 27 Nop 2022 | 11:30 WIB",
				},
				{
					Id:               id3,
					Address:          "0xy5B38Da6a701c568545dCfcB03FcB875f56beddC4",
					Tx_hash:          "0x4b47a3f5c24db6363dd841c9e009c238274bf6f0439389f9654e90395d9e7388",
					Nonce:            "5",
					Prices:           "30000",
					Description:      "Menandatangani Dokumen dengan kode : 3a0233f815f46edd8afae31413b62c6a55861ab47bc90db4dbd085735ae47ff1",
					Date_created:     times,
					Date_created_wib: "Minggu, 27 Nop 2022 | 11:30 WIB",
				},
				{
					Id:               id4,
					Address:          "0xC5fdf4076b8F3A5357c5E395ab970B5B54098Fef",
					Tx_hash:          "0xe1301136ace8d8a780df304dad4525d244fc13dc9f316bafe92e2734d2f369e8",
					Nonce:            "6",
					Prices:           "30000",
					Description:      "Menandatangani Dokumen dengan kode : cbad53ee065af3beab98fd85062076cd9d1cf38fac5760d5051080c3096bf69c",
					Date_created:     times,
					Date_created_wib: "Minggu, 27 Nop 2022 | 11:30 WIB",
				},
			},
			response1: mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, primitive.D{
				{Key: "_id", Value: id},
				{Key: "address", Value: "0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"},
				{Key: "tx_hash", Value: "0x5f44e265dbf57984ffb9a833ba9cde9c51a6bec419c44f8e40b64a9ee7033c83"},
				{Key: "nonce", Value: "3"},
				{Key: "prices", Value: "30000"},
				{Key: "description", Value: "Membuat Dokumen sample.pdf untuk tanda tangan"},
				{Key: "date_created", Value: times},
				{Key: "date_created_wib", Value: "Minggu, 27 Nop 2022 | 11:30 WIB"},
			}),
			response2: mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, primitive.D{
				{Key: "_id", Value: id2},
				{Key: "address", Value: "0xDB5Ad0D8c8acE3A9eB8fD530eD8F5254044c9fA0"},
				{Key: "tx_hash", Value: "0x5c53f765f935ae06dfdbaf820f066ce01ad30acb77bcb4d6a013d9cfe68c2b5a"},
				{Key: "nonce", Value: "4"},
				{Key: "prices", Value: "30000"},
				{Key: "description", Value: "Membuat Dokumen Sample_testing.pdf untuk tanda tangan"},
				{Key: "date_created", Value: times},
				{Key: "date_created_wib", Value: "Minggu, 27 Nop 2022 | 11:30 WIB"},
			}),
			response3: mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, primitive.D{
				{Key: "_id", Value: id3},
				{Key: "address", Value: "0xy5B38Da6a701c568545dCfcB03FcB875f56beddC4"},
				{Key: "tx_hash", Value: "0x4b47a3f5c24db6363dd841c9e009c238274bf6f0439389f9654e90395d9e7388"},
				{Key: "nonce", Value: "5"},
				{Key: "prices", Value: "30000"},
				{Key: "description", Value: "Menandatangani Dokumen dengan kode : 3a0233f815f46edd8afae31413b62c6a55861ab47bc90db4dbd085735ae47ff1"},
				{Key: "date_created", Value: times},
				{Key: "date_created_wib", Value: "Minggu, 27 Nop 2022 | 11:30 WIB"},
			}),
			response4: mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch, primitive.D{
				{Key: "_id", Value: id4},
				{Key: "address", Value: "0xC5fdf4076b8F3A5357c5E395ab970B5B54098Fef"},
				{Key: "tx_hash", Value: "0xe1301136ace8d8a780df304dad4525d244fc13dc9f316bafe92e2734d2f369e8"},
				{Key: "nonce", Value: "6"},
				{Key: "prices", Value: "30000"},
				{Key: "description", Value: "Menandatangani Dokumen dengan kode : cbad53ee065af3beab98fd85062076cd9d1cf38fac5760d5051080c3096bf69c"},
				{Key: "date_created", Value: times},
				{Key: "date_created_wib", Value: "Minggu, 27 Nop 2022 | 11:30 WIB"},
			}),
			err: nil,
		},
		{
			nameTest:    "Get Transactions Case 2: Error Failed Get Transactions Data",
			idsignature: "admin",
			output:      []models.Transac(nil),
			response1: mtest.CreateCommandErrorResponse(mtest.CommandError{
				Code:    50,
				Message: "Database Timeout",
			}),
			err: errors.New("Database Timeout"),
		},
		{
			nameTest:    "Get Transactions Case 3: Success Get Transactions Data with empty data",
			idsignature: "admin",
			output: []models.Transac{
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
			transac := repo.GetTransactions()
			assert.Equal(t, tt.output, transac)
		})
	}
}

func Test_repository_CheckSignature(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	test := []struct {
		nameTest  string
		hash      string
		publickey string
		output    bool
		response  primitive.D
		err       error
	}{
		{
			nameTest:  "Check Signature Case 1: Success Signature Data Document Exist",
			hash:      "dslajdjasodu038e29iojnewkd7a8d6y9a8hondl23123dasdassdas97",
			publickey: "0x8d6y9a8hondl23123dasdassdas97",
			output:    true,
			response: mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, primitive.D{
				{Key: "hash", Value: "dslajdjasodu038e29iojnewkd7a8d6y9a8hondl23123dasdassdas97"},
			}),
			err: nil,
		},
		{
			nameTest:  "Check Signature Case 2: Success Signature Data Document Not Exist",
			publickey: "0xdassdas979a8hondl231238d6ydas",
			output:    false,
			response:  mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, primitive.D{}),
			err:       nil,
		},
		{
			nameTest:  "Check Signature Case 3: Error Failed Decoded Data Document",
			publickey: "0xdassdas979a8hondl231238d6ydas",
			output:    false,
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
			output := repo.CheckSignature(tt.hash, tt.publickey)
			assert.Equal(t, tt.output, output)
		})
	}
}
