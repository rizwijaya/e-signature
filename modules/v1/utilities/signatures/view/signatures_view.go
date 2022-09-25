package view

import (
	"e-signature/modules/v1/utilities/signatures/repository"
	"e-signature/modules/v1/utilities/signatures/service"
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

// func (h *signaturesView) AddSignature(c *gin.Context) {
// 	c.HTML(http.StatusOK, "add_signature.html", nil)
// }

// func (h *signaturesView) ListSignature(c *gin.Context) {
// 	c.HTML(http.StatusOK, "list_signature.html", nil)
// }

func (h *signaturesView) MySignatures(c *gin.Context) {
	session := sessions.Default(c)
	title := "My Signature - SmartSign"

	c.HTML(http.StatusOK, "my_signatures.html", gin.H{
		"title": title,
		"user":  session.Get("id"),
	})
}
