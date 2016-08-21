package web

import (
	"github.com/RemmargorP/memjudge/models"
	"log"
	"net/http"
)

func (wi *WebInstance) HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["title"] = TITLE
	data["webinstanceid"] = wi.Id
	data["webinfo"] = map[string]interface{}{"isHome": true}

	user := models.GetUserFromCookie(r, wi.DB)

	models.GatherUserInfo(user, &data)

	wi.ParseCookieReason(w, r, &data)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err := wi.Templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		log.Println(err)
	}
}
