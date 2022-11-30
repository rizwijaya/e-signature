package document

import (
	"e-signature/app/config"
	"e-signature/modules/v1/utilities/signatures/models"
	"fmt"
	"log"
	"os"

	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
)

type Documents interface {
	calcImagePos(img *creator.Image, page *model.PdfPage, input models.SignDocuments) *creator.Image
	SignDocuments(imgpath string, input models.SignDocuments) string
}

type documents struct {
}

func NewDocuments() *documents {
	return &documents{}
}

func init() {
	conf, _ := config.Init()
	err := license.SetMeteredKey(conf.Signature.Key)
	if err != nil {
		log.Println(err)
	}
}

func (d *documents) calcImagePos(img *creator.Image, page *model.PdfPage, input models.SignDocuments) *creator.Image {
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

func (d *documents) SignDocuments(imgpath string, input models.SignDocuments) string {
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
			positionedImg := d.calcImagePos(img, page, input)
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
