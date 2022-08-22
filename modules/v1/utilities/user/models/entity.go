package models

type User struct {
	Profile_id     int
	Idsignature    string
	Name           string
	Password       string
	PasswordHash   string
	ImageIPFS      string
	Role           int
	Publickey      string
	Identity_card  string
	Email          string
	Phone          string
	Dateregistered string
}

type ProfileDB struct {
	Idsignature   string
	Name          string
	Email         string
	Phone         string
	Identity_card string
	Password      string
	PublicKey     string `gorm:"column:publickey"`
	Role_id       int
}
