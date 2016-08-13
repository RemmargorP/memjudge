package memjudgeweb

import (
	models "github.com/RemmargorP/memjudge/models"
	"log"
	"net/http"
)

func (wi *WebInstance) HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["title"] = TITLE
	data["webinstanceid"] = wi.Id
	data["webinfo"] = map[string]interface{}{"isHome": true}

	session, err := wi.Store.Get(r, "session")
	if err != nil {
		session, err = wi.Store.New(r, "session")
		log.Println("Sessions fkt up: ", err)
	}

	user := models.CheckUserLoginInfo(session, wi.DB)

	data["user"] = map[string]interface{}{"loggedIn": user.IsLoggedIn()}

	wi.ParseCookieReason(session, &data)

	session.Save(r, w)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err = wi.Templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		log.Println(err)
	}
}
