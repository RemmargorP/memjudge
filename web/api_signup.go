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

	var response struct {
		Status string `json:"status"`
		Reason string `json:"reason"`
		Redir  string `json:"redirect"`
	}

	var serialized []byte

	if err := decoder.Decode(&signupdata); err != nil {
		response.Reason = err.Error()
		response.Status = "FAIL"
	} else {
		u, err := api.SignUp(wi.DB, signupdata.Login, signupdata.Email, signupdata.Password)
		if u != nil {
			response.Reason = "Success"
			response.Redir = "/"
			response.Status = "OK"
		} else {
			response.Reason = err.Error()
			response.Status = "FAIL"
		}
	}

	serialized, _ = json.Marshal(response)
	w.Write(serialized)
}
