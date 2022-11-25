package view

import (
	api "e-signature/app/contracts"
	"e-signature/modules/v1/utilities/user/repository"
	"e-signature/modules/v1/utilities/user/service"
	error "e-signature/pkg/http-error"
	notif "e-signature/pkg/notification"
	"fmt"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type userView struct {
	userService service.Service
}

func NewUserView(userService service.Service) *userView {
	return &userView{userService}
}

func View(db *mongo.Database, blockhain *api.Api, client *ethclient.Client) *userView {
	Repository := repository.NewRepository(db, blockhain, client)
	Service := service.NewService(Repository)
	return NewUserView(Service)
}

func (h *userView) Index(c *gin.Context) {
	session := sessions.Default(c)
	title := "SmartSign - Smart Digital Signatures"
	page := "index"

	fm, err := notif.GetMessage(c.Writer, c.Request, "message")
	if err != nil {
		log.Println(err)
	}
	c.HTML(http.StatusOK, "landing_index.html", gin.H{
		"title":   title,
		"page":    page,
		"success": fmt.Sprintf("%s", fm),
		"userid":  session.Get("id"),
	})
}

func (h *userView) Dashboard(c *gin.Context) {
	session := sessions.Default(c)
	title := "Dashboard - SmartSign"
	page := "dashboard"

	fm, err := notif.GetMessage(c.Writer, c.Request, "message")
	if err != nil {
		log.Println(err)
	}
	cardDashboard := h.userService.GetCardDashboard(session.Get("sign").(string))
	//Logging Access
	h.userService.Logging("Mengakses Halaman Dashboard", session.Get("sign").(string), c.ClientIP(), c.Request.UserAgent())
	c.HTML(http.StatusOK, "dashboard_index.html", gin.H{
		"title":   title,
		"page":    page,
		"success": fmt.Sprintf("%s", fm),
		"userid":  session.Get("id"),
		"card":    cardDashboard,
	})
}

func (h *userView) Register(c *gin.Context) {
	title := "Pendaftaran - SmartSign"
	out := []error.Form{
		{
			Field:   "no field",
			Message: "invalid input",
		},
	}

	c.HTML(http.StatusOK, "register.html",
		gin.H{
			"title":    title,
			"errorVal": out,
		},
	)
}

func (h *userView) Login(c *gin.Context) {
	title := "Masuk - SmartSign"
	fm, _ := notif.GetMessage(c.Writer, c.Request, "registered")
	out := []error.Form{
		{
			Field:   "no field",
			Message: "invalid input",
		},
	}
	if fm == nil {
		c.HTML(http.StatusOK, "login.html",
			gin.H{
				"title":    title,
				"errorVal": out,
			},
		)
		return
	}
	c.HTML(http.StatusOK, "login.html",
		gin.H{
			"title":      title,
			"errorVal":   out,
			"registered": fmt.Sprintf("%s", fm),
		},
	)
}

func (h *userView) Logg(c *gin.Context) {
	session := sessions.Default(c)
	title := "Log Akses User - SmartSign"
	page := "log-user"
	logg, _ := h.userService.GetLogUser(session.Get("sign").(string))

	h.userService.Logging("Mengakses halaman log akses user", session.Get("sign").(string), c.ClientIP(), c.Request.UserAgent())
	c.HTML(http.StatusOK, "log.html", gin.H{
		"title":  title,
		"userid": session.Get("id"),
		"page":   page,
		"name":   session.Get("name"),
		"log":    logg,
	})
}
