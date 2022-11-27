package view

import (
	"e-signature/app/config"
	api "e-signature/app/contracts"
	"e-signature/modules/v1/utilities/signatures/repository"
	"e-signature/modules/v1/utilities/signatures/service"
	repoUser "e-signature/modules/v1/utilities/user/repository"
	serviceUser "e-signature/modules/v1/utilities/user/service"
	notif "e-signature/pkg/notification"
	"fmt"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type signaturesview struct {
	serviceSignature service.Service
	serviceUser      serviceUser.Service
}

func NewSignaturesView(serviceSignature service.Service, userService serviceUser.Service) *signaturesview {
	return &signaturesview{serviceSignature, userService}
}

func View(db *mongo.Database, blockhain *api.Api, client *ethclient.Client) *signaturesview {
	//Signatures
	Repository := repository.NewRepository(db, blockhain, client)
	Service := service.NewService(Repository)
	//User
	RepoUser := repoUser.NewRepository(db, blockhain, client)
	serviceUser := serviceUser.NewService(RepoUser)
	return NewSignaturesView(Service, serviceUser)
}

func (h *signaturesview) MySignatures(c *gin.Context) {
	session := sessions.Default(c)
	title := "Tanda Tangan Saya - SmartSign"
	page := "my-signatures"
	signatures := h.serviceSignature.GetMySignature(fmt.Sprintf("%v", session.Get("sign")), fmt.Sprintf("%v", session.Get("id")), fmt.Sprintf("%v", session.Get("name")))
	fm := notif.GetMessage(c.Writer, c.Request, "message")
	h.serviceUser.Logging("Mengakses tanda tangan saya", session.Get("sign").(string), c.ClientIP(), c.Request.UserAgent())
	c.HTML(http.StatusOK, "my_signatures.html", gin.H{
		"title":      title,
		"userid":     session.Get("id"),
		"page":       page,
		"success":    string(fm),
		"signatures": signatures,
	})
}

func (h *signaturesview) SignDocuments(c *gin.Context) {
	session := sessions.Default(c)
	title := "Tanda Tangan Dokumen - SmartSign"
	page := "sign-documents"
	getSignature := h.serviceSignature.GetMySignature(fmt.Sprintf("%v", session.Get("sign")), fmt.Sprintf("%v", session.Get("id")), fmt.Sprintf("%v", session.Get("name")))
	fm := notif.GetMessage(c.Writer, c.Request, "failed")
	succes := notif.GetMessage(c.Writer, c.Request, "success")

	h.serviceUser.Logging("Mengakses tanda tangan dan minta tanda tangan", session.Get("sign").(string), c.ClientIP(), c.Request.UserAgent())
	c.HTML(http.StatusOK, "sign_documents.html", gin.H{
		"title":      title,
		"userid":     session.Get("id"),
		"page":       page,
		"failed":     string(fm),
		"success":    string(succes),
		"signatures": getSignature,
	})
}

func (h *signaturesview) InviteSignatures(c *gin.Context) {
	session := sessions.Default(c)
	title := "Undang untuk Tanda tangan - SmartSign"
	page := "invite-signatures"
	failed := notif.GetMessage(c.Writer, c.Request, "failed")
	succes := notif.GetMessage(c.Writer, c.Request, "success")
	h.serviceUser.Logging("Mengakses undang orang lain untuk tanda tangan", session.Get("sign").(string), c.ClientIP(), c.Request.UserAgent())
	c.HTML(http.StatusOK, "invite_signatures.html", gin.H{
		"title":   title,
		"userid":  session.Get("id"),
		"page":    page,
		"failed":  string(failed),
		"success": string(succes),
	})
}

func (h *signaturesview) RequestSignatures(c *gin.Context) {
	session := sessions.Default(c)
	title := "Daftar Permintaan Tanda Tangan - SmartSign"
	page := "request-signatures"
	listDocument := h.serviceSignature.GetListDocument(fmt.Sprintf("%v", session.Get("public_key")))
	failed := notif.GetMessage(c.Writer, c.Request, "failed")
	succes := notif.GetMessage(c.Writer, c.Request, "success")
	h.serviceUser.Logging("Mengakses halaman permintaan tanda tangan", session.Get("sign").(string), c.ClientIP(), c.Request.UserAgent())
	c.HTML(http.StatusOK, "request_signatures.html", gin.H{
		"title":     title,
		"userid":    session.Get("id"),
		"page":      page,
		"name":      session.Get("name"),
		"documents": listDocument,
		"failed":    string(failed),
		"success":   string(succes),
	})
}

func (h *signaturesview) Document(c *gin.Context) {
	conf, _ := config.Init()
	session := sessions.Default(c)
	hash := c.Param("hash")
	title := "Tanda Tangan Dokumen - SmartSign"
	page := "document"
	publickey := session.Get("public_key").(string)
	check := h.serviceSignature.CheckSignature(hash, publickey)
	if !check {
		c.Redirect(http.StatusMovedPermanently, "/")
	}
	//Check creator if not sign access
	getCreator := h.serviceSignature.GetDocument(hash, publickey)
	if getCreator.Creator_string == fmt.Sprintf("%v", session.Get("public_key")) {
		c.Redirect(http.StatusMovedPermanently, "/")
	}
	getSignature := h.serviceSignature.GetMySignature(fmt.Sprintf("%v", session.Get("sign")), fmt.Sprintf("%v", session.Get("id")), fmt.Sprintf("%v", session.Get("name")))
	getDocument := h.serviceSignature.GetDocument(hash, fmt.Sprintf("%v", session.Get("public_key")))
	getDocIPFS := string(h.serviceUser.Decrypt([]byte(getDocument.IPFS), conf.App.Secret_key))
	directory := "./public/temp/pdfsign/"
	_, err := h.serviceUser.GetFileIPFS(getDocIPFS, getDocument.Hash_ori+".pdf", directory)
	if err != nil {
		log.Println(err)
		c.Redirect(http.StatusMovedPermanently, "/request-signatures")
	}
	h.serviceUser.Logging("Mengakses dokumen untuk ditanda tangani", session.Get("sign").(string), c.ClientIP(), c.Request.UserAgent())
	c.HTML(http.StatusOK, "document.html", gin.H{
		"title":      title,
		"userid":     session.Get("id"),
		"page":       page,
		"signatures": getSignature,
		"hash":       hash,
		"file":       getDocument.Hash_ori + ".pdf",
	})
}

func (h *signaturesview) Verification(c *gin.Context) {
	session := sessions.Default(c)
	title := "Verifikasi - SmartSign"
	page := "verification"
	failed := notif.GetMessage(c.Writer, c.Request, "failed")
	succes := notif.GetMessage(c.Writer, c.Request, "success")
	c.HTML(http.StatusOK, "verification.html", gin.H{
		"title":   title,
		"userid":  session.Get("id"),
		"page":    page,
		"failed":  string(failed),
		"success": string(succes),
	})
}

func (h *signaturesview) History(c *gin.Context) {
	session := sessions.Default(c)
	title := "Riwayat Tanda Tangan - SmartSign"
	page := "history"
	listDocument := h.serviceSignature.GetListDocument(fmt.Sprintf("%v", session.Get("public_key")))
	h.serviceUser.Logging("Mengakses halaman riwayat tanda tangan", session.Get("sign").(string), c.ClientIP(), c.Request.UserAgent())
	c.HTML(http.StatusOK, "history.html", gin.H{
		"title":     title,
		"userid":    session.Get("id"),
		"name":      session.Get("name"),
		"documents": listDocument,
		"page":      page,
	})
}

func (h *signaturesview) Transactions(c *gin.Context) {
	session := sessions.Default(c)
	title := "Transaksi - SmartSign"
	page := "transactions"
	transac := h.serviceSignature.GetTransactions()

	c.HTML(http.StatusOK, "transactions.html", gin.H{
		"title":   title,
		"userid":  session.Get("id"),
		"transac": transac,
		"page":    page,
	})
}

func (h *signaturesview) Download(c *gin.Context) {
	session := sessions.Default(c)
	title := "Daftar Unduh Dokumen - SmartSign"
	page := "download"
	listDocument := h.serviceSignature.GetListDocument(fmt.Sprintf("%v", session.Get("public_key")))
	failed := notif.GetMessage(c.Writer, c.Request, "failed")
	succes := notif.GetMessage(c.Writer, c.Request, "success")

	h.serviceUser.Logging("Mengakses halaman daftar unduh dokumen", session.Get("sign").(string), c.ClientIP(), c.Request.UserAgent())
	c.HTML(http.StatusOK, "download.html", gin.H{
		"title":     title,
		"userid":    session.Get("id"),
		"name":      session.Get("name"),
		"documents": listDocument,
		"page":      page,
		"failed":    string(failed),
		"success":   string(succes),
	})
}
