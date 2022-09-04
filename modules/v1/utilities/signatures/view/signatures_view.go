package view

import (
	"e-signature/modules/v1/utilities/signatures/repository"
	"e-signature/modules/v1/utilities/signatures/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type signaturesView struct {
	signaturesService service.Service
}

func NewSignaturesView(signaturesService service.Service) *signaturesView {
	return &signaturesView{signaturesService}
}

func View(db *gorm.DB) *signaturesView {
	Repository := repository.NewRepository(db)
	Service := service.NewService(Repository)
	return NewSignaturesView(Service)
}

func (h *signaturesView) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard_index.html", nil)
}
