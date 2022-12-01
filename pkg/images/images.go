package images

import (
	"e-signature/modules/v1/utilities/signatures/models"
	modelsUser "e-signature/modules/v1/utilities/user/models"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
)

type Images interface {
	CreateImageSignature(input models.AddSignature) string
	CreateImgSignatureData(input models.AddSignature, name string, font string) string
	CreateLatinSignatures(user modelsUser.User, id string, font string) string
	CreateLatinSignaturesData(user modelsUser.User, latin string, idn string, font string) string
	ResizeImages(mysign models.MySignatures, input models.SignDocuments) string
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
	if err := dc.LoadFontFace("modules/v1/utilities/signatures/font/"+font, 11); err != nil {
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
	dc.DrawStringAnchored("Sign by "+nam, 145, 254, 0, 0.5)
	dc.DrawStringAnchored("smartsign.rizwijaya.com", 145, 266, 0, 0.5)
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

func (i *images) CreateLatinSignatures(user modelsUser.User, id string, font string) string {
	dc := gg.NewContext(300, 300)
	//dc.SetRGB(1, 1, 1) // white background
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("modules/v1/utilities/signatures/font/"+font, 75); err != nil {
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

func (i *images) CreateLatinSignaturesData(user modelsUser.User, latin string, idn string, font string) string {
	im, err := gg.LoadImage(latin)
	if err != nil {
		log.Println(err)
	}

	dc := gg.NewContext(300, 300)
	//dc.SetRGB(1, 1, 1) // white background
	dc.Clear()
	dc.SetRGB255(25, 25, 112)
	if err := dc.LoadFontFace("modules/v1/utilities/signatures/font/"+font, 11); err != nil {
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

func (i *images) ResizeImages(mysign models.MySignatures, input models.SignDocuments) string {
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
