package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MySignatures struct {
	Id                 string
	Name               string
	User_id            string
	Signature          string
	Signature_id       string
	Signature_data     string
	Signature_data_id  string
	Latin              string
	Latin_id           string
	Latin_data         string
	Latin_data_id      string
	Signature_selected string
	Date_update        string
	Date_created       string
}

type SignDocs struct {
	Hash_original string
	Creator       string
	Hash          string
	IPFS          string
}

type Transac struct {
	Id           primitive.ObjectID `bson:"_id"`
	Address      string             `bson:"address"`
	Tx_hash      string             `bson:"tx_hash"`
	Nonce        string             `bson:"nonce"`
	Prices       string             `bson:"prices"`
	Description  string             `bson:"description"`
	Date_created time.Time          `bson:"date_created"`
}
