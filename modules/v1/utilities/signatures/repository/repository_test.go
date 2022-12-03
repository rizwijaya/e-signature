package repository

import (
	"e-signature/modules/v1/utilities/signatures/models"
	modelsUser "e-signature/modules/v1/utilities/user/models"
	m_blockchain "e-signature/pkg/blockchain/mock"
	"errors"
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
	t.Skip("Skip Test Add To Blockchain")
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
	t.Skip("Skip Test Document Signed")
}

func Test_repository_ListDocumentNoSign(t *testing.T) {
	t.Skip("Skip Test List Document No Sign")
}

func Test_repository_GetDocument(t *testing.T) {
	t.Skip("Skip Test Get Document")
}

func Test_repository_GetListSign(t *testing.T) {
	t.Skip("Skip Test Get List Sign")
}

func Test_repository_GetHashOriginal(t *testing.T) {
	t.Skip("Skip Test Get Hash Original")
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
