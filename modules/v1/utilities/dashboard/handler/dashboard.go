package dashboard

import (
	"TamaskaDashboard/modules/v1/utilities/dashboard/repository"
	"TamaskaDashboard/modules/v1/utilities/dashboard/service"

	"gorm.io/gorm"
)

type DashboardHandler interface {
}

type dashboardHandler struct {
	dashboardService service.Service
}

func NewDashboardHandler(dashboardService service.Service) *dashboardHandler {
	return &dashboardHandler{dashboardService}
}

func Handler(db *gorm.DB) *dashboardHandler {
	Repository := repository.NewRepository(db)
	Service := service.NewService(Repository)
	return NewDashboardHandler(Service)
}
