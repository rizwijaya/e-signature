package models

type AddSignature struct {
	Id        string `json:"unique" binding:"required"`
	Signature string `json:"signature" binding:"required"`
}
