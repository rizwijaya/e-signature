package view

import (
	"TamaskaDashboard/modules/v1/utilities/user/repository"
	"TamaskaDashboard/modules/v1/utilities/user/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userView struct {
	userService service.Service
}

func NewUserView(userService service.Service) *userView {
	return &userView{userService}
}

func View(db *gorm.DB) *userView {
	Repository := repository.NewRepository(db)
	Service := service.NewService(Repository)
	return NewUserView(Service)
}

func (h *userView) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "landing_index.html", nil)
}

func (h *userView) Dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard_index.html", nil)
}

func (h *userView) Register(c *gin.Context) {
	c.HTML(http.StatusOK, "register_index.html", nil)
}

func (h *userView) Login(c *gin.Context) {
	title := "Login - SmartSign"
	c.HTML(http.StatusOK, "login_index.html",
		gin.H{
			"title": title,
		},
	)
}
