package images

import (
	"e-signature/modules/v1/utilities/signatures/models"
	modelsUser "e-signature/modules/v1/utilities/user/models"
	times "e-signature/pkg/time"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fogleman/gg"
)

type Images interface {
	CreateImageSignature(input models.AddSignature) string
	CreateImgSignatureData(input models.AddSignature, name string, font string) string
	CreateLatinSignatures(user modelsUser.User, id string, font string) string
	CreateLatinSignaturesData(user modelsUser.User, latin string, idn string, font string) string
	ResizeImages(mysign models.MySignatures, input models.SignDocuments) string
	SignWithData(path string, sign string) string
}

type images struct {
}

func NewImages() *images {
	return &images{}
}

func (i *images) CreateImageSignature(input models.AddSignature) string {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(input.Signature))
	m, formatString, err := image.Decode(reader)
	if err != nil {
		log.Println(err)
		return ""
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

func (i *images) CreateImgSignatureData(input models.AddSignature, name string, font string) string {
	// location, _ := time.LoadLocation("Asia/Jakarta")
	// created := times.TanggalJamEnglish(time.Now().In(location))
	img := fmt.Sprintf("public/images/signatures/signatures/signatures-%s.png", input.Id)
	im, err := gg.LoadImage(img)
	if err != nil {
		log.Println(err)
		return ""
	}

	dc := gg.NewContext(300, 300)
	dc.Clear()
	dc.SetRGB255(25, 25, 112)
	if err := dc.LoadFontFace("modules/v1/utilities/signatures/font/"+font, 11); err != nil {
		log.Println(err)
		return ""
	}

	dc.DrawImage(im, 0, 0)
	nam := name
	if len(name) > 12 {
		nam = name[0:12]
	}
	dc.DrawStringAnchored("Sign by "+nam, 40, 242, 0, 0.5)
	dc.DrawStringAnchored("Verify at smartsign.rizwijaya.com", 40, 254, 0, 0.5)
	dc.DrawStringAnchored("Digital Signature In Blockchain", 40, 266, 0, 0.5)
	//dc.DrawStringAnchored("Created at "+created, 40, 278, 0, 0.5)

	dc.Clip()
	filename := fmt.Sprintf("public/images/signatures/signatures_data/signaturesdata-%s.png", input.Id)
	if _, err := os.Stat(filename); err == nil {
		os.Remove(filename)
	}

	dc.SavePNG(filename)
	return filename
}

func (i *images) CreateLatinSignatures(user modelsUser.User, id string, font string) string {
	dc := gg.NewContext(300, 300)
	//dc.SetRGB(1, 1, 1) // white background
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("modules/v1/utilities/signatures/font/"+font, 75); err != nil {
		log.Println(err)
		return ""
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

func (i *images) CreateLatinSignaturesData(user modelsUser.User, latin string, idn string, font string) string {
	// location, _ := time.LoadLocation("Asia/Jakarta")
	// created := times.TanggalJamEnglish(time.Now().In(location))
	im, err := gg.LoadImage(latin)
	if err != nil {
		log.Println(err)
		return ""
	}

	dc := gg.NewContext(300, 300)
	//dc.SetRGB(1, 1, 1) // white background
	dc.Clear()
	dc.SetRGB255(25, 25, 112)
	if err := dc.LoadFontFace("modules/v1/utilities/signatures/font/"+font, 11); err != nil {
		log.Println(err)
		return ""
	}

	dc.DrawImage(im, 0, 0)
	id := user.Idsignature
	if len(user.Idsignature) > 12 {
		id = user.Idsignature[0:12]
	}

	dc.DrawStringAnchored("Sign by "+id, 40, 242, 0, 0.5)
	dc.DrawStringAnchored("Verify at smartsign.rizwijaya.com", 40, 254, 0, 0.5)
	dc.DrawStringAnchored("Digital Signature In Blockchain", 40, 266, 0, 0.5)
	dc.Clip()
	filename := fmt.Sprintf("public/images/signatures/latin_data/latindata-%s.png", idn)
	dc.SavePNG(filename)
	return filename
}

func (i *images) ResizeImages(mysign models.MySignatures, input models.SignDocuments) string {
	signatures := mysign.Latin
	path := fmt.Sprintf("./public/images/signatures/%s", signatures)
	if mysign.Signature_selected == "signature" {
		signatures = mysign.Signature
		path = fmt.Sprintf("./public/images/signatures/%s", signatures)
	} else if mysign.Signature_selected == "signature_data" {
		signatures = mysign.Signature_data
		path = fmt.Sprintf("./public/images/signatures/%s", signatures)
		path = i.SignWithData(path, mysign.User_id)
	} else if mysign.Signature_selected == "latin" {
		signatures = mysign.Latin
		path = fmt.Sprintf("./public/images/signatures/%s", signatures)
	} else if mysign.Signature_selected == "latin_data" {
		signatures = mysign.Latin_data
		path = fmt.Sprintf("./public/images/signatures/%s", signatures)
		path = i.SignWithData(path, mysign.User_id)
	}

	return path
}

func (i *images) SignWithData(path string, sign string) string {
	location, _ := time.LoadLocation("Asia/Jakarta")
	created := times.TanggalJamEnglish(time.Now().In(location))
	font := "data.ttf"
	im, err := gg.LoadImage(path)
	if err != nil {
		log.Println(err)
		return ""
	}

	dc := gg.NewContext(300, 300)
	dc.Clear()
	dc.SetRGB255(25, 25, 112)
	if err := dc.LoadFontFace("modules/v1/utilities/signatures/font/"+font, 11); err != nil {
		log.Println(err)
		return ""
	}

	dc.DrawImage(im, 0, 0)
	dc.DrawStringAnchored("Created at "+created, 40, 278, 0, 0.5)

	dc.Clip()
	filename := fmt.Sprintf("./public/temp/signaturesdata-%s.png", sign)
	if _, err := os.Stat(filename); err == nil {
		os.Remove(filename)
	}

	dc.SavePNG(filename)
	return filename
}
