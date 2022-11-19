package service

import (
	"bytes"
	"crypto/sha256"
	"e-signature/app/config"
	"e-signature/modules/v1/utilities/signatures/models"
	"e-signature/modules/v1/utilities/signatures/repository"
	modelsUser "e-signature/modules/v1/utilities/user/models"
	"encoding/base64"
	"fmt"
	"html/template"
	"image"
	"image/png"
	"io"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
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
	GetMySignature(sign string, id string, name string) (models.MySignatures, error)
	ChangeSignatures(sign_type string, idsignature string) error
	ResizeImages(mysign models.MySignatures, input models.SignDocuments) string
	SignDocuments(imgpath string, input models.SignDocuments) string
	InvitePeople(email string, input models.SignDocuments) error
	GenerateHashDocument(input string) string
	AddToBlockhain(input models.SignDocuments) error
	AddUserDocs(input models.SignDocuments) error
	DocumentSigned(sign models.SignDocs) error
	GetListDocument(publickey string) []models.ListDocument
	GetDocument(hash string, publickey string) models.DocumentBlockchain
	GetDocumentAllSign(hash string) (models.DocumentAllSign, bool)
	GetDocumentNoSigners(hash string) models.DocumentBlockchain
	GetTransactions() []models.Transac
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

var hari = [...]string{
	"Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"}

var bulan = [...]string{
	"Januari", "Februari", "Maret", "April", "Mei", "Juni",
	"Juli", "Agustus", "September", "Oktober", "Nopember", "Desember",
}

func TanggalJam(t time.Time) string {
	return fmt.Sprintf("%s, %02d %s %d | %02d:%02d WIB",
		hari[t.Weekday()], t.Day(), bulan[t.Month()-1][:3], t.Year(), t.Hour(), int(t.Minute()),
	)
}

func Tanggal(t time.Time) string {
	return fmt.Sprintf("%02d %s %d",
		t.Day(), bulan[t.Month()-1], t.Year())
}

func Clock(t time.Time) string {
	return fmt.Sprintf("%02d:%02d",
		t.Hour(), t.Minute())
}

func init() {
	//Demo: 58f9d9cb8ba964ee240d04196ee9b1e9406107505e36d553fc2c057002de766a
	//Develop:345b46cc6941c36f6d4528a304c7b6ceb1855e56a2c76ca0db93f3f4b3586904
	err := license.SetMeteredKey("58f9d9cb8ba964ee240d04196ee9b1e9406107505e36d553fc2c057002de766a")
	if err != nil {
		log.Println(err)
	}
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
	return TanggalJam(t)
}

func (s *service) CreateImgSignature(input models.AddSignature) string {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(input.Signature))
	m, formatString, err := image.Decode(reader)
	if err != nil {
		log.Println(err)
	}
	bounds := m.Bounds()
	fmt.Println(bounds, formatString)

	//Encode from image format to writer
	pngFilename := "public/images/signatures/signatures/signatures-" + input.Id + ".png"
	f, err := os.OpenFile(pngFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Println(err)
		return ""
	}

	err = png.Encode(f, m)
	if err != nil {
		log.Println(err)
		return ""
	}

	//fmt.Println("Png file", pngFilename, "created")
	return pngFilename
}

func (s *service) CreateImgSignatureData(input models.AddSignature, name string) string {
	//const S = 200
	//im, err := gg.LoadImage("signature-632fbcd7ea74d38e23360148.png")
	//img := fmt.Sprintf("app/tmp/signature-%s.png", user.Idsignature)
	img := fmt.Sprintf("public/images/signatures/signatures/signatures-%s.png", input.Id)
	im, err := gg.LoadImage(img)
	if err != nil {
		log.Println(err)
	}

	dc := gg.NewContext(300, 300)
	//dc.SetRGB(1, 1, 1) // white background
	dc.Clear()
	dc.SetRGB255(25, 25, 112)
	if err := dc.LoadFontFace("modules/v1/utilities/signatures/font/detail_data.ttf", 11); err != nil {
		log.Println(err)
	}
	//dc.DrawStringAnchored("Hello, world!", 460/4, 180/2, 0.5, 0.5)

	//dc.DrawRoundedRectangle(0, 0, 460/4, 180/2, 0)
	dc.DrawImage(im, 0, 0)
	// id := input.Id
	// if len(input.Id) > 12 {
	// 	id = input.Id[0:12]
	// }
	nam := name
	if len(name) > 12 {
		nam = name[0:12]
	}
	//dc.DrawStringAnchored(id, 145, 242, 0, 0.5)
	dc.DrawStringAnchored("Sign-"+nam, 145, 254, 0, 0.5)
	dc.DrawStringAnchored("rizwijaya.smartsign.com", 145, 266, 0, 0.5)
	dc.DrawStringAnchored("Integrate in Blockchain", 145, 278, 0, 0.5)
	dc.Clip()
	//dc.SavePNG("out.png")
	filename := fmt.Sprintf("public/images/signatures/signatures_data/signaturesdata-%s.png", input.Id)
	if _, err := os.Stat(filename); err == nil {
		os.Remove(filename)
	}

	dc.SavePNG(filename)
	return filename
}

