package models

type AddSignature struct {
	Signature string `json:"signature" binding:"required"`
}
