package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"e-signature/app/config"
	"e-signature/modules/v1/utilities/user/models"
	"e-signature/modules/v1/utilities/user/repository"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	shell "github.com/ipfs/go-ipfs-api"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	ConnectIPFS() *shell.Shell
	UploadIPFS(path string) (error, string)
	GetFileIPFS(hash string, output string, directory string) (string, error)
	Login(input models.LoginInput) (models.ProfileDB, error)
	CreateAccount(user models.User) (string, error)
	//SaveImage(input models.RegisterUserInput, file *multipart.FileHeader) (string, error)
	CreateKey(key string) []byte
	Encrypt(data []byte, passphrase string) []byte
	Decrypt(data []byte, passphrase string) []byte
	EncryptFile(filename string, passphrase string) error
	DecryptFile(filename string, passphrase string) error
	//SavetoSystem(user models.User) error
	CheckUserExist(idsignature string) (string, error)
	CheckEmailExist(email string) (string, error)
	GetBalance(user models.ProfileDB, pw string) (string, error)
	TransferBalance(user models.ProfileDB) error
	GetPublicKey(email []string) ([]common.Address, []string)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) ConnectIPFS() *shell.Shell {
	var sh *shell.Shell
	conf, _ := config.Init()
	sh = shell.NewShell(conf.IPFS.Host + ":" + conf.IPFS.Port)
	return sh
}

func (s *service) UploadIPFS(path string) (error, string) {
	//fmt.Printf("Adding %s on IPFS \n", filename)
	sh := s.ConnectIPFS()
	f, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return err, ""
	}
	var r = f
	cid, err := sh.Add(r)
	if err != nil {
		log.Println(err)
		return err, ""
	}
	defer f.Close()

	err = os.Remove(path)
	if err != nil {
		//log.Println("The file could not be removed")
		log.Println(err)
		return err, ""
	}
	return nil, cid
}

func (s *service) GetFileIPFS(hash string, output string, directory string) (string, error) {
	sh := s.ConnectIPFS()
	outputName := directory + output
	err := sh.Get(hash, outputName)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return outputName, nil
}

func (s *service) Login(input models.LoginInput) (models.ProfileDB, error) {
	idsignature := input.IdSignature
	password := input.Password

	user, err := s.repository.CheckUserExist(idsignature)

	if user.Idsignature == "" {
		log.Println("User not found")
		return user, errors.New("user not found")
	}
	if err != nil {
		log.Println(err)
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Println(err)
		return user, errors.New("password salah")
	}

	return user, nil
}

func (s *service) CreateAccount(user models.User) (string, error) {
	conf, _ := config.Init()
	//Input Hash Password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		return "", err
	}
	user.PasswordHash = string(hash)
	//Generate account public key with passphrase user password
	user, err = s.repository.GeneratePublicKey(user)
	if err != nil {
		log.Println(err)
		return "", err
	}
	user.Publickey = string(s.Encrypt([]byte(user.Publickey), conf.App.Secret_key))
	//user.Identity_card = string(s.Encrypt([]byte(user.Identity_card), user.Password))
	//fmt.Println(string(s.Decrypt([]byte(user.Publickey), user.Password)))
	//Save to Database
	id, err := s.repository.Register(user)
	if err != nil {
		log.Println(err)
		return "", err
	}
	idn := id.(primitive.ObjectID).Hex()
	//Get Private Key user
	// key, err := s.repository.GetPrivateKey(user)
	// if err != nil {
	// 	log.Println(err)
	// 	return err
	// }
	//Isi gas dan ether pada account agar dapat melakukan transaksi
	// tx, err := s.repository.TransferBalance(user)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(tx)
	//Save to Network Blockchain account user
	// fmt.Println(user) // fmt.Println(key)
	// err = s.repository.SavetoProfile(user, key)
	// if err != nil {
	// 	return err
	// }
	//fmt.Println(user)
	return idn, nil
}

