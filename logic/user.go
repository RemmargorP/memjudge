package logic

import (
	"github.com/RemmargorP/mjudge/models"
)

func IsAdministrator(u *models.User) bool {
	if u == nil {
		return false
	}
	return u.Administrator
}
