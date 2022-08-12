package view

import (
	api "e-signature/app/contracts"
	"e-signature/modules/v1/utilities/user/repository"
	"e-signature/modules/v1/utilities/user/service"
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

func View(db *gorm.DB, blockhain *api.Api) *userView {
	Repository := repository.NewRepository(db)
	Service := service.NewService(Repository, blockhain)
	return NewUserView(Service)
}

func (h *userView) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "landing_index.html", nil)
}

func (h *userView) Dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard_index.html", nil)
}

func (h *userView) Register(c *gin.Context) {
	title := "Register - SmartSign"
	c.HTML(http.StatusOK, "register.html",
		gin.H{
			"title": title,
		},
	)

	//Algoritma Register
	//Register in record system
	//After Success register in record user
}

func (h *userView) Login(c *gin.Context) {
	title := "Login - SmartSign"
	c.HTML(http.StatusOK, "login.html",
		gin.H{
			"title": title,
		},
	)

	//Algoritma Login
	//Check idsignature, pw in record system
	//Valid, get all data
	// set session sementara with data
	//get data in record user with address
}
