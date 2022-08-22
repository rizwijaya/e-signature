package repository

import (
	"context"
	"e-signature/app/blockhain"
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

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

type Repository interface {
	SearchFile(path string, info os.FileInfo, err error) error
	GetPrivateKey(user models.User) (string, error)
	GeneratePublicKey(user models.User) (models.User, error)
	//SavetoSystem(auth *bind.TransactOpts, user models.User) error
	Register(user models.User) error
	SavetoProfile(user models.User, key string) error
	CheckUserExist(idsignature string) (models.ProfileDB, error)
	TransferBalance(user models.User) (string, error)
}

type repository struct {
	db         *gorm.DB
	blockchain *api.Api
	client     *ethclient.Client
}

func NewRepository(db *gorm.DB, blockchain *api.Api, client *ethclient.Client) *repository {
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
		log.Fatal(err)
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
		log.Fatal(e)
		return "", e
	}
	key, e := keystore.DecryptKey(keyjson, user.Password)
	if e != nil {
		log.Fatal(e)
		return "", e
	}
	//fmt.Println(key.PrivateKey)
	e = crypto.SaveECDSA(outPath, key.PrivateKey)
	if e != nil {
		log.Fatal(e)
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
		log.Fatal(err)
		return user, err
	}
	user.Publickey = account.Address.Hex()
	// file, _ := r.GetPrivateKey()

	// jsonBytes, err := ioutil.ReadFile(file)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return user, err
	// }

	// account, err = ks.Import(jsonBytes, password, password)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return user, err
	// }

	// user.Publickey = account.Address.Hex() //save to struct
	return user, nil
}

// func (r *repository) SavetoSystem(auth *bind.TransactOpts, user models.User) error {
// 	_, err := r.blockchain.AddProfilefirst(auth, user.Idsignature, user.Password, user.Publickey, user.Role)
// 	if err != nil {
// 		log.Fatal(err)
// 		return err
// 	}
// 	return nil
// }
func (r *repository) Register(user models.User) error {
	var profile models.ProfileDB
	profile.Idsignature = user.Idsignature
	profile.Password = user.PasswordHash
	profile.Name = user.Name
	profile.Email = user.Email
	profile.Phone = user.Phone
	profile.Identity_card = user.Identity_card
	profile.PublicKey = user.Publickey
	profile.Role_id = user.Role
	err := r.db.Table("users").Create(&profile).Error
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (r *repository) SavetoProfile(user models.User, key string) error {
	//auth := blockhain.GetAccountAuth(blockhain.Connect(), key)
	auth := blockhain.GetAccountAuth(blockhain.Connect(), key)
	rs, err := r.blockchain.AddProfile(auth, user.Name, user.ImageIPFS, user.Email, user.Phone, user.Dateregistered)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println(rs)
	return nil
}

func (r *repository) CheckUserExist(idsignature string) (models.ProfileDB, error) {
	var profile models.ProfileDB
	err := r.db.Table("users").Where("idsignature = ?", idsignature).Find(&profile).Error
	if err != nil {
		return profile, err
	}

	return profile, nil
}

func (r *repository) TransferBalance(user models.User) (string, error) {
	conf, _ := config.Init()
	var data []byte
	err := filepath.Walk("./app/account", r.SearchFile)
	if err != nil {
		log.Fatal(err)
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
	keyjson, e := ioutil.ReadFile(filePrivate)
	if e != nil {
		log.Fatal(e)
		return "", e
	}
	key, e := keystore.DecryptKey(keyjson, user.Password)
	if e != nil {
		log.Fatal(e)
		return "", e
	}
	//Transfer Balance
	from := common.HexToAddress(conf.Blockhain.Secret_key)
	to := common.HexToAddress(user.Publickey)

	value := big.NewInt(1000)
	gasLimit := uint64(21000)
	gasPrice, err := r.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	nonce, err := r.client.PendingNonceAt(context.Background(), from)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	txs := types.NewTransaction(nonce, to, value, gasLimit, gasPrice, data)
	chainID, err := r.client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	txs, err = types.SignTx(txs, types.NewEIP155Signer(chainID), key.PrivateKey)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	err = r.client.SendTransaction(context.Background(), txs)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return txs.Hash().Hex(), err
}
