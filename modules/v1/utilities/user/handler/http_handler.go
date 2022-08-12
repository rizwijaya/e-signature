package user

import (
	"e-signature/modules/v1/utilities/user/models"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (h *userHandler) Login(c *gin.Context) {
	session := sessions.Default(c)
	var input models.LoginInput
	err := c.ShouldBind(&input)
	if err != nil {
		log.Println(err)
		//session.AddFlash("Invalid ID Signature Format")
		c.HTML(http.StatusOK, "login", gin.H{
			"title": "Login - SmartSign",
			"flash": session.Flashes(),
		})
		return
	}

	user, err := h.userService.Login(input)
	if err != nil {
		log.Println(err)
		//session.AddFlash("Invalid ID Signature or Password")
		c.HTML(http.StatusOK, "login", gin.H{
			"title": "Login - SmartSign",
			//"flash": session.Flashes(),
		})
		return
	}
	user.Password = ""
}
