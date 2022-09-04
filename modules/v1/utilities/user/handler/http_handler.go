package user

import (
	"e-signature/modules/v1/utilities/user/models"
	"fmt"
	"log"
	"net/http"
	"time"

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

func (h *userHandler) Register(c *gin.Context) {
	var input models.RegisterUserInput
	var user models.User

	err := c.ShouldBind(&input)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title": "Register - SmartSign",
		})
		return
	}
	//Check Password and Confirm Password is same
	if input.Password != input.CPassword {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title": "Register - SmartSign",
			"flash": "Password and Confirm Password not match",
		})
		return
	}
	//Check if user already exist
	id, err := h.userService.CheckUserExist(input.IdSignature)
	if err != nil || id == "exist" {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title": "Register - SmartSign",
			"flash": "User already exist",
		})
		return
	}
	//Check if email already
	email, err := h.userService.CheckEmailExist(input.Email)
	if err != nil || email == "exist" {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title": "Register - SmartSign",
			"flash": "User already exist",
		})
		return
	}

	user.Idsignature = input.IdSignature
	user.Name = input.Name
	user.Password = input.Password
	user.Role = 2 //Role User
	user.Email = input.Email
	user.Phone = input.Phone
	user.Dateregistered = time.Now().String()

	//Binding File
	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title": "Register - SmartSign",
		})
		return
	}
	//Saving Image to Directory
	path := fmt.Sprintf("./public/images/%s-%s", input.IdSignature, file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title": "Register - SmartSign",
		})
		return
	}
	//Save File to to IPFS
	user.ImageIPFS, err = h.userService.SaveImage(input, file)
	if err != nil {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title": "Register - SmartSign",
		})
		return
	}
	//Create Account
	err = h.userService.CreateAccount(user)
	if err != nil {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title": "Register - SmartSign",
		})
		return
	}

	http.Redirect(c.Writer, c.Request, "/login", http.StatusFound)
}
