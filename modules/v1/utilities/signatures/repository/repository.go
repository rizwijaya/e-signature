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

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
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
	ListDocumentNoSign(publickey string) []models.ListDocument
	GetDocument(hash string, publickey string) models.DocumentBlockchain
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
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Fatal(err)
		return err
	}
	ctx := context.TODO()
	trans := models.Transac{
		Id:           primitive.NewObjectID(),
		Address:      fmt.Sprintf("%v", address),
		Tx_hash:      fmt.Sprintf("%v", tx_hash),
		Nonce:        nonce,
		Prices:       prices,
		Description:  desc,
		Date_created: time.Now().In(location),
	}
	c := r.db.Collection("transactions")
	_, err = c.InsertOne(ctx, &trans)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (r *repository) DefaultSignatures(user modelsUser.User, id string) error {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Fatal(err)
		return err
	}
	signatures := models.Signatures{
		Id:                 primitive.NewObjectID(),
		User:               user.Idsignature,
		Signature:          "default.png",
		Signature_data:     "default.png",
		Latin:              fmt.Sprintf("latin-%s.png", id),
		Latin_data:         fmt.Sprintf("latindata-%s.png", id),
		Signature_selected: "latin",
		Date_update:        time.Now().In(location),
		Date_created:       time.Now().In(location),
	}

	c := r.db.Collection("signatures")
	_, err = c.InsertOne(context.Background(), &signatures)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (r *repository) UpdateMySignatures(signature string, signaturedata string, sign string) error {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Fatal(err)
		return err
	}
	filter := map[string]interface{}{"user": sign}
	update := map[string]interface{}{
		"$set": map[string]interface{}{
			"signature":          signature,
			"signature_data":     signaturedata,
			"signature_selected": "signature",
			"date_update":        time.Now().In(location),
		},
	}

	c := r.db.Collection("signatures")
	_, err = c.UpdateOne(context.Background(), filter, update)
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
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Fatal(err)
		return err
	}
	filter := map[string]interface{}{"user": sign}
	update := map[string]interface{}{
		"$set": map[string]interface{}{
			"signature_selected": sign_type,
			"date_update":        time.Now().In(location),
		},
	}

	c := r.db.Collection("signatures")
	_, err = c.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (r *repository) AddToBlockhain(input models.SignDocuments, times *big.Int) error {
	conf, _ := config.Init()
	auth := blockhainAuth.GetAccountAuth(blockhainAuth.Connect(), conf.Blockhain.Secret_key)
	document, err := r.blockchain.Create(auth, input.Hash_original, common.HexToAddress(input.Creator), input.Creator_id, input.Name, input.Hash, input.IPFS, big.NewInt(1), false, times, input.Address, input.IdSignature)
	if err != nil {
		log.Println(err)
	}
	//Logging transaksi.
	r.LogTransactions(input.Creator, document.Hash().Hex(), auth.Nonce.String(), "Membuat Dokumen "+input.Name+" untuk tanda tangan", document.Cost().String())
	return err
}

func (r *repository) AddUserDocs(input models.SignDocuments) error {
	if input.Judul == "" {
		input.Judul = input.Name
	}
	if input.Note == "" {
		input.Note = "Tidak ada catatan"
	}
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Fatal(err)
		return err
	}
	for _, v := range input.Address {
		signedDocuments := struct {
			Id           primitive.ObjectID `bson:"_id,omitempty"`
			Address      string             `bson:"address"`
			Hash         string             `bson:"hash"`
			Judul        string             `bson:"judul"`
			Note         string             `bson:"note"`
			Date_Created time.Time          `bson:"date_created"`
		}{
			Id:           primitive.NewObjectID(),
			Address:      v.String(),
			Hash:         input.Hash,
			Judul:        input.Judul,
			Note:         input.Note,
			Date_Created: time.Now().In(location),
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

func (r *repository) ListDocumentNoSign(publickey string) []models.ListDocument {
	var listDocument []models.ListDocument
	c := r.db.Collection("signedDocuments")
	cursor, err := c.Find(context.Background(), bson.M{"address": publickey})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.Background()) {
		var document models.ListDocument
		err := cursor.Decode(&document)
		if err != nil {
			log.Fatal(err)
		}
		listDocument = append(listDocument, document)
	}
	return listDocument
}

func (r *repository) GetDocument(hash string, publickey string) models.DocumentBlockchain {
	var doc models.DocumentBlockchain
	var err error
	var document_id, state, Createdtime, Completedtime *big.Int
	document_id, doc.Creator, doc.Creator_id, doc.Metadata, doc.Hash, doc.IPFS, state, doc.Visibility, Createdtime, Completedtime, doc.Exist, err = r.blockchain.GetDoc(&bind.CallOpts{From: common.HexToAddress(publickey)}, hash)
	if err != nil {
		log.Println("Error Get Document")
	}
	doc.Document_id = document_id.String()
	doc.State = state.String()
	doc.Createdtime = Createdtime.String()
	doc.Completedtime = Completedtime.String()

	return doc
}
