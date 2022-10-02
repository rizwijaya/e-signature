package view

import (
	"e-signature/modules/v1/utilities/signatures/repository"
	"e-signature/modules/v1/utilities/signatures/service"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type signaturesView struct {
	signaturesService service.Service
}

func NewSignaturesView(signaturesService service.Service) *signaturesView {
	return &signaturesView{signaturesService}
}

func View(db *mongo.Database) *signaturesView {
	Repository := repository.NewRepository(db)
	Service := service.NewService(Repository)
	return NewSignaturesView(Service)
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

	c.HTML(http.StatusOK, "sign_documents.html", gin.H{
		"title": title,
		"user":  session.Get("id"),
	})
}
