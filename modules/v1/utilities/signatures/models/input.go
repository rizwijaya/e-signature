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
	Id                 primitive.ObjectID `bson:"_id"`
	User               string             `bson:"user"`
	Signature          string             `bson:"signature"`
	Signature_data     string             `bson:"signature_data"`
	Latin              string             `bson:"latin"`
	Latin_data         string             `bson:"latin_data"`
	Signature_selected string             `bson:"signature_selected"`
	Date_update        time.Time          `bson:"date_update"`
	Date_created       time.Time          `bson:"date_created"`
}

type SignDocuments struct {
	//File       string  `json:"file" form:"file" binding:"required"`
	Name       string
	SignPage   float64 `json:"signPage" form:"signPage" binding:"required"`
	X_coord    float64 `json:"signX" form:"signX" binding:"required"`
	Y_coord    float64 `json:"signY" form:"signY" binding:"required"`
	Height     float64 `json:"signH" form:"signH" binding:"required"`
	Width      float64 `json:"signW" form:"signW" binding:"required"`
	Invite_sts bool    `json:"invite_status" form:"invite_status"`
	Email      string  `json:"email" form:"email"`
}
