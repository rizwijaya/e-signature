package notify

import (
	"encoding/base64"
	"log"
	"net/http"
	"time"
)

func SetMessage(w http.ResponseWriter, name string, value []byte) {
	c := &http.Cookie{Name: name, Value: encode(value), MaxAge: time.Now().Minute() * 1}
	http.SetCookie(w, c)
}

func GetMessage(w http.ResponseWriter, r *http.Request, name string) []byte {
	c, err := r.Cookie(name)
	if err != nil {
		switch err {
		case http.ErrNoCookie:
			return nil
		default:
			log.Println(err)
			return nil
		}
	}
	value, err := decode(c.Value)
	if err != nil {
		log.Println(err)
		return nil
	}
	dc := &http.Cookie{Name: name}
	http.SetCookie(w, dc)
	return value
}

func encode(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
}

func decode(src string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(src)
}
