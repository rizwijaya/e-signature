package repository

import (
	"context"
	"e-signature/app/config"
	api "e-signature/app/contracts"
	"e-signature/modules/v1/utilities/user/models"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type Repository interface {
	SearchFile(path string, info os.FileInfo, err error) error
	GetPrivateKey(user models.User) (string, error)
	GeneratePublicKey(user models.User) (models.User, error)
	Register(user models.User) (interface{}, error)
	CheckUserExist(idsignature string) (models.ProfileDB, error)
	CheckEmailExist(email string) (models.ProfileDB, error)
	TransferBalance(user models.ProfileDB) error
	GetBalance(user models.ProfileDB, pw string) (string, error)
	GetUserByEmail(email string) (models.User, error)
	GetTotal(db string) int
	GetTotalRequestUser(sign_id string) int
	Logging(logg models.UserLog) error
	GetLogUser(idsignature string) ([]models.UserLog, error)
}

type repository struct {
	db         *mongo.Database
	blockchain *api.Api
	client     *ethclient.Client
}

func NewRepository(db *mongo.Database, blockchain *api.Api, client *ethclient.Client) *repository {
	return &repository{db, blockchain, client}
}

var filesPrivate []string

func (r *repository) SearchFile(path string, info os.FileInfo, err error) error {
	var user models.User
	pub := user.Publickey
	if err != nil {
		log.Println(err)
		return err
	}

	reg, err2 := regexp.Compile(pub + "$")
	if err2 != nil {
		log.Println(err2)
		return err
	}

	if reg.MatchString(info.Name()) {
		filesPrivate = append(filesPrivate, path)
	}

	return nil
}

func (r *repository) GetPrivateKey(user models.User) (string, error) {
	err := filepath.Walk("./app/account", r.SearchFile)
	if err != nil {
		log.Println(err)
		return "", err
	}
	var filePrivate string
	for _, file := range filesPrivate {
		filePrivate = file
	}
	if len(filePrivate) == 0 {
		return "", errors.New("file not found")
	}
	//Dapatkan Private Key dari File Keystorenya
	outPath := "./app/tmp/hash.hex"
	keyjson, e := ioutil.ReadFile(filePrivate)
	if e != nil {
		log.Println(e)
		return "", e
	}
	key, e := keystore.DecryptKey(keyjson, user.Password)
	if e != nil {
		log.Println(e)
		return "", e
	}
	//fmt.Println(key.PrivateKey)
	e = crypto.SaveECDSA(outPath, key.PrivateKey)
	if e != nil {
		log.Println(e)
		return "", e
	}
	z, _ := ioutil.ReadFile(outPath)
	//fmt.Println(string(z))
	os.Remove(outPath)

	return string(z), err
}

func (r *repository) GeneratePublicKey(user models.User) (models.User, error) {
	ks := keystore.NewKeyStore("./app/account", keystore.StandardScryptN, keystore.StandardScryptP)
	password := user.Password

	account, err := ks.NewAccount(password)
	if err != nil {
		log.Println(err)
		return user, err
	}
	user.Publickey = account.Address.Hex()
	// file, _ := r.GetPrivateKey()

	// jsonBytes, err := ioutil.ReadFile(file)
	// if err != nil {
	// 	log.Println(err)
	// 	return user, err
	// }

	// account, err = ks.Import(jsonBytes, password, password)
	// if err != nil {
	// 	log.Println(err)
	// 	return user, err
	// }

	// user.Publickey = account.Address.Hex() //save to struct
	return user, nil
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

func (r *repository) TransferBalance(user models.ProfileDB) error {
	conf, _ := config.Init()
	client, err := ethclient.Dial(conf.Blockhain.Host + conf.Blockhain.Key)
	if err != nil {
		return err
	}
	defer client.Close()

	system_address := common.HexToAddress(conf.Blockhain.Public)
	user_address := common.HexToAddress(user.PublicKey)

	value := big.NewInt(10000)
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	nonce, err := client.PendingNonceAt(context.Background(), system_address)
	if err != nil {
		return err
	}

	var data []byte
	txs := types.NewTransaction(nonce, user_address, value, gasLimit, gasPrice, data)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return err
	}

	b, err := ioutil.ReadFile(conf.Blockhain.Secret_base)
	if err != nil {
		return err
	}

	key, err := keystore.DecryptKey(b, "rizwijaya58")
	if err != nil {
		return err
	}
	txs, err = types.SignTx(txs, types.NewEIP155Signer(chainID), key.PrivateKey)

	if err != nil {
		return err
	}

	err = client.SendTransaction(context.Background(), txs)

	if err != nil {
		return err
	}
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Println(err)
		return err
	}
	ctx := context.TODO()
	trans := models.Transac{
		Id:           primitive.NewObjectID(),
		Address:      fmt.Sprintf("%v", user_address.Hex()),
		Tx_hash:      fmt.Sprintf("%v", txs.Hash().Hex()),
		Nonce:        big.NewInt(int64(nonce)).String(),
		Description:  "Kirim Ether kepada user +",
		Date_created: time.Now().In(location),
	}
	c := r.db.Collection("transactions")
	_, err = c.InsertOne(ctx, &trans)
	if err != nil {
		log.Println(err)
		return err
	}
	//fmt.Printf("tx has sent to: %s", txs.Hash().Hex())
	return nil
}

func (r *repository) GetBalance(user models.ProfileDB, pw string) (string, error) {
	conf, _ := config.Init()
	client, err := ethclient.Dial(conf.Blockhain.Host + conf.Blockhain.Key)
	if err != nil {
		return "", err
	}
	defer client.Close()
	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(user.PublicKey), nil)
	if err != nil {
		return "", err
	}
	return balance.String(), nil
}

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
