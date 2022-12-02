package repository

import (
	"context"
	"e-signature/modules/v1/utilities/user/models"
	bl "e-signature/pkg/blockchain"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type Repository interface {
	GeneratePublicKey(user models.User) (models.User, error)
	Register(user models.User) (interface{}, error)
	CheckUserExist(idsignature string) (models.ProfileDB, error)
	CheckEmailExist(email string) (models.ProfileDB, error)
	//TransferBalance(user models.ProfileDB) error
	//GetBalance(user models.ProfileDB, pw string) (string, error)
	GetUserByEmail(email string) (models.User, error)
	GetTotal(db string) int
	GetTotalRequestUser(sign_id string) int
	Logging(logg models.UserLog) error
	GetLogUser(idsignature string) ([]models.UserLog, error)
}

type repository struct {
	db         *mongo.Database
	blockchain bl.Blockchain
}

func NewRepository(db *mongo.Database, blockchain bl.Blockchain) *repository {
	return &repository{db, blockchain}
}

func (r *repository) GeneratePublicKey(user models.User) (models.User, error) {
	return r.blockchain.GeneratePublicKey(user)
}

func (r *repository) Register(user models.User) (interface{}, error) {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	user.Id = primitive.NewObjectID()
	profile := models.ProfileDB{
		Id:            user.Id,
		Idsignature:   user.Idsignature,
		Name:          user.Name,
		Email:         user.Email,
		Phone:         user.Phone,
		Identity_card: user.Identity_card,
		Password:      user.PasswordHash,
		PublicKey:     user.Publickey,
		Role_id:       user.Role,
		Date_created:  time.Now().In(location),
	}
	c := r.db.Collection("users")
	id, err := c.InsertOne(context.Background(), &profile)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return id.InsertedID, err
}

func (r *repository) CheckUserExist(idsignature string) (models.ProfileDB, error) {
	var profile models.ProfileDB
	c := r.db.Collection("users")
	err := c.FindOne(context.Background(), bson.M{"idsignature": idsignature}).Decode(&profile)
	if err != nil {
		return profile, err
	}

	return profile, nil
}

func (r *repository) CheckEmailExist(email string) (models.ProfileDB, error) {
	var profile models.ProfileDB
	c := r.db.Collection("users")
	err := c.FindOne(context.Background(), bson.M{"email": email}).Decode(&profile)
	if err != nil {
		return profile, err
	}

	return profile, nil
}

// func (r *repository) TransferBalance(user models.ProfileDB) error {
// 	user_address, txs, nonce, err := r.blockchain.TransferBalance(user)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}
// 	location, err := time.LoadLocation("Asia/Jakarta")
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}
// 	ctx := context.TODO()
// 	trans := models.Transac{
// 		Id:           primitive.NewObjectID(),
// 		Address:      fmt.Sprintf("%v", user_address),
// 		Tx_hash:      fmt.Sprintf("%v", txs),
// 		Nonce:        nonce,
// 		Description:  "Kirim Ether kepada user +",
// 		Date_created: time.Now().In(location),
// 	}
// 	c := r.db.Collection("transactions")
// 	_, err = c.InsertOne(ctx, &trans)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}
// 	//fmt.Printf("tx has sent to: %s", txs.Hash().Hex())
// 	return nil
// }

// func (r *repository) GetBalance(user models.ProfileDB, pw string) (string, error) {
// 	return r.blockchain.GetBalance(user, pw)
// }

func (r *repository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	filter := map[string]interface{}{"email": email}
	c := r.db.Collection("users")
	err := c.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		log.Println(err)
		return user, err
	}
	return user, nil
}

func (r *repository) GetTotal(db string) int {
	c := r.db.Collection(db)
	total, err := c.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		log.Println(err)
		return 0
	}
	return int(total)
}

func (r *repository) GetTotalRequestUser(sign_id string) int {
	c := r.db.Collection("users")
	total, err := c.CountDocuments(context.Background(), bson.M{"idsignature": sign_id})
	if err != nil {
		log.Println(err)
		return 0
	}
	return int(total)
}

func (r *repository) Logging(logg models.UserLog) error {
	logg.Id = primitive.NewObjectID()
	c := r.db.Collection("user_log")
	_, err := c.InsertOne(context.Background(), logg)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *repository) GetLogUser(idsignature string) ([]models.UserLog, error) {
	var logg []models.UserLog
	c := r.db.Collection("user_log")
	cur, err := c.Find(context.Background(), bson.M{"idsignature": idsignature})
	if err != nil {
		log.Println(err)
		return logg, err
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var result models.UserLog
		err := cur.Decode(&result)
		if err != nil {
			log.Println(err)
			return logg, err
		}
		logg = append(logg, result)
	}
	if err := cur.Err(); err != nil {
		log.Println(err)
		return logg, err
	}
	return logg, nil
}
