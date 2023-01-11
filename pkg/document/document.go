package document

import (
	"e-signature/app/config"
	"e-signature/modules/v1/utilities/signatures/models"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
)

type Documents interface {
	//Init()
	//CalcImagePos(img *creator.Image, page *model.PdfPage, input models.SignDocuments) *creator.Image
	SignDocuments(imgpath string, input models.SignDocuments) string
	WaterMarking(path string) string
}

type documents struct {
}

func NewDocuments() *documents {
	return &documents{}
}

func (d *documents) Init() {
	conf, _ := config.Init()
	err := license.SetMeteredKey(conf.Signature.Key)
	if err != nil {
		log.Println(err)
	}
}

func (d *documents) CalcImagePos(img *creator.Image, page *model.PdfPage, input models.SignDocuments) *creator.Image {
	bbox, err := page.GetMediaBox()

	if err != nil {
		log.Println(err)
		return nil
	}

	pageWidth := (*bbox).Urx - (*bbox).Llx
	pageHeight := (*bbox).Ury - (*bbox).Lly

	// X and Y is actually a image pos relative
	// to page width (element width in html)
	lCorrection := 1 + 0.039
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
	d.Init()
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
			positionedImg := d.CalcImagePos(img, page, input)
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

func (d *documents) WaterMarking(path string) string {
	d.Init()
	//Add Invisible text to pdf
	location, _ := time.LoadLocation("Asia/Jakarta")
	t := time.Now().In(location)
	pageNum := -1
	text := "smartsign at " + t.Format("02-01-2006 15:04:05")
	xPos := 0.000000001
	yPos := 0.000000001

	f, err := os.Open(path)
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

	c := creator.New()

	// Load the pages.
	for i := 0; i < numPages; i++ {
		page, err := pdfReader.GetPage(i + 1)
		if err != nil {
			log.Println(err)
			return ""
		}

		err = c.AddPage(page)
		if err != nil {
			log.Println(err)
			return ""
		}

		if i == pageNum || pageNum == -1 {
			p := c.NewParagraph(text)
			// Change to times bold font (default is helvetica).
			timesBold, err := model.NewStandard14Font("Times-Bold")
			if err != nil {
				panic(err)
			}
			p.SetColor(creator.ColorRGBFromHex("#ffffff00"))
			p.SetFontSize(0)
			p.SetFont(timesBold)
			p.SetPos(xPos, yPos)

			_ = c.Draw(p)
		}

	}

	err = c.WriteToFile(path)
	if err != nil {
		log.Println(err)
		return ""
	}
	return path
}

// func (d *documents) SignDocuments(imgpath string, input models.SignDocuments) string {
// 	var pt pdft.PDFt
// 	inputPath := fmt.Sprintf("./public/temp/pdfsign/%s", input.Name)
// 	err := pt.Open(inputPath)
// 	if err != nil {
// 		log.Println(err)
// 		return ""
// 	}
// 	//Read PDF and Get pages size media box with pdfcpu
// 	dim, err := pdf.PageDimsFile(inputPath)
// 	if err != nil {
// 		log.Println(err)
// 		return ""
// 	}

// 	pic, err := ioutil.ReadFile(imgpath)
// 	if err != nil {
// 		log.Println(err)
// 		return ""
// 	}

// 	pageWidth := dim[int(input.SignPage)-1].Width
// 	pageHeight := dim[int(input.SignPage)-1].Height

// 	lCorrection := 1 + 0.038
// 	wCorrection := 1 - 0.05
// 	tCorrection := 1 + 0.096
// 	imgLeft := input.X_coord * pageWidth * lCorrection
// 	imgTop := input.Y_coord * pageHeight * tCorrection
// 	imgWidth := input.Width * pageWidth * wCorrection
// 	imgHeight := input.Height * pageHeight

// 	// insert image to pdf
// 	err = pt.InsertImg(pic, int(input.SignPage), imgLeft, imgTop, imgWidth, imgHeight)
// 	if err != nil {
// 		log.Println(err)
// 		return ""
// 	}
// 	inputPath2 := fmt.Sprintf("./public/temp/pdfsign/signed_%s", input.Name)
// 	err = pt.Save(inputPath2)
// 	if err != nil {
// 		log.Println(err)
// 		return ""
// 	}
// 	return inputPath2
// }
