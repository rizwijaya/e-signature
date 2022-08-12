package models

type LoginInput struct {
	IdSignature string `json:"idsignature" form:"idsignature" binding:"required"`
	Password    string `json:"password" form:"password" binding:"required"`
}
