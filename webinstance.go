package myjudge

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	HTTPPORT  = ":8080"
	TITLE     = "Online Judge"
	PublicDir = "/var/www/myjudge/public/"
)

type WebInstance struct {
	submits chan SubmitionProtocol
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><title>%s</title><body>", TITLE)
	buffer, _ := ioutil.ReadFile(PublicDir + "index.html")
	fmt.Fprintf(w, string(buffer))
	fmt.Fprintf(w, "</body></html>")
}

func (wi *WebInstance) init() {
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(HTTPPORT, nil))
}

func (wi *WebInstance) Start(submits chan SubmitionProtocol) {
	wi.submits = submits
	wi.init()
}
