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
	session := sessions.Default(c)
	title := "Dashboard - SmartSign"
	page := "dashboard"
	c.HTML(http.StatusOK, "dashboard_index.html", gin.H{
		"title":  title,
		"userid": session.Get("id"),
		"page":   page,
	})
}

func (h *signaturesView) MySignatures(c *gin.Context) {
	session := sessions.Default(c)
	title := "My Signature - SmartSign"
	page := "my-signatures"
	signatures, _ := h.signaturesService.GetMySignature(fmt.Sprintf("%v", session.Get("sign")), fmt.Sprintf("%v", session.Get("id")), fmt.Sprintf("%v", session.Get("name")))

	c.HTML(http.StatusOK, "my_signatures.html", gin.H{
		"title":      title,
		"userid":     session.Get("id"),
		"page":       page,
		"signatures": signatures,
	})
}

func (h *signaturesView) SignDocuments(c *gin.Context) {
	session := sessions.Default(c)
	title := "Sign Documents - SmartSign"
	page := "sign-documents"
	getSignature, err := h.signaturesService.GetMySignature(fmt.Sprintf("%v", session.Get("sign")), fmt.Sprintf("%v", session.Get("id")), fmt.Sprintf("%v", session.Get("name")))
	if err != nil {
		log.Println(err)
	}
	c.HTML(http.StatusOK, "sign_documents.html", gin.H{
		"title":      title,
		"userid":     session.Get("id"),
		"page":       page,
		"signatures": getSignature,
	})
}

func (h *signaturesView) InviteSignatures(c *gin.Context) {
	session := sessions.Default(c)
	title := "Invite Signatures - SmartSign"
	page := "invite-signatures"
	c.HTML(http.StatusOK, "invite_signatures.html", gin.H{
		"title":  title,
		"userid": session.Get("id"),
		"page":   page,
	})
}

func (h *signaturesView) RequestSignatures(c *gin.Context) {
	session := sessions.Default(c)
	title := "Signature Request - SmartSign"
	page := "request-signatures"
	listDocument := h.signaturesService.GetListDocument(fmt.Sprintf("%v", session.Get("public_key")))

	c.HTML(http.StatusOK, "request_signatures.html", gin.H{
		"title":     title,
		"userid":    session.Get("id"),
		"page":      page,
		"name":      session.Get("name"),
		"documents": listDocument,
	})
}

func (h *signaturesView) Document(c *gin.Context) {
	conf, _ := config.Init()
	session := sessions.Default(c)
	hash := c.Param("hash")
	title := "Signature Document - SmartSign"
	page := "document"
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
	c.HTML(http.StatusOK, "document.html", gin.H{
		"title":      title,
		"userid":     session.Get("id"),
		"page":       page,
		"signatures": getSignature,
		"hash":       hash,
		"file":       getDocument.Hash_ori + ".pdf",
	})
}

func (h *signaturesView) Verification(c *gin.Context) {
	session := sessions.Default(c)
	title := "Verification - SmartSign"
	page := "verification"
	c.HTML(http.StatusOK, "verification.html", gin.H{
		"title":  title,
		"userid": session.Get("id"),
		"page":   page,
	})
}

func (h *signaturesView) History(c *gin.Context) {
	session := sessions.Default(c)
	title := "History Signatures - SmartSign"
	page := "history"
	listDocument := h.signaturesService.GetListDocument(fmt.Sprintf("%v", session.Get("public_key")))

	c.HTML(http.StatusOK, "history.html", gin.H{
		"title":     title,
		"userid":    session.Get("id"),
		"name":      session.Get("name"),
		"documents": listDocument,
		"page":      page,
	})
}

func (h *signaturesView) Transactions(c *gin.Context) {
	session := sessions.Default(c)
	title := "Transactions - SmartSign"
	page := "transactions"
	transac := h.signaturesService.GetTransactions()

	c.HTML(http.StatusOK, "transactions.html", gin.H{
		"title":   title,
		"userid":  session.Get("id"),
		"transac": transac,
		"page":    page,
	})
}

func (h *signaturesView) Download(c *gin.Context) {
	session := sessions.Default(c)
	title := "Daftar Unduh Dokumen - SmartSign"
	page := "history"
	listDocument := h.signaturesService.GetListDocument(fmt.Sprintf("%v", session.Get("public_key")))

	c.HTML(http.StatusOK, "download.html", gin.H{
		"title":     title,
		"userid":    session.Get("id"),
		"name":      session.Get("name"),
		"documents": listDocument,
		"page":      page,
	})
}
