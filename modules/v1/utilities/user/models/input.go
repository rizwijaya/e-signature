package models

type LoginInput struct {
	IdSignature string `json:"idsignature" form:"idsignature" binding:"required" label:"id signature"`
	Password    string `json:"password" form:"password" binding:"required" label:"kata sandi"`
}

type RegisterUserInput struct {
	IdSignature string `json:"idsignature" form:"idsignature" binding:"required,alphanum,min=6" label:"id signature"`
	Name        string `json:"name" form:"name" binding:"required" label:"nama"`
	Email       string `json:"email" form:"email" binding:"required,email" label:"email"`
	Phone       string `json:"phone" form:"phone" binding:"required,numeric" label:"nomor hp"`
	Password    string `json:"password" form:"password" binding:"required,alphanum,min=8,eqfield=CPassword" label:"kata sandi"`
	CPassword   string `json:"cpassword" form:"cpassword" binding:"required,alphanum,min=8" label:"konfirmasi kata sandi"`
}
