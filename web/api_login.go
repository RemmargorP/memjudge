package web

import (
	"encoding/json"
	"github.com/RemmargorP/memjudge/api"
	"net/http"
	"time"
)

func (wi *WebInstance) APILoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	decoder := json.NewDecoder(r.Body)
	var logindata struct {
		Login    string
		Password string
		Duration int
	}

	var response struct {
		Status string `json:"status"`
		Reason string `json:"reason"`
		Redir  string `json:"redirect"`
		SID    string `json:"sid"`
	}

	var serialized []byte

	if err := decoder.Decode(&logindata); err != nil {
		response.Reason = err.Error()
		response.Status = "FAIL"
	} else {
		u, err := api.Login(wi.DB, logindata.Login, logindata.Password, r.RemoteAddr, time.Second*time.Duration(logindata.Duration))
		if u != nil {
			response.Reason = "Success"
			response.Redir = "/"
			response.Status = "OK"
			response.SID = u.LastSID
		} else {
			response.Reason = err.Error()
			response.Status = "FAIL"
		}
	}

	serialized, _ = json.Marshal(response)
	w.Write(serialized)
}
