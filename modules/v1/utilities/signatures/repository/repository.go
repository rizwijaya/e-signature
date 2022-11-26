package repository

import (
	"context"
	blockhainAuth "e-signature/app/blockhain"
	"e-signature/app/config"
	api "e-signature/app/contracts"
	"e-signature/modules/v1/utilities/signatures/models"
	modelsUser "e-signature/modules/v1/utilities/user/models"
	tm "e-signature/pkg/time"
	"fmt"
	"log"
	"math/big"
	"time"

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
	GetSigners(hash string, publickey string) models.Signers
	GetHashOriginal(hash string, publickey string) string
	GetListSign(hash string) []models.SignersData
	GetUserByIdSignatures(idsignature string) modelsUser.ProfileDB
	VerifyDoc(hash string) bool
	GetTransactions() []models.Transac
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
		log.Println(err)
		return err
	}
	tim := time.Now().In(location)
	ctx := context.TODO()
	trans := models.Transac{
		Id:               primitive.NewObjectID(),
		Address:          fmt.Sprintf("%v", address),
		Tx_hash:          fmt.Sprintf("%v", tx_hash),
		Nonce:            nonce,
		Prices:           prices,
		Description:      desc,
		Date_created:     tim,
		Date_created_wib: tm.TanggalJam(tim),
	}
	c := r.db.Collection("transactions")
	_, err = c.InsertOne(ctx, &trans)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *repository) DefaultSignatures(user modelsUser.User, id string) error {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Println(err)
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
		log.Println(err)
		return err
	}
	return nil
}

func (r *repository) UpdateMySignatures(signature string, signaturedata string, sign string) error {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Println(err)
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
		log.Println(err)
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
		log.Println(err)
		return signatures, err
	}
	return signatures, nil
}

func (r *repository) ChangeSignature(sign_type string, sign string) error {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Println(err)
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
		log.Println(err)
		return err
	}
	return nil
}

func (r *repository) AddToBlockhain(input models.SignDocuments, times *big.Int) error {
	conf, _ := config.Init()
	mode := new(big.Int)
	mode.SetString(input.Mode, 10)
	auth := blockhainAuth.GetAccountAuth(blockhainAuth.Connect(), conf.Blockhain.Secret_key)
	document, err := r.blockchain.Create(auth, input.Hash_original, common.HexToAddress(input.Creator), input.Creator_id, input.Name, input.Hash, input.IPFS, big.NewInt(1), mode, times, input.Address, input.IdSignature)
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
		log.Println(err)
		return err
	}
	for _, v := range input.Address {
		signedDocuments := struct {
			Id               primitive.ObjectID `bson:"_id,omitempty"`
			Address          string             `bson:"address"`
			Hash_Ori         string             `bson:"hash_ori"`
			Hash             string             `bson:"hash"`
			Judul            string             `bson:"judul"`
			Note             string             `bson:"note"`
			Date_Created     time.Time          `bson:"date_created"`
			Date_Created_wib string             `bson:"date_created_wib"`
		}{
			Id:               primitive.NewObjectID(),
			Address:          v.String(),
			Hash_Ori:         input.Hash_original,
			Hash:             input.Hash,
			Judul:            input.Judul,
			Note:             input.Note,
			Date_Created:     time.Now().In(location),
			Date_Created_wib: tm.TanggalJam(time.Now().In(location)),
		}

		c := r.db.Collection("signedDocuments")
		_, err := c.InsertOne(context.Background(), &signedDocuments)
		if err != nil {
			log.Println(err)
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
		log.Println(err)
	}
	for cursor.Next(context.Background()) {
		var document models.ListDocument
		err := cursor.Decode(&document)
		if err != nil {
			log.Println(err)
		}
		listDocument = append(listDocument, document)
	}
	return listDocument
}

func (r *repository) GetDocument(hash string, publickey string) models.DocumentBlockchain {
	var doc models.DocumentBlockchain
	var err error
	var document_id, state, Createdtime, Completedtime, Mode *big.Int
	document_id, doc.Creator, doc.Creator_id, doc.Metadata, doc.Hash_ori, doc.Hash, doc.IPFS, state, Mode, Createdtime, Completedtime, doc.Exist, err = r.blockchain.GetDoc(&bind.CallOpts{From: common.HexToAddress(publickey)}, hash)
	if err != nil {
		log.Println("Error Get Document")
	}
	doc.Document_id = document_id.String()
	doc.State = state.String()
	doc.Createdtime = Createdtime.String()
	doc.Completedtime = Completedtime.String()
	doc.Mode = Mode.String()
	doc.Creator_string = doc.Creator.String()

	return doc
}

func (r *repository) GetSigners(hash string, publickey string) models.Signers {
	var signers models.Signers
	var err error
	var sign_id, sign_time *big.Int
	sign_id, signers.Signers_id, signers.Signers_hash, signers.Signers_state, sign_time, err = r.blockchain.GetSign(&bind.CallOpts{From: common.HexToAddress(publickey)}, hash, common.HexToAddress(publickey))
	if err != nil {
		log.Println("Error Get Signers")
	}
	signers.Sign_id = sign_id.String()
	signers.Sign_time = sign_time.String()

	return signers
}

func (r *repository) GetHashOriginal(hash string, publickey string) string {
	hash_ori, err := r.blockchain.GetDocSigned(&bind.CallOpts{From: common.HexToAddress(publickey)}, hash)
	if err != nil {
		log.Println(err)
		log.Println("Error Get Hash Original")
	}
	return hash_ori
}

func (r *repository) GetListSign(hash string) []models.SignersData {
	var sign []models.SignersData
	filter := bson.M{"hash_ori": hash}
	c := r.db.Collection("signedDocuments")
	cursor, err := c.Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	for cursor.Next(context.Background()) {
		var signer models.SignersData
		err := cursor.Decode(&signer)
		if err != nil {
			log.Println(err)
		}
		sign = append(sign, signer)
	}
	return sign
}

func (r *repository) GetUserByIdSignatures(idsignature string) modelsUser.ProfileDB {
	var user modelsUser.ProfileDB
	c := r.db.Collection("users")
	filter := bson.M{"idsignature": idsignature}
	err := c.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		log.Println(err)
	}
	return user
}

func (r *repository) VerifyDoc(hash string) bool {
	conf, _ := config.Init()
	publickey := "0x" + conf.Blockhain.Public
	check, err := r.blockchain.VerifyDoc(&bind.CallOpts{From: common.HexToAddress(publickey)}, hash)
	if err != nil {
		log.Println(err)
	}
	return check
}

func (r *repository) GetTransactions() []models.Transac {
	var transac []models.Transac
	c := r.db.Collection("transactions")
	cursor, err := c.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println(err)
	}
	for cursor.Next(context.Background()) {
		var model models.Transac
		err := cursor.Decode(&model)
		if err != nil {
			log.Println(err)
		}
		transac = append(transac, model)
	}
	return transac
}
