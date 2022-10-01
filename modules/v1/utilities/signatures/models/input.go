package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddSignature struct {
	Id        string `json:"unique" binding:"required"`
	Signature string `json:"signature" binding:"required"`
}

type Signatures struct {
	Id             primitive.ObjectID `bson:"_id"`
	User           string             `bson:"user"`
	Signature      string             `bson:"signature"`
	Signature_data string             `bson:"signature_data"`
	Latin          string             `bson:"latin"`
	Latin_data     string             `bson:"latin_data"`
	Date_update    time.Time          `bson:"date_update"`
	Date_created   time.Time          `bson:"date_created"`
}
