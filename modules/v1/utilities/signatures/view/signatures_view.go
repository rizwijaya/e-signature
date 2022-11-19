package view

import (
	"e-signature/app/config"
	api "e-signature/app/contracts"
	"e-signature/modules/v1/utilities/signatures/repository"
	"e-signature/modules/v1/utilities/signatures/service"
	repoUser "e-signature/modules/v1/utilities/user/repository"
	serviceUser "e-signature/modules/v1/utilities/user/service"
	"fmt"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type signaturesView struct {
	signaturesService service.Service
	serviceUser       serviceUser.Service
}

func NewSignaturesView(signaturesService service.Service, userService serviceUser.Service) *signaturesView {
	return &signaturesView{signaturesService, userService}
}

func View(db *mongo.Database, blockhain *api.Api, client *ethclient.Client) *signaturesView {
	//Signatures
	Repository := repository.NewRepository(db, blockhain, client)
	Service := service.NewService(Repository)
	//User
	RepoUser := repoUser.NewRepository(db, blockhain, client)
	serviceUser := serviceUser.NewService(RepoUser)
	return NewSignaturesView(Service, serviceUser)
}

func (h *signaturesView) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard_index.html", nil)
}

func (h *signaturesView) MySignatures(c *gin.Context) {
	session := sessions.Default(c)
	title := "My Signature - SmartSign"
	signatures, _ := h.signaturesService.GetMySignature(fmt.Sprintf("%v", session.Get("sign")), fmt.Sprintf("%v", session.Get("id")), fmt.Sprintf("%v", session.Get("name")))
	// sign := []models.Signatures{
	// 	signatures,
	// }

	c.HTML(http.StatusOK, "my_signatures.html", gin.H{
		"title":      title,
		"user":       session.Get("id"),
		"signatures": signatures,
	})
}

func (h *signaturesView) SignDocuments(c *gin.Context) {
	session := sessions.Default(c)
	title := "Sign Documents - SmartSign"
	getSignature, err := h.signaturesService.GetMySignature(fmt.Sprintf("%v", session.Get("sign")), fmt.Sprintf("%v", session.Get("id")), fmt.Sprintf("%v", session.Get("name")))
	if err != nil {
		log.Println(err)
	}
	c.HTML(http.StatusOK, "sign_documents.html", gin.H{
		"title":      title,
		"user":       session.Get("id"),
		"signatures": getSignature,
	})
}

func (h *signaturesView) InviteSignatures(c *gin.Context) {
	session := sessions.Default(c)
	title := "Invite Signatures - SmartSign"
	c.HTML(http.StatusOK, "invite_signatures.html", gin.H{
		"title": title,
		"user":  session.Get("id"),
	})
}

func (h *signaturesView) RequestSignatures(c *gin.Context) {
	session := sessions.Default(c)
	title := "Signature Request - SmartSign"
	listDocument := h.signaturesService.GetListDocument(fmt.Sprintf("%v", session.Get("public_key")))

	c.HTML(http.StatusOK, "request_signatures.html", gin.H{
		"title":     title,
		"user":      session.Get("id"),
		"name":      session.Get("name"),
		"documents": listDocument,
	})
}

func (h *signaturesView) Document(c *gin.Context) {
	conf, _ := config.Init()
	session := sessions.Default(c)
	hash := c.Param("hash")
	title := "Signature Document - SmartSign"
	getSignature, err := h.signaturesService.GetMySignature(fmt.Sprintf("%v", session.Get("sign")), fmt.Sprintf("%v", session.Get("id")), fmt.Sprintf("%v", session.Get("name")))
	if err != nil {
		log.Println(err)
		c.Redirect(http.StatusMovedPermanently, "/request-signatures")
	}
	getDocument := h.signaturesService.GetDocument(hash, fmt.Sprintf("%v", session.Get("public_key")))
	getDocIPFS := string(h.serviceUser.Decrypt([]byte(getDocument.IPFS), conf.App.Secret_key))
	directory := "./public/temp/pdfsign/"
	_, err = h.serviceUser.GetFileIPFS(getDocIPFS, getDocument.Hash_ori+".pdf", directory)
	if err != nil {
		log.Println(err)
		c.Redirect(http.StatusMovedPermanently, "/request-signatures")
	}
	//fmt.Println(dir)
	c.HTML(http.StatusOK, "document.html", gin.H{
		"title":      title,
		"user":       session.Get("id"),
		"signatures": getSignature,
		"hash":       hash,
		"file":       getDocument.Hash_ori + ".pdf",
	})
}

func (h *signaturesView) Verification(c *gin.Context) {
	session := sessions.Default(c)
	title := "Verification - SmartSign"
	c.HTML(http.StatusOK, "verification.html", gin.H{
		"title": title,
		"user":  session.Get("id"),
	})
}

func (h *signaturesView) History(c *gin.Context) {
	session := sessions.Default(c)
	title := "History Signatures - SmartSign"
	listDocument := h.signaturesService.GetListDocument(fmt.Sprintf("%v", session.Get("public_key")))

	c.HTML(http.StatusOK, "history.html", gin.H{
		"title":     title,
		"user":      session.Get("id"),
		"name":      session.Get("name"),
		"documents": listDocument,
	})
}

func (h *signaturesView) Transactions(c *gin.Context) {
	session := sessions.Default(c)
	title := "Transactions - SmartSign"
	transac := h.signaturesService.GetTransactions()

	c.HTML(http.StatusOK, "transactions.html", gin.H{
		"title":   title,
		"user":    session.Get("id"),
		"transac": transac,
	})
}

// func (h *signaturesView) VerificationResult(c *gin.Context) {
// 	session := sessions.Default(c)
// 	title := "Verification - SmartSign"
// 	c.HTML(http.StatusOK, "verification_result.html", gin.H{
// 		"title": title,
// 		"user":  session.Get("id"),
// 		// "hash":        hash,
// 		// "verif_state": exist,
// 		// "data":        data,
// 	})
// }
