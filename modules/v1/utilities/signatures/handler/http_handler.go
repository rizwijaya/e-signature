package signatures

import (
	"e-signature/app/config"
	"e-signature/modules/v1/utilities/signatures/models"
	api "e-signature/pkg/api_response"
	notif "e-signature/pkg/notification"
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
	h.serviceSignature.CreateImgSignature(input)
	h.serviceSignature.CreateImgSignatureData(input, sign)
	//Update Database MySignatures
	h.serviceSignature.UpdateMySignatures(fmt.Sprintf("signatures-%s.png", input.Id), fmt.Sprintf("signaturesdata-%s.png", input.Id), sign)

	//Return Response API
	response := api.APIRespon("Success Add Signatures", 200, "success", nil)
	//Logging Access
	h.serviceUser.Logging("Menambakan tanda tangan baru", sessions.Get("sign").(string), c.ClientIP(), c.Request.UserAgent())
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
	h.serviceSignature.ChangeSignatures(signing, user)
	fm := []byte("Mengganti Tanda Tangan!")
	notif.SetMessage(c.Writer, "message", fm)
	//Logging Access
	h.serviceUser.Logging("Mengganti tanda tangan ke "+sign_type, session.Get("sign").(string), c.ClientIP(), c.Request.UserAgent())
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
		fm := []byte("melakukan tanda tangan")
		notif.SetMessage(c.Writer, "failed", fm)
		c.Redirect(302, "/sign-documents")
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		fm := []byte("melakukan tanda tangan")
		notif.SetMessage(c.Writer, "failed", fm)
		log.Println(err)
		c.Redirect(302, "/sign-documents")
		return
	}
	if file.Header.Get("Content-Type") != "application/pdf" || file.Filename[len(file.Filename)-4:] != ".pdf" {
		log.Println("File not pdf")
		fm := []byte("melakukan tanda tangan")
		notif.SetMessage(c.Writer, "failed", fm)
		c.Redirect(302, "/sign-documents")
		return
	}
	input.Name = file.Filename
	//Saving Document to Directory
	path := fmt.Sprintf("./public/temp/pdfsign/%s", input.Name)
	_ = c.SaveUploadedFile(file, path)
	//Generate hash document original
	input.Hash_original = h.serviceSignature.GenerateHashDocument(path)
	//Get Address Creator
	input.Creator = fmt.Sprintf("%v", session.Get("public_key"))
	input.Creator_id = fmt.Sprintf("%v", session.Get("sign"))
	//Get Images signatures
	mysignatures := h.serviceSignature.GetMySignature(fmt.Sprintf("%v", session.Get("sign")), fmt.Sprintf("%v", session.Get("id")), fmt.Sprintf("%v", session.Get("name")))
	//Resize Images Signatures
	img := h.serviceSignature.ResizeImages(mysignatures, input)
	//Signing Documents to PDF
	sign := h.serviceSignature.SignDocuments(img, input)
	signDocs.Hash = h.serviceSignature.GenerateHashDocument(sign)
	input.Hash = input.Hash_original
	//Input to IPFS
	IPFS, err := h.serviceUser.UploadIPFS(sign)
	if err != nil {
		fm := []byte("melakukan tanda tangan")
		notif.SetMessage(c.Writer, "failed", fm)
		log.Println(err)
		c.Redirect(302, "/sign-documents")
		return
	}
	//Delete file uploaded sign
	_ = os.Remove(path)
	//Encript IPFS and Get Signatures Data
	input.IPFS = string(h.serviceUser.Encrypt([]byte(IPFS), conf.App.Secret_key))
	if input.Invite_sts { //Check invite or not
		input.Address, input.IdSignature = h.serviceUser.GetPublicKey(input.Email)
	}
	//Add Creator for signatures
	input.Address = append(input.Address, common.HexToAddress(input.Creator))
	input.IdSignature = append(input.IdSignature, input.Creator_id)
	if input.Invite_sts { //Mode ttd with Invite
		input.Mode = "1"
	} else { //Mode ttd without Invite
		input.Mode = "3"
	}
	//Input to blockchain
	err = h.serviceSignature.AddToBlockhain(input)
	if err != nil {
		fm := []byte("melakukan tanda tangan")
		notif.SetMessage(c.Writer, "failed", fm)
		log.Println(err)
		c.Redirect(302, "/sign-documents")
		return
	}
	//Signing Creator in Documents
	signDocs.Hash_original = input.Hash_original
	signDocs.Creator = input.Creator
	signDocs.IPFS = input.IPFS
	h.serviceSignature.DocumentSigned(signDocs)

	//invite people
	if input.Invite_sts { //Check invite or not
		for _, email := range input.Email { //Invite Via Email
			if email != "" {
				users, _ := h.serviceUser.GetUserByEmail(email)
				h.serviceSignature.InvitePeople(email, input, users)
			}
		}
	}
	input.Hash = signDocs.Hash
	err = h.serviceSignature.AddUserDocs(input)
	if err != nil {
		fm := []byte("melakukan tanda tangan")
		notif.SetMessage(c.Writer, "failed", fm)
		log.Println(err)
		c.Redirect(302, "/sign-documents")
		return
	}
	fm := []byte("melakukan tanda tangan")
	notif.SetMessage(c.Writer, "success", fm)
	//fmt.Println(input)
	//Delete Image Sign Resize
	err = os.Remove(img)
	if err != nil {
		log.Println(err)
	}
	//Logging Access
	h.serviceUser.Logging("Menandatangani dokumen "+input.Name, session.Get("sign").(string), c.ClientIP(), c.Request.UserAgent())
	c.Redirect(302, "/download")
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
		fm := []byte("mengundang orang lain untuk tanda tangan")
		notif.SetMessage(c.Writer, "failed", fm)
		c.Redirect(302, "/invite-signatures")
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
		fm := []byte("mengundang orang lain untuk tanda tangan")
		notif.SetMessage(c.Writer, "failed", fm)
		c.Redirect(302, "/invite-signatures")
		return
	}
	if file.Header.Get("Content-Type") != "application/pdf" || file.Filename[len(file.Filename)-4:] != ".pdf" {
		log.Println("File not pdf")
		fm := []byte("mengundang orang lain untuk tanda tangan")
		notif.SetMessage(c.Writer, "failed", fm)
		c.Redirect(302, "/invite-signatures")
		return
	}
	DocData.Name = file.Filename
	//Saving Image to Directory
	path := fmt.Sprintf("./public/temp/pdfsign/%s", DocData.Name)
	_ = c.SaveUploadedFile(file, path)
	DocData.Email = input.Email
	DocData.Judul = input.Judul
	DocData.Note = input.Note
	//Generate hash document
	DocData.Hash_original = h.serviceSignature.GenerateHashDocument(path)
	DocData.Hash = DocData.Hash_original
	//Get Address Creator
	DocData.Creator = fmt.Sprintf("%v", session.Get("public_key"))
	DocData.Creator_id = fmt.Sprintf("%v", session.Get("sign"))
	//Input to IPFS
	DocData.IPFS, err = h.serviceUser.UploadIPFS(path)
	if err != nil {
		log.Println(err)
		fm := []byte("mengundang orang lain untuk tanda tangan")
		notif.SetMessage(c.Writer, "failed", fm)
		c.Redirect(302, "/invite-signatures")
		return
	}
	//Encript IPFS and Get Signatures Data
	DocData.IPFS = string(h.serviceUser.Encrypt([]byte(DocData.IPFS), conf.App.Secret_key))
	DocData.Address, DocData.IdSignature = h.serviceUser.GetPublicKey(DocData.Email)
	DocData.Mode = "2"
	//Input to blockchain
	err = h.serviceSignature.AddToBlockhain(DocData)
	if err != nil {
		log.Println(err)
		fm := []byte("mengundang orang lain untuk tanda tangan")
		notif.SetMessage(c.Writer, "failed", fm)
		c.Redirect(302, "/invite-signatures")
		return
	}

	//Invite Via Email
	for _, email := range input.Email {
		if email != "" {
			users, _ := h.serviceUser.GetUserByEmail(email)
			h.serviceSignature.InvitePeople(email, DocData, users)
		}
	}
	//Add Creator for view signatures documents
	DocData.Address = append(DocData.Address, common.HexToAddress(DocData.Creator))
	DocData.IdSignature = append(DocData.IdSignature, DocData.Creator_id)

	err = h.serviceSignature.AddUserDocs(DocData)
	if err != nil {
		log.Println(err)
		fm := []byte("mengundang orang lain untuk tanda tangan")
		notif.SetMessage(c.Writer, "failed", fm)
		c.Redirect(302, "/invite-signatures")
		return
	}
	fm := []byte("mengundang orang lain untuk tanda tangan")
	notif.SetMessage(c.Writer, "success", fm)
	h.serviceUser.Logging("Mengundang orang lain untuk tanda tangan", session.Get("sign").(string), c.ClientIP(), c.Request.UserAgent())
	c.Redirect(302, "/download")
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
		fm := []byte("melakukan tanda tangan")
		notif.SetMessage(c.Writer, "failed", fm)
		c.Redirect(302, "/request-signatures")
		return
	}
	input.Hash_original = c.Param("hash")
	input.Name = input.Hash_original + ".pdf"
	//Get Images signatures
	mysignatures := h.serviceSignature.GetMySignature(fmt.Sprintf("%v", session.Get("sign")), fmt.Sprintf("%v", session.Get("id")), fmt.Sprintf("%v", session.Get("name")))
	//Resize Images Signatures
	img := h.serviceSignature.ResizeImages(mysignatures, input)
	//Signing Document to PDF
	signing := h.serviceSignature.SignDocuments(img, input)
	//Generate Hash Document Signed and Upload to IPFS
	input.Hash = h.serviceSignature.GenerateHashDocument(signing)
	input.IPFS, err = h.serviceUser.UploadIPFS(signing)
	if err != nil {
		log.Println(err)
		fm := []byte("melakukan tanda tangan")
		notif.SetMessage(c.Writer, "failed", fm)
		c.Redirect(302, "/request-signatures")
		return
	}
	input.IPFS = string(h.serviceUser.Encrypt([]byte(input.IPFS), conf.App.Secret_key))
	//Signing Documents in Blockchain
	signDocs.Hash_original = input.Hash_original
	signDocs.Creator = fmt.Sprintf("%v", session.Get("public_key"))
	signDocs.Hash = input.Hash
	signDocs.IPFS = input.IPFS
	h.serviceSignature.DocumentSigned(signDocs)
	//Remove document
	inputPath := fmt.Sprintf("./public/temp/pdfsign/%s", input.Name)
	_ = os.Remove(inputPath)
	//Delete Image Sign Resize
	_ = os.Remove(img)
	fm := []byte("melakukan tanda tangan")
	notif.SetMessage(c.Writer, "success", fm)
	h.serviceUser.Logging("Melakukan tanda tangan dari permintaan tanda tangan", session.Get("sign").(string), c.ClientIP(), c.Request.UserAgent())
	c.Redirect(302, "/download")
}

