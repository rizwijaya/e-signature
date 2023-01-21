package document

import (
	"e-signature/modules/v1/utilities/signatures/models"
	"fmt"
	"log"
	"os"
	"os/exec"

	pdf "github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/signintech/pdft"
)

type Documents interface {
	SignDocuments(imgpath string, input models.SignDocuments) string
	WaterMarking(path string) string
}

type documents struct {
}

func NewDocuments() *documents {
	return &documents{}
}

func (d *documents) SignDocuments(imgpath string, input models.SignDocuments) string {
	var pt pdft.PDFt
	inputPath := fmt.Sprintf("./public/temp/pdfsign/%s", input.Name)
	err := pt.Open(inputPath)
	if err != nil {
		log.Println(err)
		return ""
	}
	//Read PDF and Get pages size media box with pdfcpu
	dim, err := pdf.PageDimsFile(inputPath)
	if err != nil {
		log.Println(err)
		return ""
	}

	pageWidth := dim[int(input.SignPage)-1].Width
	pageHeight := dim[int(input.SignPage)-1].Height

	lCorrection := 1 + 0.039
	wCorrection := 1 - 0.05
	//tCorrection := 1 + 0.096
	x := input.X_coord * pageWidth * lCorrection
	y := input.Y_coord * pageHeight
	w := input.Width * pageWidth * wCorrection
	h := input.Height * pageHeight
	result := fmt.Sprintf("./public/temp/pdfsign/signed_%s", input.Name)
	mode := "sign"
	//Run Python Script document.py with Arguments
	out, err := exec.Command("cmd", "/c", "python3", "./document.py", "-m", mode, "-s", "\""+inputPath+"\"", "-d", "\""+result+"\"", "-i", "\""+imgpath+"\"", "-x", fmt.Sprintf("%v", x), "-y", fmt.Sprintf("%v", y), "-w", fmt.Sprintf("%v", w), "-t", fmt.Sprintf("%v", h), "-p", fmt.Sprintf("%v", input.SignPage)).Output()
	if err != nil {
		log.Println(err)
		return ""
	}
	fmt.Println(string(out))
	//check file exist
	if _, err := os.Stat(result); os.IsNotExist(err) {
		log.Println(err)
		return ""
	}
	return result
}

func (d *documents) WaterMarking(path string) string {
	x := 500
	y := 500
	w := 13.0
	h := 13.0
	mode := "watermark"
	//Run Python Script document.py with Arguments
	//In Linux running without cmd /c
	out, err := exec.Command("cmd", "/c", "python3", "./document.py", "-m", mode, "-s", "\""+path+"\"", "-d", "\""+path+"\"", "-x", fmt.Sprintf("%v", x), "-y", fmt.Sprintf("%v", y), "-w", fmt.Sprintf("%v", w), "-t", fmt.Sprintf("%v", h), "-p", "1").Output()
	if err != nil {
		log.Println(err)
		return path
	}
	fmt.Println(string(out))

	return path
}
