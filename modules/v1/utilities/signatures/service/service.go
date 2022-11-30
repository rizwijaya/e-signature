package service

import (
	"bytes"
	"crypto/sha256"
	"e-signature/app/config"
	"e-signature/modules/v1/utilities/signatures/models"
	"e-signature/modules/v1/utilities/signatures/repository"
	modelsUser "e-signature/modules/v1/utilities/user/models"
	docs "e-signature/pkg/document"
	img "e-signature/pkg/images"
	tm "e-signature/pkg/time"
	"fmt"
	"html/template"
	"io"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"

	gomail "gopkg.in/gomail.v2"
)

type Service interface {
	TimeFormating(times string) string
	CreateImgSignature(input models.AddSignature) string
	CreateImgSignatureData(input models.AddSignature, name string) string
	CreateLatinSignatures(user modelsUser.User, id string) string
	CreateLatinSignaturesData(user modelsUser.User, latin string, idn string) string
	DefaultSignatures(user modelsUser.User, id string) error
	UpdateMySignatures(signature string, signaturedata string, sign string) error
	GetMySignature(sign string, id string, name string) models.MySignatures
	ChangeSignatures(sign_type string, idsignature string) error
	ResizeImages(mysign models.MySignatures, input models.SignDocuments) string
	SignDocuments(imgpath string, input models.SignDocuments) string
	InvitePeople(email string, input models.SignDocuments, users modelsUser.User) error
	GenerateHashDocument(input string) string
	AddToBlockhain(input models.SignDocuments) error
	AddUserDocs(input models.SignDocuments) error
	DocumentSigned(sign models.SignDocs) error
	GetListDocument(publickey string) []models.ListDocument
	GetDocument(hash string, publickey string) models.DocumentBlockchain
	GetDocumentAllSign(hash string) (models.DocumentAllSign, bool)
	GetDocumentNoSigners(hash string) models.DocumentBlockchain
	GetTransactions() []models.Transac
	CheckSignature(hash string, publickey string) bool
}

type service struct {
	repository repository.Repository
	images     img.Images
	documents  docs.Documents
}

func NewService(repository repository.Repository, img img.Images, documents docs.Documents) *service {
	return &service{repository, img, documents}
}

func (s *service) TimeFormating(times string) string {
	if len(times) == 13 {
		times = "0" + times
	}
	h, _ := strconv.Atoi(times[0:2])
	m, _ := strconv.Atoi(times[2:4])
	se, _ := strconv.Atoi(times[4:6])
	d, _ := strconv.Atoi(times[6:8])
	mo, _ := strconv.Atoi(times[8:10])
	y, _ := strconv.Atoi(times[10:14])
	t := time.Date(y, time.Month(mo), d, h, m, se, 0, time.UTC)
	return tm.TanggalJam(t)
}

func (s *service) CreateImgSignature(input models.AddSignature) string {
	return s.images.CreateImageSignature(input)
}

func (s *service) CreateImgSignatureData(input models.AddSignature, name string) string {
	font := "detail_data.ttf"
	return s.images.CreateImgSignatureData(input, name, font)
}

func (s *service) CreateLatinSignatures(user modelsUser.User, id string) string {
	latin := "latin.ttf"
	return s.images.CreateLatinSignatures(user, id, latin)
}

func (s *service) CreateLatinSignaturesData(user modelsUser.User, latin string, idn string) string {
	font := "detail_data.ttf"
	return s.images.CreateLatinSignaturesData(user, latin, idn, font)
}

func (s *service) DefaultSignatures(user modelsUser.User, id string) error {
	err := s.repository.DefaultSignatures(user, id)
	return err
}

func (s *service) UpdateMySignatures(signature string, signaturedata string, sign string) error {
	return s.repository.UpdateMySignatures(signature, signaturedata, sign)
}

func (s *service) GetMySignature(sign string, id string, name string) models.MySignatures {
	signature, err := s.repository.GetMySignature(sign)
	if err != nil {
		log.Println(err)
	}
	mysign := models.MySignatures{
		Id:                 signature.Id.Hex(),
		Name:               name,
		User_id:            id,
		Signature:          fmt.Sprintf("signatures/%s", signature.Signature),
		Signature_id:       fmt.Sprintf("sign-%s", signature.Id.Hex()),
		Signature_data:     fmt.Sprintf("signatures_data/%s", signature.Signature_data),
		Signature_data_id:  fmt.Sprintf("sign_data-%s", signature.Id.Hex()),
		Latin:              fmt.Sprintf("latin/%s", signature.Latin),
		Latin_id:           fmt.Sprintf("latin-%s", signature.Id.Hex()),
		Latin_data:         fmt.Sprintf("latin_data/%s", signature.Latin_data),
		Latin_data_id:      fmt.Sprintf("latin_data-%s", signature.Id.Hex()),
		Signature_selected: signature.Signature_selected,
		Date_update:        fmt.Sprintf("%s | %s WIB", tm.Tanggal(signature.Date_update), tm.Jam(signature.Date_update)),
		Date_created:       fmt.Sprintf("%s | %s WIB", tm.Tanggal(signature.Date_created), tm.Jam(signature.Date_created)),
	}
	return mysign
}

func (s *service) ChangeSignatures(sign_type string, idsignature string) error {
	return s.repository.ChangeSignature(sign_type, idsignature)
}

func (s *service) ResizeImages(mysign models.MySignatures, input models.SignDocuments) string {
	return s.images.ResizeImages(mysign, input)
}

func (s *service) SignDocuments(imgpath string, input models.SignDocuments) string {
	return s.documents.SignDocuments(imgpath, input)
}

