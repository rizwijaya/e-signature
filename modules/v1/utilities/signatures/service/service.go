package service

import (
	"e-signature/modules/v1/utilities/signatures/models"
	"e-signature/modules/v1/utilities/signatures/repository"
	modelsUser "e-signature/modules/v1/utilities/user/models"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/fogleman/gg"
)

type Service interface {
	CreateImgSignature(input models.AddSignature) string
	CreateImgSignatureData(input models.AddSignature, name string) string
	CreateLatinSignatures(user modelsUser.User) string
	CreateLatinSignaturesData(user modelsUser.User, latin string) string
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
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
	pngFilename := "public/images/signatures/signatures-" + input.Id + ".png"
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
	img := fmt.Sprintf("public/images/signatures/signatures-%s.png", input.Id)
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
	id := input.Id
	if len(input.Id) > 12 {
		id = input.Id[0:12]
	}
	nam := name
	if len(name) > 12 {
		nam = name[0:12]
	}
	dc.DrawStringAnchored(id, 145, 242, 0, 0.5)
	dc.DrawStringAnchored(nam, 145, 254, 0, 0.5)
	dc.DrawStringAnchored("rizwijaya.smartsign.com", 145, 266, 0, 0.5)
	dc.DrawStringAnchored("Integrate in Blockchain", 145, 278, 0, 0.5)
	dc.Clip()
	//dc.SavePNG("out.png")
	filename := fmt.Sprintf("public/images/signatures/signaturesdata-%s.png", input.Id)
	dc.SavePNG(filename)
	return filename
}

func (s *service) CreateLatinSignatures(user modelsUser.User) string {
	dc := gg.NewContext(300, 300)
	//dc.SetRGB(1, 1, 1) // white background
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("modules/v1/utilities/signatures/font/latin.ttf", 75); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored(user.Name[0:12], 300/2, 130, 0.5, 0.5)
	dc.Clip()
	filename := fmt.Sprintf("app/tmp/latin-%s.png", user.Idsignature)
	dc.SavePNG(filename)
	return filename
}

func (s *service) CreateLatinSignaturesData(user modelsUser.User, latin string) string {
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
	dc.DrawStringAnchored(user.Idsignature[0:12], 145, 242, 0, 0.5)
	dc.DrawStringAnchored(user.Name[0:12], 145, 254, 0, 0.5)
	dc.DrawStringAnchored("rizwijaya.smartsign.com", 145, 266, 0, 0.5)
	dc.DrawStringAnchored("Integrate in Blockchain", 145, 278, 0, 0.5)
	//dc.DrawStringAnchored("Integrate in Blockchain", 145, 290, 0, 0.5)
	dc.Clip()
	filename := fmt.Sprintf("app/tmp/latindata-%s.png", user.Idsignature)
	dc.SavePNG(filename)
	return filename
}