// func (s *service) SaveImage(input models.RegisterUserInput, file *multipart.FileHeader) (string, error) {
// 	path := fmt.Sprintf("./public/images/%s-%s", input.IdSignature, file.Filename)
// 	//Encrypt file Image with AES and Passphrase password
// 	err := s.EncryptFile(path, input.Password)
// 	if err != nil {
// 		return "", err
// 	}
// 	//Upload to Network IPFS
// 	err, cidr := s.UploadIPFS(path)
// 	if err != nil {
// 		return "", err
// 	}
// 	return cidr, nil
// }

func (s *service) CreateKey(key string) []byte {
	hash := sha256.Sum256([]byte(key))
	return hash[:]
}

func (s *service) Encrypt(data []byte, passphrase string) []byte {
	key := s.CreateKey(passphrase)
	block, _ := aes.NewCipher(key)
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Println(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Println(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)

	base64Cipher := make([]byte, base64.RawStdEncoding.EncodedLen(len(ciphertext)))
	base64.RawStdEncoding.Encode(base64Cipher, ciphertext)
	//return string(base64Cipher)
	return base64Cipher
}

func (s *service) Decrypt(data []byte, passphrase string) []byte {
	cipherText := make([]byte, base64.RawStdEncoding.DecodedLen(len(data)))
	_, err := base64.RawStdEncoding.Decode(cipherText, data)
	if err != nil {
		return nil
	}
	key := s.CreateKey(passphrase)
	//key := "791f13d3e2552bcf31c4f8d0e5d6a1ed"
	//fmt.Println("key", key)
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Println(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := cipherText[:nonceSize], cipherText[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Println(err.Error())
	}
	return plaintext
}

func (s *service) EncryptFile(filename string, passphrase string) error {
	b, err := ioutil.ReadFile(filename) //Read the target file
	if err != nil {
		fmt.Printf("Unable to open the input file!\n")
		return err
	}
	ciphertext := s.Encrypt(b, passphrase)
	err = ioutil.WriteFile(filename, ciphertext, 0644)
	if err != nil {
		fmt.Printf("Unable to create encrypted file!\n")
		return err
	}
	//fmt.Println(ciphertext)
	return nil
}

func (s *service) DecryptFile(filename string, passphrase string) error {
	z, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Unable to open the input file!\n")
		return err
	}
	result := s.Decrypt(z, passphrase)
	//fmt.Printf("Decrypted file was created with file permissions 0777\n")
	err = ioutil.WriteFile(filename, result, 0777)
	if err != nil {
		fmt.Printf("Unable to create decrypted file!\n")
		return err
	}
	return nil
}

func (s *service) CheckUserExist(idsignature string) (string, error) {
	id, err := s.repository.CheckUserExist(idsignature)
	if err != nil {
		log.Println(err)
		return "no-exist", err
	}
	if id.Email == "" {
		return "no-exist", nil
	}
	return "exist", nil
}

func (s *service) CheckEmailExist(email string) (string, error) {
	id, err := s.repository.CheckEmailExist(email)
	if err != nil {
		log.Println(err)
		return "no-exist", err
	}
	if id.Email == "" {
		return "no-exist", nil
	}
	return "exist", nil
}

func (s *service) GetBalance(user models.ProfileDB, pw string) (string, error) {
	balance, err := s.repository.GetBalance(user, pw)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return balance, err
}

func (s *service) TransferBalance(user models.ProfileDB) error {
	err := s.repository.TransferBalance(user)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (s *service) GetPublicKey(email []string) ([]common.Address, []string) {
	var idSignature []string
	var addr []common.Address

	conf, _ := config.Init()
	for _, v := range email {
		user, err := s.repository.GetUserByEmail(v)
		if err != nil {
			log.Println(err)
			continue
		}
		addr = append(addr, common.HexToAddress(string(s.Decrypt([]byte(user.Publickey), conf.App.Secret_key))))
		idSignature = append(idSignature, user.Idsignature)
	}

	return addr, idSignature
}

// func (s *service) SavetoSystem(user models.User) error {
// 	conf, _ := config.Init()
// 	//System Authentication
// 	auth := blockhain.GetAccountAuth(blockhain.Connect(), conf.Blockhain.Secret_key)
// 	err := s.repository.SavetoSystem(auth, user)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}
// 	return nil
// }
