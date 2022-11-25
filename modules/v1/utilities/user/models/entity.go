package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id           primitive.ObjectID `bson:"_id"`
	Idsignature  string             `bson:"idsignature"`
	Name         string             `bson:"name"`
	Password     string             `bson:"password"`
	PasswordHash string
	//ImageIPFS      string
	Role           int    `bson:"role"`
	Publickey      string `bson:"public_key"`
	Identity_card  string `bson:"identity_card"`
	Email          string `bson:"email"`
	Phone          string `bson:"phone"`
	Dateregistered string
}

type ProfileDB struct {
	Id            primitive.ObjectID `bson:"_id"`
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

type CardDashboard struct {
	TotalUser        int
	TotalTx          int
	TotalRequest     int
	TotalRequestUser int
}

type UserLog struct {
	Id              primitive.ObjectID `bson:"_id"`
	Idsignature     string             `bson:"idsignature"`
	User_agent      string             `bson:"user_agent"`
	Ip_address      string             `bson:"ip_address"`
	Action          string             `bson:"action"`
	Date_access     time.Time          `bson:"date_accessed"`
	Date_access_wib string
}
