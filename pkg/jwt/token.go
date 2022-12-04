package token

import (
	"e-signature/app/config"
	"errors"

	"github.com/dgrijalva/jwt-go"
)

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	conf, _ := config.Init()
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid token")
		}

		return []byte(conf.App.Secret_key), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}

func GenerateToken(email string, pw string) (string, error) {
	conf, _ := config.Init()
	SECRET_KEY := []byte(conf.App.Secret_key)
	claim := jwt.MapClaims{}
	claim["email"] = email
	claim["pw"] = pw

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
