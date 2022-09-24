package database

import (
	"context"
	"e-signature/app/config"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init(conf config.Conf) *mongo.Database {
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI("mongodb://" + conf.Db.Host + ":" + conf.Db.Port)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println(err)
		return nil
	}
	db := client.Database(conf.Db.Name)
	return db
}
