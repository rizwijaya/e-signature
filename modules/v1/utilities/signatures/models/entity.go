package models

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
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
	Id               primitive.ObjectID `bson:"_id"`
	Ids              string
	Address          string    `bson:"address"`
	Tx_hash          string    `bson:"tx_hash"`
	Nonce            string    `bson:"nonce"`
	Prices           string    `bson:"prices"`
	Description      string    `bson:"description"`
	Date_created     time.Time `bson:"date_created"`
	Date_created_wib string    `bson:"date_created_wib"`
}

type ListDocument struct {
	Id               primitive.ObjectID `bson:"_id"`
	Address          string             `bson:"address"`
	Hash             string             `bson:"hash"`
	Hash_original    string             `bson:"hash_ori"`
	Judul            string             `bson:"judul"`
	Note             string             `bson:"note"`
	Date_created     time.Time          `bson:"date_created"`
	Date_created_WIB string             `bson:"date_created_wib"`
	Documents        DocumentBlockchain `bson:"documents"`
}

type DocumentBlockchain struct {
	Document_id    string
	Creator        common.Address
	Creator_string string
	Creator_id     string
	Metadata       string
	Hash_ori       string
	Hash           string
	IPFS           string
	State          string
	Mode           string
	Createdtime    string
	Completedtime  string
	Exist          bool
	Signers        Signers
}

type DocumentAllSign struct {
	Document_id   string
	Creator       common.Address
	Creator_id    string
	Metadata      string
	Hash_ori      string
	Hash          string
	IPFS          string
	State         string
	Mode          string
	Createdtime   string
	Completedtime string
	Exist         bool
	Signers       []SignersData
}

type Signers struct {
	Sign_addr     common.Address
	Sign_id       string
	Signers_id    string
	Signers_hash  string
	Signers_state bool
	Sign_time     string
}

type SignersData struct {
	Sign_addr     string `bson:"address"`
	Sign_name     string `bson:"name"`
	Sign_email    string `bson:"email"`
	Sign_id_db    string `bson:"_id"`
	Sign_id       string
	Signers_id    string `bson:"idsignature"`
	Signers_hash  string `bson:"hash"`
	Signers_state bool
	Sign_time     string
}
