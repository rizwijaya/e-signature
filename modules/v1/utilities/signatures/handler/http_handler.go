package signatures

import (
	"e-signature/app/config"
	"e-signature/modules/v1/utilities/signatures/models"
	api "e-signature/pkg/api_response"
	"fmt"
	"log"

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
	input.Name = file.Filename
	//Saving Image to Directory
	path := fmt.Sprintf("./public/temp/%s", input.Name)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		log.Println(err)
	}
	//Get Images signatures
	mysignatures, _ := h.signaturesService.GetMySignature(fmt.Sprintf("%v", session.Get("sign")), fmt.Sprintf("%v", session.Get("id")), fmt.Sprintf("%v", session.Get("name")))
	//Resize Images Signatures
	img := h.signaturesService.ResizeImages(mysignatures, input)
	//Signing Documents to PDF
	sign := h.signaturesService.SignDocuments(img, input)
	//Generate hash document
	input.Hash_original = h.signaturesService.GenerateHashDocument(path)
	input.Hash = h.signaturesService.GenerateHashDocument(sign)
	//Get Address Creator
	input.Creator = fmt.Sprintf("%v", session.Get("public_key"))
	input.Creator_id = fmt.Sprintf("%v", session.Get("sign"))
	//Input to IPFS
	err, input.IPFS = h.serviceUser.UploadIPFS(sign)
	if err != nil {
		log.Println(err)
	}
	//Encript IPFS and Get Signatures Data
	input.IPFS = string(h.serviceUser.Encrypt([]byte(input.IPFS), conf.App.Secret_key))
	input.Address, input.IdSignature = h.serviceUser.GetPublicKey(input.Email)
	//Add Creator for signatures
	input.Address = append(input.Address, common.HexToAddress(input.Creator))
	input.IdSignature = append(input.IdSignature, input.Creator)
	//Input to blockchain
	err = h.signaturesService.AddToBlockhain(input)
	if err != nil {
		log.Println(err)
	}
	//Signing Creator in Documents
	signDocs.Hash_original = input.Hash_original
	signDocs.Creator = input.Creator
	signDocs.Hash = input.Hash
	signDocs.IPFS = input.IPFS
	h.signaturesService.DocumentSigned(signDocs)

	//invite people
	if input.Invite_sts { //Check invite or not
		fmt.Println("Sisa Invite People Via Email Habis itu selesai")
		//h.signaturesService.InvitePeople(input.Email) //Invite Via Email
	}
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
	var signDocs models.SignDocs
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
	DocData.Name = file.Filename
	//Saving Image to Directory
	path := fmt.Sprintf("./public/temp/%s", DocData.Name)
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
	//Signing Creator in Documents
	signDocs.Hash_original = DocData.Hash_original
	signDocs.Creator = DocData.Creator
	signDocs.Hash = DocData.Hash
	signDocs.IPFS = DocData.IPFS
	h.signaturesService.DocumentSigned(signDocs)

	//invite people
	fmt.Println("Sisa Invite People Via Email Habis itu selesai")
	//h.signaturesService.InvitePeople(DocData.Email) //Invite Via Email

	err = h.signaturesService.AddUserDocs(DocData)
	if err != nil {
		log.Println(err)
	}
	c.Redirect(302, "/invite-signatures")
}
