package logic

import (
	"github.com/RemmargorP/mjudge/models"
)

func CanRead(r *models.Rights) bool {
	return r.Rights >= models.R_Read
}

func CanWrite(r *models.Rights) bool {
	return r.Rights >= models.R_Write
}

func CanSetRights(r *models.Rights) bool {
	return r.Rights >= models.R_SetRights
}
