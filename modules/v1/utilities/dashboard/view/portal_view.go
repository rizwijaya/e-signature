package view

import (
	"TamaskaDashboard/modules/v1/utilities/dashboard/repository"
	"TamaskaDashboard/modules/v1/utilities/dashboard/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type dashboardView struct {
	dashboardService service.Service
}

func NewDashboardView(dashboardService service.Service) *dashboardView {
	return &dashboardView{dashboardService}
}

func View(db *gorm.DB) *dashboardView {
	Repository := repository.NewRepository(db)
	Service := service.NewService(Repository)
	return NewDashboardView(Service)
}

func (h *dashboardView) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard_index.html", nil)
}
