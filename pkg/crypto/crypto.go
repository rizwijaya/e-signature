package crypto

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(pw string) (string, error) {
	if len(pw) < 6 {
		return "", errors.New("password must be more than 6 characters")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func Compare(hash string, pw string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
	if err != nil {
		return errors.New("password salah")
	}
	return nil
}
