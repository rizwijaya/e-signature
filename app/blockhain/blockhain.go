package blockhain

import (
	"context"
	"crypto/ecdsa"
	"e-signature/app/config"
	api "e-signature/app/contracts"
	database "e-signature/app/databases"
	"e-signature/modules/v1/utilities/user/models"
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

func Init(conf config.Conf) (*api.Api, *ethclient.Client) {
	client := Connect()

	//---------------Uncomment For New Deploy Smart Contract----------//
	auth := GetAccountAuth(client, conf.Blockhain.Secret_key)
	address, tx, _, err := api.DeployApi(auth, client)
	if err != nil {
		//log.Println("Error Deploy API")
		log.Println(err)
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
		Id:           primitive.NewObjectID(),
		Address:      fmt.Sprintf("%v", address.Hex()),
		Tx_hash:      fmt.Sprintf("%v", tx.Hash().Hex()),
		Nonce:        auth.Nonce.String(),
		Description:  "Aplikasi Berjalan, Membuat Kontrak",
		Date_created: time.Now().In(location),
	}
	c := db.Collection("transactions")
	_, err = c.InsertOne(ctx, &trans)
	if err != nil {
		log.Println(err)
	}
	//------------End Uncomment For New Deploy Smart Contract----------//
	//address := common.HexToAddress("0x8101c772c3af62bb3096b5dd9dfd9b53cd50652e")
	conn, err := api.NewApi(common.HexToAddress(address.Hex()), client)
	if err != nil {
		log.Println("Error Create New API")
		log.Println(err)
	}

	return conn, client
}