func (s *service) InvitePeople(email string, input models.SignDocuments, users modelsUser.User) error {
	conf, _ := config.Init()
	emData := struct {
		Judul         string
		Creator_id    string
		Note          string
		Hash_original string
		Name          string
		Member_name   string
	}{
		Judul:         input.Judul,
		Creator_id:    input.Creator_id,
		Note:          input.Note,
		Hash_original: input.Hash_original,
		Name:          input.Name,
		Member_name:   users.Name,
	}
	// Parse the html file.
	dir := "./public/templates/users/pages/email.html"
	t := template.New("email.html")
	var err error
	t, err = t.ParseFiles(dir)
	if err != nil {
		log.Println(err)
		return err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, emData); err != nil {
		log.Println(err)
		return err
	}

	result := tpl.String()

	// Set up authentication information for send email.
	m := gomail.NewMessage()
	m.SetHeader("From", conf.Email.User)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Permintaan Tanda Tangan Digital - SmartSign")
	m.SetBody("text/html", result)
	port, _ := strconv.Atoi(conf.Email.Port)
	d := gomail.NewDialer(conf.Email.Host, port, conf.Email.User, conf.Email.Pass)

	// Send the email.
	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *service) GenerateHashDocument(input string) string {
	f, err := os.Open(input)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Println(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (s *service) AddToBlockhain(input models.SignDocuments) error {
	timeSign := new(big.Int)
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return err
	}
	timeNow := time.Now().In(location)
	timeFormat := timeNow.Format("15040502012006")
	timeSign, _ = timeSign.SetString(timeFormat, 10)
	err = s.repository.AddToBlockhain(input, timeSign)
	return err
}

func (s *service) AddUserDocs(input models.SignDocuments) error {
	err := s.repository.AddUserDocs(input)
	return err
}

func (s *service) DocumentSigned(sign models.SignDocs) error {
	timeSign := new(big.Int)
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Println(err)
		return err
	}
	timeNow := time.Now().In(location)
	timeFormat := timeNow.Format("15040502012006")
	timeSign, _ = timeSign.SetString(timeFormat, 10)
	err = s.repository.DocumentSigned(sign, timeSign)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *service) GetListDocument(publickey string) []models.ListDocument {
	listDoc := s.repository.ListDocumentNoSign(publickey)
	for i := range listDoc {
		listDoc[i].Documents = s.repository.GetDocument(listDoc[i].Hash_original, publickey)
		listDoc[i].Documents.Signers = s.repository.GetSigners(listDoc[i].Hash_original, publickey)
	}
	return listDoc
}

func (s *service) GetDocument(hash string, publickey string) models.DocumentBlockchain {
	doc := s.repository.GetDocument(hash, publickey)
	doc.Signers = s.repository.GetSigners(hash, publickey)
	return doc
}

func (s *service) GetDocumentAllSign(hash string) (models.DocumentAllSign, bool) {
	conf, _ := config.Init()
	var docSigned models.DocumentAllSign
	checkDoc := s.repository.VerifyDoc(hash)
	if !checkDoc {
		return docSigned, false
	}
	//Get Original Hash from Blockchain with Hash signed
	hash_ori := s.repository.GetHashOriginal(hash, "0x"+conf.Blockhain.Public)
	doc := s.repository.GetDocument(hash_ori, "0x"+conf.Blockhain.Public)
	docSigned.Document_id = doc.Document_id
	docSigned.Creator = doc.Creator
	docSigned.Creator_id = doc.Creator_id
	docSigned.Metadata = doc.Metadata
	docSigned.Hash_ori = doc.Hash_ori
	docSigned.Hash = doc.Hash
	docSigned.IPFS = doc.IPFS
	docSigned.State = doc.State
	docSigned.Mode = doc.Mode
	docSigned.Createdtime = s.TimeFormating(doc.Createdtime)
	docSigned.Completedtime = s.TimeFormating(doc.Completedtime)
	docSigned.Exist = doc.Exist
	//Get list sign documents from db
	signData := s.repository.GetListSign(doc.Hash_ori)
	for i := range signData {
		//Get Signer Data in Blockchain
		signer := s.repository.GetSigners(hash_ori, signData[i].Sign_addr)
		if signer.Signers_state {
			//Get Signer Data in Database
			signDB := s.repository.GetUserByIdSignatures(signer.Signers_id)
			SignersData := models.SignersData{
				Sign_addr:     signer.Sign_addr.String(),
				Sign_id:       signer.Sign_id,
				Signers_id:    signer.Signers_id,
				Signers_hash:  signer.Signers_hash,
				Signers_state: signer.Signers_state,
				Sign_time:     s.TimeFormating(signer.Sign_time),
				Sign_name:     signDB.Name,
				Sign_email:    signDB.Email,
				Sign_id_db:    signDB.Id.String(),
			}
			docSigned.Signers = append(docSigned.Signers, SignersData)
		}
	}
	return docSigned, true
}

func (s *service) GetDocumentNoSigners(hash string) models.DocumentBlockchain {
	conf, _ := config.Init()
	doc := s.repository.GetDocument(hash, "0x"+conf.Blockhain.Public)
	return doc
}

func (s *service) GetTransactions() []models.Transac {
	transac := s.repository.GetTransactions()
	for i := range transac {
		transac[i].Ids = transac[i].Id.Hex()
	}
	return transac
}

func (s *service) CheckSignature(hash string, publickey string) bool {
	check := s.repository.CheckSignature(hash, publickey)
	return check
}
