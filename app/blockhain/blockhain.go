package blockhain

import (
	"context"
	"crypto/ecdsa"
	"e-signature/app/config"
	api "e-signature/app/contracts"
	database "e-signature/app/databases"
	"e-signature/modules/v1/utilities/user/models"
	tm "e-signature/pkg/time"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAccountAuth(client *ethclient.Client, privateKeyAddress string) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(privateKeyAddress)
	if err != nil {
		log.Println(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("Invalid key")
	}
	//ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	//defer cancel()
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//fmt.Println(fromAddress)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Println(err)
	}
	//fmt.Println(fromAddress)
	//fmt.Println("nounce= ", nonce)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Println(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	//auth, err = bind.NewTransactor(privateKey)
	if err != nil {
		log.Println(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = big.NewInt(1000000)
	return auth
}
func Connect() *ethclient.Client {
	conf, _ := config.Init()
	//client, err := ethclient.Dial(conf.Blockhain.Host + ":" + conf.Blockhain.Port)
	client, err := ethclient.Dial(conf.Blockhain.Host + conf.Blockhain.Key)
	if err != nil {
		log.Println(err)
	}
	return client
}

func NewContract(client *ethclient.Client, conf config.Conf) common.Address {
	auth := GetAccountAuth(client, conf.Blockhain.Secret_key)
	address, tx, _, err := api.DeployApi(auth, client)
	if err != nil {
		log.Println("Error Deploy API")
		//log.Println(err)
	}

	// fmt.Println(address.Hex())
	// fmt.Println("instance : ", instance)
	fmt.Println("tx : ", tx.Hash().Hex())

	db := database.Init(conf)
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Println(err)
	}
	ctx := context.TODO()
	trans := models.Transac{
		Id:               primitive.NewObjectID(),
		Address:          fmt.Sprintf("%v", address.Hex()),
		Tx_hash:          fmt.Sprintf("%v", tx.Hash().Hex()),
		Nonce:            auth.Nonce.String(),
		Description:      "Aplikasi Berjalan, Membuat Kontrak",
		Date_created:     time.Now().In(location),
		Date_created_wib: tm.TanggalJam(time.Now().In(location)),
	}
	c := db.Collection("transactions")
	_, err = c.InsertOne(ctx, &trans)
	if err != nil {
		log.Println(err)
	}
	return address
}

func Init(conf config.Conf) (*api.Api, *ethclient.Client) {
	client := Connect()
	//Deploy New Contract
	address := NewContract(client, conf)
	//Connect to Existing Contract
	//address := common.HexToAddress(conf.Contract.Smart_contract)
	conn, err := api.NewApi(address, client)
	if err != nil {
		log.Println("Error Connect to Smart Contract")
		//log.Println(err)
	}

	return conn, client
}
