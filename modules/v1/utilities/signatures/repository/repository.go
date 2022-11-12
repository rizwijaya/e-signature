package repository

import (
	"context"
	"e-signature/modules/v1/utilities/signatures/models"
	modelsUser "e-signature/modules/v1/utilities/user/models"
	"fmt"
	"log"
	"math/big"
	"time"

	blockhainAuth "e-signature/app/blockhain"
	"e-signature/app/config"
	api "e-signature/app/contracts"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	LogTransactions(address string, tx_hash string, nonce string, desc string, prices string) error
	DefaultSignatures(user modelsUser.User, id string) error
	UpdateMySignatures(signature string, signaturedata string, sign string) error
	GetMySignature(sign string) (models.Signatures, error)
	ChangeSignature(sign_type string, sign string) error
	AddToBlockhain(input models.SignDocuments, times *big.Int) error
	AddUserDocs(input models.SignDocuments) error
	DocumentSigned(sign models.SignDocs, timeSign *big.Int) error
}

type repository struct {
	db         *mongo.Database
	blockchain *api.Api
	client     *ethclient.Client
}

func NewRepository(db *mongo.Database, blockchain *api.Api, client *ethclient.Client) *repository {
	return &repository{db, blockchain, client}
}

func (r *repository) LogTransactions(address string, tx_hash string, nonce string, desc string, prices string) error {
	ctx := context.TODO()
	trans := models.Transac{
		Id:           primitive.NewObjectID(),
		Address:      fmt.Sprintf("%v", address),
		Tx_hash:      fmt.Sprintf("%v", tx_hash),
		Nonce:        nonce,
		Prices:       prices,
		Description:  desc,
		Date_created: time.Now(),
	}
	c := r.db.Collection("transactions")
	_, err := c.InsertOne(ctx, &trans)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (r *repository) DefaultSignatures(user modelsUser.User, id string) error {
	signatures := models.Signatures{
		Id:                 primitive.NewObjectID(),
		User:               user.Idsignature,
		Signature:          "default.png",
		Signature_data:     "default.png",
		Latin:              fmt.Sprintf("latin-%s.png", id),
		Latin_data:         fmt.Sprintf("latindata-%s.png", id),
		Signature_selected: "latin",
		Date_update:        time.Now(),
		Date_created:       time.Now(),
	}

	c := r.db.Collection("signatures")
	_, err := c.InsertOne(context.Background(), &signatures)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (r *repository) UpdateMySignatures(signature string, signaturedata string, sign string) error {
	filter := map[string]interface{}{"user": sign}
	update := map[string]interface{}{
		"$set": map[string]interface{}{
			"signature":          signature,
			"signature_data":     signaturedata,
			"signature_selected": "signature",
			"date_update":        time.Now(),
		},
	}

	c := r.db.Collection("signatures")
	_, err := c.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (r *repository) GetMySignature(sign string) (models.Signatures, error) {
	var signatures models.Signatures
	filter := map[string]interface{}{"user": sign}
	c := r.db.Collection("signatures")
	err := c.FindOne(context.Background(), filter).Decode(&signatures)
	if err != nil {
		log.Fatal(err)
		return signatures, err
	}
	return signatures, nil
}

func (r *repository) ChangeSignature(sign_type string, sign string) error {
	filter := map[string]interface{}{"user": sign}
	update := map[string]interface{}{
		"$set": map[string]interface{}{
			"signature_selected": sign_type,
			"date_update":        time.Now(),
		},
	}

	c := r.db.Collection("signatures")
	_, err := c.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (r *repository) AddToBlockhain(input models.SignDocuments, times *big.Int) error {
	conf, _ := config.Init()
	auth := blockhainAuth.GetAccountAuth(blockhainAuth.Connect(), conf.Blockhain.Secret_key)
	document, err := r.blockchain.Create(auth, input.Hash_original, common.HexToAddress(input.Creator), input.Name, input.Hash, input.IPFS, big.NewInt(1), false, times, input.Address, input.IdSignature)
	if err != nil {
		log.Println(err)
	}
	//Logging transaksi.
	r.LogTransactions(input.Creator, document.Hash().Hex(), auth.Nonce.String(), "Membuat Dokumen "+input.Name+" untuk tanda tangan", document.Cost().String())
	return err
}

func (r *repository) AddUserDocs(input models.SignDocuments) error {
	for _, v := range input.Address {
		signedDocuments := struct {
			Id           primitive.ObjectID `bson:"_id,omitempty"`
			Address      string             `bson:"address"`
			Hash         string             `bson:"hash"`
			Date_Created time.Time          `bson:"date_created"`
		}{
			Id:           primitive.NewObjectID(),
			Address:      v.String(),
			Hash:         input.Hash,
			Date_Created: time.Now(),
		}

		c := r.db.Collection("signedDocuments")
		_, err := c.InsertOne(context.Background(), &signedDocuments)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}

func (r *repository) DocumentSigned(sign models.SignDocs, timeSign *big.Int) error {
	conf, _ := config.Init()
	auth := blockhainAuth.GetAccountAuth(blockhainAuth.Connect(), conf.Blockhain.Secret_key)
	signDocs, err := r.blockchain.SignDoc(auth, sign.Hash_original, common.HexToAddress(sign.Creator), sign.Hash, sign.IPFS, timeSign)
	//Logging Transaction
	r.LogTransactions(sign.Creator, signDocs.Hash().Hex(), auth.Nonce.String(), "Menandatangani Dokumen dengan kode : "+sign.Hash_original, signDocs.Cost().String())
	return err
}
