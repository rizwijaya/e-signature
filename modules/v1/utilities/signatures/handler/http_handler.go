package signatures

import (
	"e-signature/app/config"
	"e-signature/modules/v1/utilities/signatures/models"
	api "e-signature/pkg/api_response"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (h *signaturesHandler) AddSignatures(c *gin.Context) {
	sessions := sessions.Default(c)
	var input models.AddSignature
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := api.APIRespon(err.Error(), 300, "error", nil)
		c.JSON(300, response)
		return
	}

	sign := fmt.Sprintf("%v", sessions.Get("sign"))
	h.signaturesService.CreateImgSignature(input)
	h.signaturesService.CreateImgSignatureData(input, sign)
	//Update Database MySignatures
	h.signaturesService.UpdateMySignatures(fmt.Sprintf("signatures-%s.png", input.Id), fmt.Sprintf("signaturesdata-%s.png", input.Id), sign)

	//Return Response API
	response := api.APIRespon("Success Add Signatures", 200, "success", nil)
	c.JSON(200, response)
}

func (h *signaturesHandler) ChangeSignatures(c *gin.Context) {
	var signing string
	session := sessions.Default(c)
	sign_type := c.Param("sign_type")
	if sign_type == "signature" || sign_type == "signature_data" || sign_type == "latin" || sign_type == "latin_data" {
		signing = sign_type
	} else {
		signing = "latin"
	}
	user := fmt.Sprintf("%v", session.Get("sign"))
	h.signaturesService.ChangeSignatures(signing, user)
	c.Redirect(302, "/my-signatures")
}

func (h *signaturesHandler) SignDocuments(c *gin.Context) {
	conf, _ := config.Init()
	session := sessions.Default(c)
	var input models.SignDocuments
	var signDocs models.SignDocs
	//Input Mapping
	err := c.ShouldBind(&input)
	if err != nil {
		log.Println(err)
		c.Redirect(302, "/sign-documents")
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
	}
	if file.Header.Get("Content-Type") != "application/pdf" || file.Filename[len(file.Filename)-4:] != ".pdf" {
		log.Println("File not pdf")
		c.Redirect(302, "/verification")
		return
	}
	input.Name = file.Filename
	//Saving Document to Directory
	path := fmt.Sprintf("./public/temp/pdfsign/%s", input.Name)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		log.Println(err)
	}
	//Generate hash document original
	input.Hash_original = h.signaturesService.GenerateHashDocument(path)
	//Get Address Creator
	input.Creator = fmt.Sprintf("%v", session.Get("public_key"))
	input.Creator_id = fmt.Sprintf("%v", session.Get("sign"))
	//Get Images signatures
	mysignatures, _ := h.signaturesService.GetMySignature(fmt.Sprintf("%v", session.Get("sign")), fmt.Sprintf("%v", session.Get("id")), fmt.Sprintf("%v", session.Get("name")))
	//Resize Images Signatures
	img := h.signaturesService.ResizeImages(mysignatures, input)
	//Signing Documents to PDF
	sign := h.signaturesService.SignDocuments(img, input)
	signDocs.Hash = h.signaturesService.GenerateHashDocument(sign)
	input.Hash = input.Hash_original
	//Input to IPFS
	err, IPFS := h.serviceUser.UploadIPFS(sign)
	if err != nil {
		log.Println(err)
	}
	//Delete file uploaded sign
	err = os.Remove(path)
	if err != nil {
		log.Println(err)
	}
	//Encript IPFS and Get Signatures Data
	input.IPFS = string(h.serviceUser.Encrypt([]byte(IPFS), conf.App.Secret_key))
	input.Address, input.IdSignature = h.serviceUser.GetPublicKey(input.Email)
	//Add Creator for signatures
	input.Address = append(input.Address, common.HexToAddress(input.Creator))
	input.IdSignature = append(input.IdSignature, input.Creator_id)
	//Input to blockchain
	err = h.signaturesService.AddToBlockhain(input)
	if err != nil {
		log.Println(err)
	}
	//Signing Creator in Documents
	signDocs.Hash_original = input.Hash_original
	signDocs.Creator = input.Creator
	signDocs.IPFS = input.IPFS
	h.signaturesService.DocumentSigned(signDocs)

	//invite people
	if input.Invite_sts { //Check invite or not
		fmt.Println("Sisa Invite People Via Email Habis itu selesai")
		//h.signaturesService.InvitePeople(input.Email) //Invite Via Email
	}
	input.Hash = signDocs.Hash
	err = h.signaturesService.AddUserDocs(input)
	if err != nil {
		log.Println(err)
	}
	//fmt.Println(input)
	c.Redirect(302, "/sign-documents")
}

