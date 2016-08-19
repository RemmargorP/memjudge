package web

import (
	"encoding/json"
	"github.com/RemmargorP/memjudge/api"
	"github.com/RemmargorP/memjudge/models"
	"log"
	"net/http"
)

func (wi *WebInstance) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["title"] = TITLE
	data["webinstanceid"] = wi.Id
	data["webinfo"] = map[string]interface{}{"isSignUp": true}
	data["additionalJS"] = []string{"signup.js"}

	session, err := wi.Store.Get(r, "session")
	if err != nil {
		session, err = wi.Store.New(r, "session")
		log.Println("Sessions fkt up: ", err)
	}

	user := models.GetUserFromSession(session, wi.DB)

	if user.IsLoggedIn() {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	data["user"] = map[string]interface{}{"loggedIn": false}

	wi.ParseCookieReason(session, &data)
	session.Save(r, w)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err = wi.Templates.ExecuteTemplate(w, "signup.html", data)
	if err != nil {
		log.Println(err)
	}
}

func (wi *WebInstance) APISignUpHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	decoder := json.NewDecoder(r.Body)
	var signupdata struct {
		Login    string
		Email    string
		Password string
	}

	if err := decoder.Decode(&signupdata); err != nil {
		s, _ := json.Marshal(map[string]interface{}{"reason": err.Error()})
		w.Write(s)
		return
	}

	u, err := api.SignUp(wi.DB, signupdata.Login, signupdata.Email, signupdata.Password)

	if u != nil {
		log.Println("tuta4ki")
		s, _ := json.Marshal(u)
		w.Write(s)
	} else {
		log.Println("tama4ki")
		s, _ := json.Marshal(map[string]interface{}{"reason": err.Error()})
		w.Write(s)
	}

}
