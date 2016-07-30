package memjudgeweb

import (
	"html/template"
	"log"
	"net/http"
)

func (wi *WebInstance) WriteFooter(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(PublicDir + "html/footer.html")
	if err != nil {
		log.Fatal(err)
	}

	data := make(map[string]interface{})

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}
