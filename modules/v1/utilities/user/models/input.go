package models

type LoginInput struct {
	IdSignature string `json:"idsignature" form:"idsignature" binding:"required"`
	Password    string `json:"password" form:"password" binding:"required"`
}

type RegisterUserInput struct {
	IdSignature string `json:"idsignature" form:"idsignature" binding:"required"`
	Name        string `json:"name" form:"name" binding:"required"`
	Email       string `json:"email" form:"email" binding:"required,email"`
	Phone       string `json:"phone" form:"phone" binding:"required"`
	Password    string `json:"password" form:"password" binding:"required"`
	CPassword   string `json:"cpassword" form:"cpassword" binding:"required"`
}
