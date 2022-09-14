package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"e-signature/modules/v1/utilities/user/models"
	"e-signature/modules/v1/utilities/user/repository"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	ConnectIPFS() *shell.Shell
	UploadIPFS(path string) (error, string)
	GetFileIPFS(hash string, output string) (string, error)
	Login(input models.LoginInput) (models.ProfileDB, error)
	CreateAccount(user models.User) error
	SaveImage(input models.RegisterUserInput, file *multipart.FileHeader) (string, error)
	CreateKey(key string) []byte
	Encrypt(data []byte, passphrase string) []byte
	Decrypt(data []byte, passphrase string) []byte
	EncryptFile(filename string, passphrase string) error
	DecryptFile(filename string, passphrase string) error
	//SavetoSystem(user models.User) error
	CheckUserExist(idsignature string) (string, error)
	CheckEmailExist(email string) (string, error)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) ConnectIPFS() *shell.Shell {
	var sh *shell.Shell
	sh = shell.NewShell("localhost:5001")
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
	return err, cid
}

func (s *service) GetFileIPFS(hash string, output string) (string, error) {
	sh := s.ConnectIPFS()
	outputName := "./public/images/" + output
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
	if err != nil {
		log.Println(err)
		return user, err
	}

	if user.User_id == 0 {
		log.Println("User not found")
		return user, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Println(err)
		return user, errors.New("password salah")
	}

	return user, nil
}

func (s *service) CreateAccount(user models.User) error {
	//Input Hash Password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		return err
	}
	user.PasswordHash = string(hash)
	//Generate account public key with passphrase user password
	user, err = s.repository.GeneratePublicKey(user)
	if err != nil {
		log.Println(err)
		return err
	}
	user.Publickey = string(s.Encrypt([]byte(user.Publickey), user.Password))
	user.Identity_card = string(s.Encrypt([]byte(user.Identity_card), user.Password))
	//fmt.Println(string(s.Decrypt([]byte(user.Publickey), user.Password)))
	//Save to Database
	err = s.repository.Register(user)
	if err != nil {
		log.Println(err)
		return err
	}
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
	return nil
}

func (s *service) SaveImage(input models.RegisterUserInput, file *multipart.FileHeader) (string, error) {
	path := fmt.Sprintf("./public/images/%s-%s", input.IdSignature, file.Filename)
	//Encrypt file Image with AES and Passphrase password
	err := s.EncryptFile(path, input.Password)
	if err != nil {
		return "", err
	}
	//Upload to Network IPFS
	err, cidr := s.UploadIPFS(path)
	if err != nil {
		return "", err
	}
	return cidr, nil
}

func (s *service) CreateKey(key string) []byte {
	hash := sha256.Sum256([]byte(key))
	return hash[:]
}

func (s *service) Encrypt(data []byte, passphrase string) []byte {
	key := s.CreateKey(passphrase)
	block, _ := aes.NewCipher(key)
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
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
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := cipherText[:nonceSize], cipherText[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func (s *service) EncryptFile(filename string, passphrase string) error {
	b, err := ioutil.ReadFile(filename) //Read the target file
	if err != nil {
		fmt.Printf("Unable to open the input file!\n")
		os.Exit(0)
		return err
	}
	ciphertext := s.Encrypt(b, passphrase)
	err = ioutil.WriteFile(filename, ciphertext, 0644)
	if err != nil {
		fmt.Printf("Unable to create encrypted file!\n")
		os.Exit(0)
		return err
	}
	//fmt.Println(ciphertext)
	return nil
}

func (s *service) DecryptFile(filename string, passphrase string) error {
	z, err := ioutil.ReadFile(filename)
	result := s.Decrypt(z, passphrase)
	//fmt.Printf("Decrypted file was created with file permissions 0777\n")
	err = ioutil.WriteFile(filename, result, 0777)
	if err != nil {
		fmt.Printf("Unable to create decrypted file!\n")
		os.Exit(0)
		return err
	}
	return nil
}

func (s *service) CheckUserExist(idsignature string) (string, error) {
	id, err := s.repository.CheckUserExist(idsignature)
	if err != nil {
		log.Println(err)
		return "", err
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
		return "", err
	}
	if id.Email == "" {
		return "no-exist", nil
	}
	return "exist", nil
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