func (h *signaturesHandler) Verification(c *gin.Context) {
	session := sessions.Default(c)
	title := "Hasil Verifikasi - SmartSign"
	page := "verification"
	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
		fm := []byte("melakukan verifikasi")
		notif.SetMessage(c.Writer, "failed", fm)
		c.Redirect(302, "/verification")
		return
	}
	//Filter file pdf
	if file.Header.Get("Content-Type") != "application/pdf" || file.Filename[len(file.Filename)-4:] != ".pdf" {
		log.Println("File not pdf")
		fm := []byte("melakukan verifikasi")
		notif.SetMessage(c.Writer, "failed", fm)
		c.Redirect(302, "/verification")
		return
	}
	//Saving File to Directory
	path := fmt.Sprintf("./public/temp/pdfverify/%s", file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		log.Println(err)
		fm := []byte("melakukan verifikasi")
		notif.SetMessage(c.Writer, "failed", fm)
		c.Redirect(302, "/verification")
		return
	}
	//Generate Hash Document
	hash := h.serviceSignature.GenerateHashDocument(path)
	//Get Data Document
	data, exist := h.serviceSignature.GetDocumentAllSign(hash)
	if !exist {
		log.Println("Document not signed")
	}
	//remove Document
	err = os.Remove(path)
	if err != nil {
		log.Println(err)
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
	session := sessions.Default(c)
	hash := c.Param("hash")
	doc := h.serviceSignature.GetDocumentNoSigners(hash)
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
		failed := []byte("mengunduh dokumen")
		notif.SetMessage(c.Writer, "failed", failed)
		c.Redirect(302, "/download")
	}
	sucess := []byte("mengunduh dokumen")
	notif.SetMessage(c.Writer, "success", sucess)
	h.serviceUser.Logging("Mengunduh dokumen "+doc.Metadata+".pdf", session.Get("sign").(string), c.ClientIP(), c.Request.UserAgent())
	c.Redirect(302, "/download")
}

// Test verification
// func (h *signaturesHandler) Verif(c *gin.Context) {
// 	hash := c.Param("hash")
// 	data, exist := h.serviceSignature.GetDocumentAllSign(hash)
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
// 	docs := h.serviceSignature.GetDocument(hash, id)
// 	fmt.Println(docs)
// 	c.JSON(200, docs)
// }
//----- End Get Documents and Signatures -----//
