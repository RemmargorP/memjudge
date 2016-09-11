package web

import (
	"html/template"
	"net/http"

	"github.com/RemmargorP/mjudge/interfaces"
	"github.com/RemmargorP/mjudge/logic"
	"github.com/RemmargorP/mjudge/models"

	"github.com/go-mangos/mangos"
)

type WebInstance struct {
	Id        models.Id
	Repos     *interfaces.Repos
	Templates *template.Template
}
