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

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err := wi.Templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		log.Println(err)
	}
}
