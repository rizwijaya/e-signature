package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Profile_id     int
	Idsignature    string
	Name           string
	Password       string
	PasswordHash   string
	ImageIPFS      string
	Role           int
	Publickey      string
	Identity_card  string
	Email          string
	Phone          string
	Dateregistered string
}

type ProfileDB struct {
	Id            primitive.ObjectID `bson:"_user_id"`
	Idsignature   string             `bson:"idsignature"`
	Name          string             `bson:"name"`
	Email         string             `bson:"email"`
	Phone         string             `bson:"phone"`
	Identity_card string             `bson:"identity_card"`
	Password      string             `bson:"password"`
	PublicKey     string             `bson:"public_key"`
	Role_id       int                `bson:"role"`
	Date_created  time.Time          `bson:"date_created"`
}

type Transac struct {
	Id           primitive.ObjectID `bson:"_id"`
	Address      string             `bson:"address"`
	Tx_hash      string             `bson:"tx_hash"`
	Nonce        string             `bson:"nonce"`
	Description  string             `bson:"description"`
	Date_created time.Time          `bson:"date_created"`
}
