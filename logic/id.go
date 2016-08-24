package logic

import (
	"github.com/RemmargorP/mjudge/models"

	"github.com/satori/go.uuid"
)

func GenerateId() models.Id {
	return models.Id(uuid.NewV4().String())
}
