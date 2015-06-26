package myjudge

import (
	"fmt"
	"log"
	"net/http"
)

const (
	HTTPPORT = ":8080"
	TITLE    = "Online Judge"
)

type WebInstance struct {
	submits chan SubmitionProtocol
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><title>%s</title><body>", TITLE)
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
