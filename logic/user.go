package logic

import (
	"crypto/md5"
	"fmt"
	"io"

	"github.com/RemmargorP/mjudge/models"
)

func IsAdministrator(u *models.User) bool {
	if u == nil {
		return false
	}
	return u.Administrator
}

func Encrypt(s, leftSalt, rightSalt string) string {
	h := md5.New()
	io.WriteString(h, leftSalt)
	io.WriteString(h, s)
	io.WriteString(h, rightSalt)
	return fmt.Sprintf("%x", h.Sum(nil))
}
