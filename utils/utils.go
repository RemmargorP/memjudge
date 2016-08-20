package utils

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func GetCookieValue(r *http.Request, name string) string {
	cookie, err := r.Cookie(name)
	if err != nil {
		return ""
	}
	return cookie.Value
}

func EraseCookie(w http.ResponseWriter, r *http.Request, name string) {
	cookie, err := r.Cookie(name)
	if err != nil {
		return
	}
	cookie.MaxAge = 0
	cookie.Value = ""
	http.SetCookie(w, cookie)
}

func ValidateEmail(e string) (bool, error) {
	return regexp.MatchString(`\S+@\S+\.\S+`, e)
}

var encryptLeftSalt, encryptRightSalt string

func Encrypt(s string) string {
	h := md5.New()
	io.WriteString(h, encryptLeftSalt)
	io.WriteString(h, s)
	io.WriteString(h, encryptRightSalt)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func init() {
	db_auth_json, err := ioutil.ReadFile("DB_auth.json")
	if err != nil {
		log.Fatal(err)
	}

	var db_auth struct {
		PWDLeftSalt, PWDRightSalt string
	}
	err = json.Unmarshal(db_auth_json, &db_auth)

	encryptLeftSalt = db_auth.PWDLeftSalt
	encryptRightSalt = db_auth.PWDRightSalt
}
