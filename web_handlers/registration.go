package memjudgeweb

import (
	"encoding/json"
	models "github.com/RemmargorP/memjudge/models"
	"log"
	"net/http"
)

func (wi *WebInstance) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["title"] = TITLE
	data["webinstanceid"] = wi.Id
	data["webinfo"] = map[string]interface{}{"isSignUp": true}

	session, err := wi.Store.Get(r, "session")
	if err != nil {
		session, err = wi.Store.New(r, "session")
		log.Println("Sessions fkt up: ", err)
	}

	user := models.CheckUserLoginInfo(session, wi.DB)

	if user.IsLoggedIn() {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	data["user"] = map[string]interface{}{"loggedIn": false}

	wi.ParseCookieReason(session, &data)
	session.Save(r, w)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err = wi.Templates.ExecuteTemplate(w, "registration.html", data)
	if err != nil {
		log.Println(err)
	}
}

func (wi *WebInstance) SignUpCheckHandler(w http.ResponseWriter, r *http.Request) {
	session, err := wi.Store.Get(r, "session")
	ok := true
	if err != nil {
		session, err = wi.Store.New(r, "session")
		log.Println("Sessions fkt up: ", err)
	}

	json_raw := r.PostFormValue("registration_data")
	var data []struct {
		login     string
		pwd       string
		pwdcheck  string
		email     string
		firstname string
		lastname  string
	}

	json.Unmarshal([]byte(json_raw), &data)

	session.Save(r, w)
	if ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
	}
}
