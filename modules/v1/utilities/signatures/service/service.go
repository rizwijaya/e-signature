package service

import (
	"crypto/sha256"
	"e-signature/modules/v1/utilities/signatures/models"
	"e-signature/modules/v1/utilities/signatures/repository"
	modelsUser "e-signature/modules/v1/utilities/user/models"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"image"
	"image/png"
	"io"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
)

type Service interface {
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
	InvitePeople(email []string) error
	GenerateHashDocument(input string) string
	AddToBlockhain(input models.SignDocuments) error
	AddUserDocs(input models.SignDocuments) error
	DocumentSigned(sign models.SignDocs) error
	GetListDocument(publickey string) []models.ListDocument
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

func (s *service) CreateImgSignature(input models.AddSignature) string {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(input.Signature))
	m, formatString, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()
	fmt.Println(bounds, formatString)

	//Encode from image format to writer
	pngFilename := "public/images/signatures/signatures/signatures-" + input.Id + ".png"
	f, err := os.OpenFile(pngFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	err = png.Encode(f, m)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	dc := gg.NewContext(300, 300)
	//dc.SetRGB(1, 1, 1) // white background
	dc.Clear()
	dc.SetRGB255(25, 25, 112)
	if err := dc.LoadFontFace("modules/v1/utilities/signatures/font/detail_data.ttf", 11); err != nil {
		panic(err)
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
		panic(err)
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
		log.Fatal(err)
	}

	dc := gg.NewContext(300, 300)
	//dc.SetRGB(1, 1, 1) // white background
	dc.Clear()
	dc.SetRGB255(25, 25, 112)
	if err := dc.LoadFontFace("modules/v1/utilities/signatures/font/detail_data.ttf", 11); err != nil {
		panic(err)
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
		log.Fatal(err)
	}
	defer r.Close()
	img, err := png.Decode(r)
	if err != nil {
		log.Fatal(err)
	}
	m := resize.Resize(uint(input.Width*0.75), uint(input.Height*0.75), img, resize.Lanczos3)
	path2 := fmt.Sprintf("./public/temp/sizes-%s.png", mysign.Signature_selected)
	out, err := os.Create(path2)
	if err != nil {
		log.Fatal(err)
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
	inputPath := fmt.Sprintf("./public/temp/%s", input.Name)
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
	inputPath2 := fmt.Sprintf("./public/temp/signed_%s", input.Name)
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

func (s *service) InvitePeople(email []string) error {
	// const CONFIG_SMTP_HOST = "smtp.gmail.com"
	// const CONFIG_SMTP_PORT = 587
	// const CONFIG_SENDER_NAME = "PT. Telkom Indonesia <contact@tamaska.id>"
	// const CONFIG_AUTH_EMAIL = "contact@tamaska.id"
	// const CONFIG_AUTH_PASSWORD = "uyhiqdzkcknojmfh"
	// to := []string{"contact@tamaska.id"}
	// cc := []string{}

	// body := "Subject: " + formHubungiKami.SenderSubject + "\n\n" + "Nama: " + formHubungiKami.SenderName + "\n\n" + "Nama Instansi: " + formHubungiKami.SenderInstution + "\n\n" + "Nomor HP: " + formHubungiKami.SenderPhone + "\n\n" + "Email: " + formHubungiKami.SenderEmail + "\n\n" + "Pesan: " + formHubungiKami.SenderMsg + "\n\n"

	// auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
	// smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	// err := smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, append(to, cc...), []byte(body))
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (s *service) GenerateHashDocument(input string) string {
	file := strings.NewReader(input)
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(hash.Sum(nil))
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
		listDoc[i].Documents = s.repository.GetDocument(listDoc[i].Hash, publickey)
		//fmt.Println(listDoc[i].Documents)
	}
	// fmt.Println(publickey)
	return listDoc
}
