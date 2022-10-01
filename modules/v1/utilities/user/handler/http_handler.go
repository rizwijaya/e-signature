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
	PublicKey := h.userService.Decrypt([]byte(user.PublicKey), input.Password)
	user.PublicKey = string(PublicKey)
	//------ Enabled in Production karena melakukan transfer balance ----------------
	// //Check Balance Accounts
	// mybalance, _ := h.userService.GetBalance(user, input.Password)
	// //fmt.Println(mybalance)
	// if mybalance == "0" { //transfer balance if balance is 0
	// 	err := h.userService.TransferBalance(user)
	// 	if err != nil {
	// 		c.HTML(http.StatusOK, "login.html", gin.H{
	// 			"title":   "Login - SmartSign",
	// 			"message": "Terjadi kesalahan, harap hubungi administrator.",
	// 		})
	// 		return
	// 	}
	// }
	//------ End Enabled in Production karena melakukan transfer balance ----------------
	session.Set("id", user.Id.Hex())
	session.Set("sign", user.Idsignature)
	session.Set("name", user.Name)
	session.Set("public_key", user.PublicKey)
	session.Set("role", user.Role_id)
	session.Set("passph", string(h.userService.Encrypt([]byte(input.Password), user.PublicKey)))
	session.Save()

	// fm := []byte("Berhasil Masuk, Selamat Datang " + user.Name)
	// notif.SetMessage(c.Writer, "success", fm)
	//fmt.Println(session.Get("id"))
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
	id, _ := h.userService.CheckUserExist(input.IdSignature)
	if id == "exist" {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title":   "Register - SmartSign",
			"message": "ID Signature sudah terdaftar.",
		})
		return
	}
	//Check if email already
	email, _ := h.userService.CheckEmailExist(input.Email)
	if email == "exist" {
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
	path := fmt.Sprintf("./public/images/identity_card/card-%s.%s", input.IdSignature, file.Filename[len(file.Filename)-3:])
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title":   "Register - SmartSign",
			"message": "Terjadi kesalahan, harap hubungi administrator.",
		})
		return
	}
	user.Identity_card = fmt.Sprintf("card-%s.%s", input.IdSignature, file.Filename[len(file.Filename)-3:])
	h.userService.EncryptFile(path, input.Password)

	//Create Account
	idn, err := h.userService.CreateAccount(user)
	if err != nil {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title":   "Register - SmartSign",
			"message": "Terjadi kesalahan, harap hubungi administrator.",
		})
		return
	}
	//fmt.Println(idn)
	//Create Default Latin Signatures
	latin := h.signatureService.CreateLatinSignatures(user, idn)
	h.signatureService.CreateLatinSignaturesData(user, latin, idn)
	//Save Signatures
	err = h.signatureService.DefaultSignatures(user, idn)
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
