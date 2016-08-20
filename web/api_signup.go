package web

import (
	"encoding/json"
	"github.com/RemmargorP/memjudge/api"
	"net/http"
)

func (wi *WebInstance) APISignUpHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	decoder := json.NewDecoder(r.Body)
	var signupdata struct {
		Login    string
		Email    string
		Password string
	}

	// var response struct {
	// 	Reason string
	// 	Redir  string
	// 	SID    string
	// }

	if err := decoder.Decode(&signupdata); err != nil {
		s, _ := json.Marshal(map[string]interface{}{"reason": err.Error()})
		w.Write(s)
		return
	}

	u, err := api.SignUp(wi.DB, signupdata.Login, signupdata.Email, signupdata.Password)

	if u != nil {
		s, _ := json.Marshal(map[string]interface{}{"reason": "Success"})
		w.Write(s)
	} else {
		s, _ := json.Marshal(map[string]interface{}{"reason": err.Error()})
		w.Write(s)
	}

}