func (s *service) CreateLatinSignatures(user modelsUser.User, id string) string {
	dc := gg.NewContext(300, 300)
	//dc.SetRGB(1, 1, 1) // white background
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("modules/v1/utilities/signatures/font/latin.ttf", 75); err != nil {
		log.Println(err)
	}
	name := user.Name
	if len(user.Name) > 12 {
		name = user.Name[0:12]
	}

	dc.DrawStringAnchored(name, 300/2, 130, 0.5, 0.5)
	dc.Clip()
	filename := fmt.Sprintf("public/images/signatures/latin/latin-%s.png", id)
	dc.SavePNG(filename)
	return filename
}

func (s *service) CreateLatinSignaturesData(user modelsUser.User, latin string, idn string) string {
	im, err := gg.LoadImage(latin)
	if err != nil {
		log.Println(err)
	}

	dc := gg.NewContext(300, 300)
	//dc.SetRGB(1, 1, 1) // white background
	dc.Clear()
	dc.SetRGB255(25, 25, 112)
	if err := dc.LoadFontFace("modules/v1/utilities/signatures/font/detail_data.ttf", 11); err != nil {
		log.Println(err)
	}

	dc.DrawImage(im, 0, 0)
	id := user.Idsignature
	if len(user.Idsignature) > 12 {
		id = user.Idsignature[0:12]
	}
	// nam := user.Name
	// if len(user.Name) > 12 {
	// 	nam = user.Name[0:12]
	// }
	//dc.DrawStringAnchored(id, 145, 242, 0, 0.5)
	dc.DrawStringAnchored("Sign-"+id, 145, 254, 0, 0.5)
	dc.DrawStringAnchored("rizwijaya.smartsign.com", 145, 266, 0, 0.5)
	dc.DrawStringAnchored("Integrate in Blockchain", 145, 278, 0, 0.5)
	//dc.DrawStringAnchored("Integrate in Blockchain", 145, 290, 0, 0.5)
	dc.Clip()
	filename := fmt.Sprintf("public/images/signatures/latin_data/latindata-%s.png", idn)
	dc.SavePNG(filename)
	return filename
}

func (s *service) DefaultSignatures(user modelsUser.User, id string) error {
	err := s.repository.DefaultSignatures(user, id)
	return err
}

func (s *service) UpdateMySignatures(signature string, signaturedata string, sign string) error {
	err := s.repository.UpdateMySignatures(signature, signaturedata, sign)
	return err
}

func (s *service) GetMySignature(sign string, id string, name string) (models.MySignatures, error) {
	signature, err := s.repository.GetMySignature(sign)
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
		Date_update:        fmt.Sprintf("%s | %s WIB", Tanggal(signature.Date_update), Clock(signature.Date_update)),
		Date_created:       fmt.Sprintf("%s | %s WIB", Tanggal(signature.Date_created), Clock(signature.Date_created)),
	}
	return mysign, err
}

func (s *service) ChangeSignatures(sign_type string, idsignature string) error {
	err := s.repository.ChangeSignature(sign_type, idsignature)
	return err
}

