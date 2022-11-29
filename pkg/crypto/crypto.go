package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
	"io/ioutil"
	"log"

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

func CreateKey(key string) []byte {
	hash := sha256.Sum256([]byte(key))
	return hash[:]
}

func Encrypt(data []byte, passphrase string) []byte {
	key := CreateKey(passphrase)
	block, _ := aes.NewCipher(key)
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Println(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Println(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)

	base64Cipher := make([]byte, base64.RawStdEncoding.EncodedLen(len(ciphertext)))
	base64.RawStdEncoding.Encode(base64Cipher, ciphertext)
	//return string(base64Cipher)
	return base64Cipher
}

func Decrypt(data []byte, passphrase string) []byte {
	cipherText := make([]byte, base64.RawStdEncoding.DecodedLen(len(data)))
	_, err := base64.RawStdEncoding.Decode(cipherText, data)
	if err != nil {
		return nil
	}
	key := CreateKey(passphrase)
	//key := "791f13d3e2552bcf31c4f8d0e5d6a1ed"
	//fmt.Println("key", key)
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Println(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := cipherText[:nonceSize], cipherText[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Println(err.Error())
	}
	return plaintext
}

func EncryptFile(filename string, passphrase string) error {
	b, err := ioutil.ReadFile(filename) //Read the target file
	if err != nil {
		log.Println("Unable to open the input file!")
		return err
	}
	ciphertext := Encrypt(b, passphrase)
	err = ioutil.WriteFile(filename, ciphertext, 0644)
	if err != nil {
		log.Println("Unable to create encrypted file!")
		return err
	}
	//fmt.Println(ciphertext)
	return nil
}

func DecryptFile(filename string, passphrase string) error {
	z, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("Unable to open the input file!")
		return err
	}
	result := Decrypt(z, passphrase)
	//fmt.Printf("Decrypted file was created with file permissions 0777\n")
	err = ioutil.WriteFile(filename, result, 0777)
	if err != nil {
		log.Println("Unable to create decrypted file!")
		return err
	}
	return nil
}
