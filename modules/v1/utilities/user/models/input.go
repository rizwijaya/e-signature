package models

type LoginInput struct {
	IdSignature string `json:"idsignature" form:"idsignature" binding:"required,alphanum"`
	Password    string `json:"password" form:"password" binding:"required,alphanum"`
}

type RegisterUserInput struct {
	IdSignature string `json:"idsignature" form:"idsignature" binding:"required,alphanum,min=6"`
	Name        string `json:"name" form:"name" binding:"required"`
	Email       string `json:"email" form:"email" binding:"required,email"`
	Phone       string `json:"phone" form:"phone" binding:"required,numeric"`
	Password    string `json:"password" form:"password" binding:"required,alphanum,min=8,eqfield=CPassword"`
	CPassword   string `json:"cpassword" form:"cpassword" binding:"required,alphanum,min=8"`
}