func (s *service) ResizeImages(mysign models.MySignatures, input models.SignDocuments) string {
	signatures := mysign.Latin
	if mysign.Signature_selected == "signature" {
		signatures = mysign.Signature
	} else if mysign.Signature_selected == "signature_data" {
		signatures = mysign.Signature_data
	} else if mysign.Signature_selected == "latin" {
		signatures = mysign.Latin
	} else if mysign.Signature_selected == "latin_data" {
		signatures = mysign.Latin_data
	}
	path := fmt.Sprintf("./public/images/signatures/%s", signatures)
	r, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}
	defer r.Close()
	img, err := png.Decode(r)
	if err != nil {
		log.Println(err)
	}
	m := resize.Resize(uint(input.Width*0.75), uint(input.Height*0.75), img, resize.Lanczos3)
	path2 := fmt.Sprintf("./public/temp/sizes-%s.png", mysign.Signature_selected)
	out, err := os.Create(path2)
	if err != nil {
		log.Println(err)
	}
	defer out.Close()
	png.Encode(out, m)
	return path2
}

func (s *service) SignDocuments(imgpath string, input models.SignDocuments) string {
	c := creator.New()
	img, err := c.NewImageFromFile(imgpath)
	if err != nil {
		log.Println(err)
		return ""
	}
	inputPath := fmt.Sprintf("./public/temp/pdfsign/%s", input.Name)
	// Read the input pdf file.
	f, err := os.Open(inputPath)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer f.Close()

	pdfReader, err := model.NewPdfReader(f)
	if err != nil {
		log.Println(err)
		return ""
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		log.Println(err)
		return ""
	}

	// Load the pages.
	for i := 0; i < numPages; i++ {
		page, err := pdfReader.GetPage(i + 1)
		if err != nil {
			log.Println(err)
			return ""
		}
		// Add the page.
		err = c.AddPage(page)
		if err != nil {
			log.Println(err)
			return ""
		}

		// If the specified page, or -1, apply the image to the page.
		if i+1 == int(input.SignPage) || int(input.SignPage) == -1 {
			positionedImg := calcImagePos(img, page, input)
			err = c.Draw(positionedImg)
			if err != nil {
				log.Println(err)
				return ""
			}
		}
	}
	inputPath2 := fmt.Sprintf("./public/temp/pdfsign/signed_%s", input.Name)
	err = c.WriteToFile(inputPath2)
	if err != nil {
		log.Println(err)
		return ""
	}
	return inputPath2
}

func calcImagePos(img *creator.Image, page *model.PdfPage, input models.SignDocuments) *creator.Image {
	bbox, err := page.GetMediaBox()

	if err != nil {
		log.Println(err)
		return nil
	}

	pageWidth := (*bbox).Urx - (*bbox).Llx
	pageHeight := (*bbox).Ury - (*bbox).Lly

	// X and Y is actually a image pos relative
	// to page width (element width in html)
	lCorrection := 1 + 0.09
	wCorrection := 1 - 0.05
	imgLeft := input.X_coord * pageWidth * lCorrection
	imgTop := input.Y_coord * pageHeight
	imgWidth := input.Width * pageWidth * wCorrection
	imgHeight := input.Height * pageHeight

	img.SetPos(imgLeft, imgTop)
	img.SetWidth(imgWidth)
	img.SetHeight(imgHeight)

	return img
}

func (s *service) InvitePeople(email string, input models.SignDocuments) error {
	conf, _ := config.Init()
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
	if err := t.Execute(&tpl, input); err != nil {
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
		listDoc[i].Date_created_WIB = TanggalJam(listDoc[i].Date_created)
		listDoc[i].Documents = s.repository.GetDocument(listDoc[i].Hash_original, publickey)
		listDoc[i].Documents.Signers = s.repository.GetSigners(listDoc[i].Hash_original, publickey)
	}
	// fmt.Println(publickey)
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
	docSigned.Visibility = doc.Visibility
	docSigned.Createdtime = s.TimeFormating(doc.Createdtime)
	docSigned.Completedtime = s.TimeFormating(doc.Completedtime)
	docSigned.Exist = doc.Exist
	//Get list sign documents from db
	signData := s.repository.GetListSign(doc.Hash_ori)
	for i := range signData {
		//Get Signer Data in Blockchain
		signer := s.repository.GetSigners(hash_ori, signData[i].Sign_addr)
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
		transac[i].Date_created_wib = TanggalJam(transac[i].Date_created)
		transac[i].Ids = transac[i].Id.Hex()
	}
	return transac
}
