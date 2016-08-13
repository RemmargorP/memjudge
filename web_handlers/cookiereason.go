package memjudgeweb

import (
	"github.com/gorilla/sessions"
)

func (wi *WebInstance) ParseCookieReason(session *sessions.Session, data *map[string]interface{}) bool {
	v := session.Values["reason"]
	switch v.(type) {
	case string:
		s := string(v.(string))
		(*data)["hasReason"] = true
		(*data)["reason"] = s
		session.Values["reason"] = 0
		return true
	default:
		return false
	}
}
