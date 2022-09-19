package view

import (
	api "e-signature/app/contracts"
	"e-signature/modules/v1/utilities/user/repository"
	"e-signature/modules/v1/utilities/user/service"
	notif "e-signature/pkg/notification"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userView struct {
	userService service.Service
}

func NewUserView(userService service.Service) *userView {
	return &userView{userService}
}

func View(db *gorm.DB, blockhain *api.Api, client *ethclient.Client) *userView {
	Repository := repository.NewRepository(db, blockhain, client)
	Service := service.NewService(Repository)
	return NewUserView(Service)
}

func (h *userView) Index(c *gin.Context) {
	session := sessions.Default(c)
	title := "SmartSign - Smart Digital Signatures"
	page := "index"

	if session.Get("id") == nil {
		c.HTML(http.StatusOK, "landing_index.html", gin.H{
			"title":  title,
			"page":   page,
			"userid": 0,
		})
	} else {
		c.HTML(http.StatusOK, "landing_index.html", gin.H{
			"title":  title,
			"page":   page,
			"userid": session.Get("id"),
		})
	}
}

func (h *userView) Dashboard(c *gin.Context) {
	session := sessions.Default(c)
	title := "Dashboard - SmartSign"
	page := "dashboard"
	//fm, _ := notif.GetMessage(c.Writer, c.Request, "success")
	// if fm == nil {
	// 	c.HTML(http.StatusOK, "dashboard_index.html",
	// 		gin.H{
	// 			"title": title,
	// 			"page":  page,
	// 		})
	// }
	c.HTML(http.StatusOK, "dashboard_index.html", gin.H{
		"title":  title,
		"page":   page,
		"userid": session.Get("id"),
		//"success": fmt.Sprintf("%s", fm),
	})
}

func (h *userView) Register(c *gin.Context) {
	title := "Pendaftaran - SmartSign"
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
	title := "Masuk - SmartSign"
	fm, _ := notif.GetMessage(c.Writer, c.Request, "registered")
	if fm == nil {
		c.HTML(http.StatusOK, "login.html",
			gin.H{
				"title": title,
			},
		)
		return
	}
	c.HTML(http.StatusOK, "login.html",
		gin.H{
			"title":      title,
			"registered": fmt.Sprintf("%s", fm),
		},
	)
}
