package web

import (
	"github.com/RemmargorP/memjudge/utils"
	"net/http"
)

func (wi *WebInstance) ParseCookieReason(w http.ResponseWriter, r *http.Request, data *map[string]interface{}) bool {
	v := utils.GetCookieValue(r, "reason")

	if len(v) == 0 {
		return false
	}

	(*data)["hasReason"] = true
	(*data)["reason"] = v
	utils.EraseCookie(w, r, "reason")
	return true
}