func (h *signaturesHandler) InviteSignatures(c *gin.Context) {
	conf, _ := config.Init()
	session := sessions.Default(c)
	var input models.InviteSignatures
	var DocData models.SignDocuments
	//Input Mapping
	err := c.ShouldBind(&input)
	if err != nil {
		log.Println(err)
		c.Redirect(302, "/invite-signatures")
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
	}
	if file.Header.Get("Content-Type") != "application/pdf" || file.Filename[len(file.Filename)-4:] != ".pdf" {
		log.Println("File not pdf")
		c.Redirect(302, "/verification")
		return
	}
	DocData.Name = file.Filename
	//Saving Image to Directory
	path := fmt.Sprintf("./public/temp/pdfsign/%s", DocData.Name)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		log.Println(err)
	}
	DocData.Email = input.Email
	DocData.Judul = input.Judul
	DocData.Note = input.Note
	//Generate hash document
	DocData.Hash_original = h.signaturesService.GenerateHashDocument(path)
	DocData.Hash = DocData.Hash_original
	//Get Address Creator
	DocData.Creator = fmt.Sprintf("%v", session.Get("public_key"))
	DocData.Creator_id = fmt.Sprintf("%v", session.Get("sign"))
	//Input to IPFS
	err, DocData.IPFS = h.serviceUser.UploadIPFS(path)
	if err != nil {
		log.Println(err)
	}
	//Encript IPFS and Get Signatures Data
	DocData.IPFS = string(h.serviceUser.Encrypt([]byte(DocData.IPFS), conf.App.Secret_key))
	DocData.Address, DocData.IdSignature = h.serviceUser.GetPublicKey(DocData.Email)
	//Input to blockchain
	err = h.signaturesService.AddToBlockhain(DocData)
	if err != nil {
		log.Println(err)
	}

	//invite people
	fmt.Println("Sisa Invite People Via Email Habis itu selesai")
	//h.signaturesService.InvitePeople(DocData.Email) //Invite Via Email

	err = h.signaturesService.AddUserDocs(DocData)
	if err != nil {
		log.Println(err)
	}
	c.Redirect(302, "/invite-signatures")
}

func (h *signaturesHandler) Document(c *gin.Context) {
	var input models.SignDocuments
	var signDocs models.SignDocs
	conf, _ := config.Init()
	session := sessions.Default(c)
	//Input Mapping
	err := c.ShouldBind(&input)
	if err != nil {
		log.Println(err)
	}
	input.Hash_original = c.Param("hash")
	input.Name = input.Hash_original + ".pdf"
	//Get Images signatures
	mysignatures, _ := h.signaturesService.GetMySignature(fmt.Sprintf("%v", session.Get("sign")), fmt.Sprintf("%v", session.Get("id")), fmt.Sprintf("%v", session.Get("name")))
	//Resize Images Signatures
	img := h.signaturesService.ResizeImages(mysignatures, input)
	//Signing Document to PDF
	signing := h.signaturesService.SignDocuments(img, input)
	//Generate Hash Document Signed and Upload to IPFS
	input.Hash = h.signaturesService.GenerateHashDocument(signing)
	err, input.IPFS = h.serviceUser.UploadIPFS(signing)
	if err != nil {
		log.Println(err)
	}
	input.IPFS = string(h.serviceUser.Encrypt([]byte(input.IPFS), conf.App.Secret_key))
	//Signing Documents in Blockchain
	signDocs.Hash_original = input.Hash_original
	signDocs.Creator = fmt.Sprintf("%v", session.Get("public_key"))
	signDocs.Hash = input.Hash
	signDocs.IPFS = input.IPFS
	h.signaturesService.DocumentSigned(signDocs)
	//Remove document
	inputPath := fmt.Sprintf("./public/temp/pdfsign/%s", input.Name)
	err = os.Remove(inputPath)
	if err != nil {
		log.Println(err)
	}

	c.Redirect(302, "/request-signatures")
}

func (h *signaturesHandler) Verification(c *gin.Context) {
	session := sessions.Default(c)
	title := "Hasil Verifikasi - SmartSign"
	page := "verification"
	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
	}
	//Filter file pdf
	if file.Header.Get("Content-Type") != "application/pdf" || file.Filename[len(file.Filename)-4:] != ".pdf" {
		log.Println("File not pdf")
		c.Redirect(302, "/verification")
		return
	}
	//Saving File to Directory
	path := fmt.Sprintf("./public/temp/pdfverify/%s", file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		log.Println(err)
	}
	//Generate Hash Document
	hash := h.signaturesService.GenerateHashDocument(path)
	//Get Data Document
	data, exist := h.signaturesService.GetDocumentAllSign(hash)
	if exist {
		fmt.Println(data)
	} else {
		log.Println("Document not signed")
	}
	c.HTML(http.StatusOK, "verification_result.html", gin.H{
		"title":       title,
		"userid":      session.Get("id"),
		"page":        page,
		"hash":        hash,
		"verif_state": exist,
		"data":        data,
	})
}

//Download Document Signed From Blockchain and IPFS to Client.
func (h *signaturesHandler) Download(c *gin.Context) {
	conf, _ := config.Init()
	hash := c.Param("hash")
	doc := h.signaturesService.GetDocumentNoSigners(hash)
	doc.IPFS = string(h.serviceUser.Decrypt([]byte(doc.IPFS), conf.App.Secret_key))
	//Download File From
	directory := "./public/temp/pdfdownload/"
	res, _ := h.serviceUser.GetFileIPFS(doc.IPFS, doc.Metadata+".pdf", directory)
	//Download Dokumen
	c.FileAttachment(res, doc.Metadata+".pdf")
	//Delete File
	err := os.Remove(res)
	if err != nil {
		log.Println(err)
	}
}

// Test verification
// func (h *signaturesHandler) Verif(c *gin.Context) {
// 	hash := c.Param("hash")
// 	data, exist := h.signaturesService.GetDocumentAllSign(hash)
// 	if !exist {
// 		log.Println("Document not signed")
// 		c.JSON(200, "Document not signed")
// 	} else {
// 		c.JSON(200, data)
// 	}
// }

//----- Get Documents and Signatures -----//
// func (h *signaturesHandler) GetDocs(c *gin.Context) {
// 	session := sessions.Default(c)
// 	hash := c.Param("hash")
// 	id := c.Param("id")
// 	if id == "" {
// 		id = fmt.Sprintf("%v", session.Get("id"))
// 	}
// 	fmt.Println(hash)
// 	docs := h.signaturesService.GetDocument(hash, id)
// 	fmt.Println(docs)
// 	c.JSON(200, docs)
// }
//----- End Get Documents and Signatures -----//
