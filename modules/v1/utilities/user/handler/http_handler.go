package user

import (
	"e-signature/modules/v1/utilities/user/models"
	notif "e-signature/pkg/notification"
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
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title":   "Login - SmartSign",
			"message": "ID Signature/Password salah!",
		})
		return
	}

	user, err := h.userService.Login(input)
	if err != nil {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title":   "Login - SmartSign",
			"message": "ID Signature/Password salah!",
		})
		return
	}

	// token, err := token.GenerateToken(user, input.Password)
	// if err != nil {
	// 	c.HTML(http.StatusOK, "login.html", gin.H{
	// 		"title":   "Login - SmartSign",
	// 		"message": "Kesalahan! Harap hubungi administrator.",
	// 	})
	// 	return
	// }

	// session.Set("session", token)
	session.Set("id", user.User_id)
	session.Set("public_key", user.PublicKey)
	session.Set("role", user.Role_id)
	session.Set("passph", input.Password)
	session.Save()

	// fm := []byte("Berhasil Masuk, Selamat Datang " + user.Name)
	// notif.SetMessage(c.Writer, "success", fm)

	c.Redirect(http.StatusFound, "/dashboard")
}

func (h *userHandler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	http.SetCookie(c.Writer, &http.Cookie{
		Name:   "message",
		MaxAge: -1,
	})

	c.Redirect(http.StatusFound, "/")
}

func (h *userHandler) Register(c *gin.Context) {
	var input models.RegisterUserInput
	var user models.User
	//Input Validation
	err := c.ShouldBind(&input)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title": "pendaftaran - SmartSign",
			//"message": err.Error(),
			"message": "Harap masukan input dengan benar.",
		})
		return
	}

	//Check if user already exist
	id, err := h.userService.CheckUserExist(input.IdSignature)
	if err != nil || id == "exist" {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title":   "Register - SmartSign",
			"message": "ID Signature sudah terdaftar.",
		})
		return
	}
	//Check if email already
	email, err := h.userService.CheckEmailExist(input.Email)
	if err != nil || email == "exist" {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title":   "Register - SmartSign",
			"message": "Email sudah terdaftar.",
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
			"title":   "Register - SmartSign",
			"message": "Terjadi kesalahan, harap hubungi administrator.",
		})
		return
	}
	//Saving Image to Directory
	path := fmt.Sprintf("./public/images/%s-%s", input.IdSignature, file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title":   "Register - SmartSign",
			"message": "Terjadi kesalahan, harap hubungi administrator.",
		})
		return
	}
	//Save File to to IPFS
	user.ImageIPFS, err = h.userService.SaveImage(input, file)
	if err != nil {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title":   "Register - SmartSign",
			"message": "Terjadi kesalahan, harap hubungi administrator.",
		})
		return
	}
	//Create Account
	err = h.userService.CreateAccount(user)
	if err != nil {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title":   "Register - SmartSign",
			"message": "Terjadi kesalahan, harap hubungi administrator.",
		})
		return
	}

	fm := []byte("Berhasil melakukan pendaftaran, silahkan login untuk melanjutkan.")
	notif.SetMessage(c.Writer, "registered", fm)

	http.Redirect(c.Writer, c.Request, "/login", http.StatusFound)
}
