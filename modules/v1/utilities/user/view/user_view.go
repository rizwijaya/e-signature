package view

import (
	api "e-signature/app/contracts"
	"e-signature/modules/v1/utilities/user/repository"
	"e-signature/modules/v1/utilities/user/service"
	error "e-signature/pkg/http-error"
	notif "e-signature/pkg/notification"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type userview struct {
	userService service.Service
}

func NewUserView(userService service.Service) *userview {
	return &userview{userService}
}

func View(db *mongo.Database, blockhain *api.Api, client *ethclient.Client) *userview {
	Repository := repository.NewRepository(db, blockhain, client)
	Service := service.NewService(Repository)
	return NewUserView(Service)
}

func (h *userview) Index(c *gin.Context) {
	session := sessions.Default(c)
	title := "SmartSign - Smart Digital Signatures"
	page := "index"

	fm := notif.GetMessage(c.Writer, c.Request, "message")
	c.HTML(http.StatusOK, "landing_index.html", gin.H{
		"title":   title,
		"page":    page,
		"success": string(fm),
		"userid":  session.Get("id"),
	})
}

func (h *userview) Dashboard(c *gin.Context) {
	session := sessions.Default(c)
	title := "Dashboard - SmartSign"
	page := "dashboard"

	fm := notif.GetMessage(c.Writer, c.Request, "message")
	cardDashboard := h.userService.GetCardDashboard(session.Get("sign").(string))
	//Logging Access
	h.userService.Logging("Mengakses Halaman Dashboard", session.Get("sign").(string), c.ClientIP(), c.Request.UserAgent())
	c.HTML(http.StatusOK, "dashboard_index.html", gin.H{
		"title":   title,
		"page":    page,
		"success": string(fm),
		"userid":  session.Get("id"),
		"card":    cardDashboard,
	})
}

func (h *userview) Register(c *gin.Context) {
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

func (h *userview) Login(c *gin.Context) {
	title := "Masuk - SmartSign"
	fm := notif.GetMessage(c.Writer, c.Request, "registered")
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
			"registered": string(fm),
		},
	)
}

func (h *userview) Logg(c *gin.Context) {
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
