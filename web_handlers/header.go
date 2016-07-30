package memjudgeweb

import (
	"html/template"
	"log"
	"net/http"
)

func (wi *WebInstance) WriteHeader(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(PublicDir + "html/header.html")
	if err != nil {
		log.Fatal(err)
	}

	data := make(map[string]interface{})

	data["title"] = TITLE

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}
