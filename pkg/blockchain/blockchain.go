package blockchain

import (
	"context"
	blockhainAuth "e-signature/app/blockhain"
	"e-signature/app/config"
	api "e-signature/app/contracts"
	modelsSign "e-signature/modules/v1/utilities/signatures/models"
	modelsUser "e-signature/modules/v1/utilities/user/models"
	"errors"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"regexp"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Blockchain interface {
	GetPrivateKey(user modelsUser.User) (string, error)
	GeneratePublicKey(user modelsUser.User) (modelsUser.User, error)
	TransferBalance(user modelsUser.ProfileDB) (string, string, string, error)
	GetBalance(user modelsUser.ProfileDB, pw string) (string, error)
	AddToBlockhain(input modelsSign.SignDocuments, times *big.Int) (*types.Transaction, *bind.TransactOpts, error)
	DocumentSigned(sign modelsSign.SignDocs, timeSign *big.Int) (*types.Transaction, *bind.TransactOpts, error)
	GetDocument(hash string, publickey string) modelsSign.DocumentBlockchain
	GetSigners(hash string, publickey string) modelsSign.Signers
	GetHashOriginal(hash string, publickey string) string
	VerifyDoc(hash string) bool
	GetListSign(hash string) []common.Address
}

type blockchain struct {
	contracts *api.Api
	client    *ethclient.Client
}

func NewBlockchain(contracts *api.Api, client *ethclient.Client) *blockchain {
	return &blockchain{contracts, client}
}

var filesPrivate []string

func (b *blockchain) SearchFile(path string, info os.FileInfo, err error) error {
	var user modelsUser.User
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

func (b *blockchain) GetPrivateKey(user modelsUser.User) (string, error) {
	err := filepath.Walk("./app/account", b.SearchFile)
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

func (b *blockchain) GeneratePublicKey(user modelsUser.User) (modelsUser.User, error) {
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

func (b *blockchain) TransferBalance(user modelsUser.ProfileDB) (string, string, string, error) {
	conf, _ := config.Init()
	client, err := ethclient.Dial(conf.Blockhain.Host + conf.Blockhain.Key)
	if err != nil {
		return "", "", "", err
	}
	defer client.Close()

	system_address := common.HexToAddress(conf.Blockhain.Public)
	user_address := common.HexToAddress(user.PublicKey)

	value := big.NewInt(10000)
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", "", "", err
	}

	nonce, err := client.PendingNonceAt(context.Background(), system_address)
	if err != nil {
		return "", "", "", err
	}

	var data []byte
	txs := types.NewTransaction(nonce, user_address, value, gasLimit, gasPrice, data)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", "", "", err
	}

	bs, err := ioutil.ReadFile(conf.Blockhain.Secret_base)
	if err != nil {
		return "", "", "", err
	}

	key, err := keystore.DecryptKey(bs, "rizwijaya58")
	if err != nil {
		return "", "", "", err
	}
	txs, err = types.SignTx(txs, types.NewEIP155Signer(chainID), key.PrivateKey)

	if err != nil {
		return "", "", "", err
	}

	err = client.SendTransaction(context.Background(), txs)

	if err != nil {
		return "", "", "", err
	}
	return user_address.Hex(), txs.Hash().Hex(), big.NewInt(int64(nonce)).String(), nil
}

func (b *blockchain) GetBalance(user modelsUser.ProfileDB, pw string) (string, error) {
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

func (b *blockchain) AddToBlockhain(input modelsSign.SignDocuments, times *big.Int) (*types.Transaction, *bind.TransactOpts, error) {
	conf, _ := config.Init()
	mode := new(big.Int)
	mode.SetString(input.Mode, 10)
	auth := blockhainAuth.GetAccountAuth(blockhainAuth.Connect(), conf.Blockhain.Secret_key)
	document, err := b.contracts.Create(auth, input.Hash_original, common.HexToAddress(input.Creator), input.Creator_id, input.Name, input.Hash, input.IPFS, big.NewInt(1), mode, times, input.Address, input.IdSignature)
	if err != nil {
		log.Println(err)
	}
	return document, auth, err
}

func (b *blockchain) DocumentSigned(sign modelsSign.SignDocs, timeSign *big.Int) (*types.Transaction, *bind.TransactOpts, error) {
	conf, _ := config.Init()
	auth := blockhainAuth.GetAccountAuth(blockhainAuth.Connect(), conf.Blockhain.Secret_key)
	signDocs, err := b.contracts.SignDoc(auth, sign.Hash_original, common.HexToAddress(sign.Creator), sign.Hash, sign.IPFS, timeSign)
	return signDocs, auth, err
}

func (b *blockchain) GetDocument(hash string, publickey string) modelsSign.DocumentBlockchain {
	var doc modelsSign.DocumentBlockchain
	var err error
	var document_id, state, Createdtime, Completedtime, Mode *big.Int
	document_id, doc.Creator, doc.Creator_id, doc.Metadata, doc.Hash_ori, doc.Hash, doc.IPFS, state, Mode, Createdtime, Completedtime, doc.Exist, err = b.contracts.GetDoc(&bind.CallOpts{From: common.HexToAddress(publickey)}, hash)
	if err != nil {
		log.Println("Error Get Document")
	}
	doc.Document_id = document_id.String()
	doc.State = state.String()
	doc.Createdtime = Createdtime.String()
	doc.Completedtime = Completedtime.String()
	doc.Mode = Mode.String()
	doc.Creator_string = doc.Creator.String()

	return doc
}

func (b *blockchain) GetSigners(hash string, publickey string) modelsSign.Signers {
	var signers modelsSign.Signers
	var err error
	var sign_id, sign_time *big.Int
	sign_id, signers.Signers_id, signers.Signers_hash, signers.Signers_state, sign_time, err = b.contracts.GetSign(&bind.CallOpts{From: common.HexToAddress(publickey)}, hash, common.HexToAddress(publickey))
	if err != nil {
		log.Println("Error Get Signers")
	}
	signers.Sign_id = sign_id.String()
	signers.Sign_time = sign_time.String()
	return signers
}

func (b *blockchain) GetHashOriginal(hash string, publickey string) string {
	hash_ori, err := b.contracts.GetDocSigned(&bind.CallOpts{From: common.HexToAddress(publickey)}, hash)
	if err != nil {
		log.Println(err)
		log.Println("Error Get Hash Original")
	}
	return hash_ori
}

func (b *blockchain) VerifyDoc(hash string) bool {
	conf, _ := config.Init()
	publickey := "0x" + conf.Blockhain.Public
	check, err := b.contracts.VerifyDoc(&bind.CallOpts{From: common.HexToAddress(publickey)}, hash)
	if err != nil {
		log.Println(err)
	}
	return check
}

func (b *blockchain) GetListSign(hash string) []common.Address {
	conf, _ := config.Init()
	publickey := "0x" + conf.Blockhain.Public
	list, err := b.contracts.GetListSign(&bind.CallOpts{From: common.HexToAddress(publickey)}, hash)
	if err != nil {
		log.Println(err)
	}
	return list
}
