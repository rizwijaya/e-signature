package repository

import (
	"context"
	"e-signature/modules/v1/utilities/signatures/models"
	modelsUser "e-signature/modules/v1/utilities/user/models"
	bl "e-signature/pkg/blockchain"
	tm "e-signature/pkg/time"
	"fmt"
	"log"
	"math/big"
	"time"

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
	CheckSignature(hash string, publickey string) bool
}

type repository struct {
	db         *mongo.Database
	blockchain bl.Blockchain
}

func NewRepository(db *mongo.Database, blockchain bl.Blockchain) *repository {
	return &repository{db, blockchain}
}

func (r *repository) LogTransactions(address string, tx_hash string, nonce string, desc string, prices string) error {
	location, _ := time.LoadLocation("Asia/Jakarta")
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
	_, err := c.InsertOne(ctx, &trans)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *repository) DefaultSignatures(user modelsUser.User, id string) error {
	location, _ := time.LoadLocation("Asia/Jakarta")
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
	_, err := c.InsertOne(context.Background(), &signatures)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *repository) UpdateMySignatures(signature string, signaturedata string, sign string) error {
	location, _ := time.LoadLocation("Asia/Jakarta")
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
	_, err := c.UpdateOne(context.Background(), filter, update)
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
	location, _ := time.LoadLocation("Asia/Jakarta")
	filter := map[string]interface{}{"user": sign}
	update := map[string]interface{}{
		"$set": map[string]interface{}{
			"signature_selected": sign_type,
			"date_update":        time.Now().In(location),
		},
	}

	c := r.db.Collection("signatures")
	_, err := c.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *repository) AddToBlockhain(input models.SignDocuments, times *big.Int) error {
	document, auth, err := r.blockchain.AddToBlockhain(input, times)
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
	signDocs, auth, err := r.blockchain.DocumentSigned(sign, timeSign)
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
	return r.blockchain.GetDocument(hash, publickey)
}

func (r *repository) GetSigners(hash string, publickey string) models.Signers {
	return r.blockchain.GetSigners(hash, publickey)
}

func (r *repository) GetHashOriginal(hash string, publickey string) string {
	return r.blockchain.GetHashOriginal(hash, publickey)
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
	return r.blockchain.VerifyDoc(hash)
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

func (r *repository) CheckSignature(hash string, publickey string) bool {
	var sign struct {
		Hash string `bson:"hash"`
	}
	c := r.db.Collection("signedDocuments")
	filter := bson.M{"hash_ori": hash, "address": publickey}
	err := c.FindOne(context.Background(), filter).Decode(&sign)
	if err != nil {
		log.Println(err)
	}
	if sign.Hash != "" {
		return true
	} else {
		return false
	}
}
