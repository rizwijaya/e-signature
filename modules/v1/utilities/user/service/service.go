package service

import (
	"e-signature/app/config"
	"e-signature/modules/v1/utilities/user/models"
	"e-signature/modules/v1/utilities/user/repository"
	pw "e-signature/pkg/crypto"
	"e-signature/pkg/ipfs"
	tm "e-signature/pkg/time"
	"errors"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	UploadIPFS(path string) (string, error)
	GetFileIPFS(hash string, output string, directory string) (string, error)
	Login(input models.LoginInput) (models.ProfileDB, error)
	CreateAccount(user models.User) (string, error)
	//SaveImage(input models.RegisterUserInput, file *multipart.FileHeader) (string, error)
	Encrypt(data []byte, passphrase string) []byte
	Decrypt(data []byte, passphrase string) []byte
	EncryptFile(filename string, passphrase string) error
	DecryptFile(filename string, passphrase string) error
	CheckUserExist(idsignature string) (string, error)
	CheckEmailExist(email string) (string, error)
	GetBalance(user models.ProfileDB, pw string) (string, error)
	TransferBalance(user models.ProfileDB) error
	GetPublicKey(email []string) ([]common.Address, []string)
	GetCardDashboard(sign_id string) models.CardDashboard
	Logging(action string, idsignature string, ip string, user_agent string) error
	GetLogUser(idsignature string) ([]models.UserLog, error)
	GetUserByEmail(email string) (models.User, error)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) UploadIPFS(path string) (string, error) {
	return ipfs.UploadIPFS(path)
}

func (s *service) GetFileIPFS(hash string, output string, directory string) (string, error) {
	return ipfs.GetFileIPFS(hash, output, directory)
}

func (s *service) Login(input models.LoginInput) (models.ProfileDB, error) {
	idsignature := input.IdSignature
	password := input.Password

	user, err := s.repository.CheckUserExist(idsignature)

	if user.Idsignature == "" || err != nil {
		log.Println(err)
		return user, errors.New("user not found")
	}

	err = pw.Compare(user.Password, password)
	if err != nil {
		log.Println(err)
		return user, errors.New("password salah")
	}

	return user, nil
}

func (s *service) CreateAccount(user models.User) (string, error) {
	conf, _ := config.Init()
	var err error
	//Input Hash Password
	user.PasswordHash, err = pw.GenerateHash(user.Password)
	if err != nil {
		log.Println(err)
		return "", err
	}
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
// 	cidr, err := s.UploadIPFS(path)
// 	if err != nil {
// 		return "", err
// 	}
// 	return cidr, nil
// }

func (s *service) Encrypt(data []byte, passphrase string) []byte {
	return pw.Encrypt(data, passphrase)
}

func (s *service) Decrypt(data []byte, passphrase string) []byte {
	return pw.Decrypt(data, passphrase)
}

func (s *service) EncryptFile(filename string, passphrase string) error {
	return pw.EncryptFile(filename, passphrase)
}

func (s *service) DecryptFile(filename string, passphrase string) error {
	return pw.DecryptFile(filename, passphrase)
}

func (s *service) CheckUserExist(idsignature string) (string, error) {
	id, err := s.repository.CheckUserExist(idsignature)
	if err != nil || id.Email == "" {
		log.Println(err)
		return "no-exist", err
	}
	return "exist", nil
}

func (s *service) CheckEmailExist(email string) (string, error) {
	id, err := s.repository.CheckEmailExist(email)
	if err != nil || id.Email == "" {
		log.Println(err)
		return "no-exist", err
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

func (s *service) GetCardDashboard(sign_id string) models.CardDashboard {
	var card models.CardDashboard
	card.TotalRequest = s.repository.GetTotal("signedDocuments")
	card.TotalUser = s.repository.GetTotal("users")
	card.TotalTx = s.repository.GetTotal("transactions")
	card.TotalRequestUser = s.repository.GetTotalRequestUser(sign_id)
	return card
}

func (s *service) Logging(action string, idsignature string, ip string, user_agent string) error {
	var logg models.UserLog
	logg.Action = action
	logg.Idsignature = idsignature
	logg.Ip_address = ip
	logg.User_agent = user_agent
	location, _ := time.LoadLocation("Asia/Jakarta")
	logg.Date_access = time.Now().In(location)
	logg.Date_access_wib = tm.TanggalJam(time.Now().In(location))
	err := s.repository.Logging(logg)
	return err
}

func (s *service) GetLogUser(idsignature string) ([]models.UserLog, error) {
	log, err := s.repository.GetLogUser(idsignature)
	return log, err
}

func (s *service) GetUserByEmail(email string) (models.User, error) {
	user, err := s.repository.GetUserByEmail(email)
	return user, err
}
