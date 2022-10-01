package repository

import (
	"context"
	"e-signature/modules/v1/utilities/signatures/models"
	modelsUser "e-signature/modules/v1/utilities/user/models"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	DefaultSignatures(user modelsUser.User, id string) error
	UpdateMySignatures(signature string, signaturedata string, sign string) error
	GetMySignature(sign string) (models.Signatures, error)
}

type repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) *repository {
	return &repository{db}
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
