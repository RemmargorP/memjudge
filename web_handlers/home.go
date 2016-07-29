package memjudgeweb

import (
	"html/template"
	"log"
	"net/http"
)

func (wi *WebInstance) HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(PublicDir + "html/index.html")
	if err != nil {
		log.Fatal(err)
	}

	data := make(map[string]interface{})

	var name string
	name = r.FormValue("name")
	log.Println(name)
	if len(name) == 0 {
		name = "$USER"
	}

	data["name"] = name
	data["title"] = TITLE

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}
