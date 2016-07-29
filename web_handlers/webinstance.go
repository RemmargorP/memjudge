package memjudgeweb

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

const (
	HTTPPORT  = ":8080"
	TITLE     = "MemJudge"
	PublicDir = "./public/"
)

type WebInstance struct {
	DB     *mgo.Database
	Store  *sessions.CookieStore
	Stop   <-chan interface{}
	Router *mux.Router
}

func (wi *WebInstance) init() {
	wi.Router = mux.NewRouter()
	wi.Router.HandleFunc("/", wi.HomeHandler)
}

func (wi *WebInstance) Serve() {
	log.Fatal(http.ListenAndServe(HTTPPORT, wi.Router))
}

func (wi *WebInstance) Start(stop <-chan interface{}, db *mgo.Database, cookieStore *sessions.CookieStore) {
	wi.DB = db
	wi.Stop = stop
	wi.Store = cookieStore
	wi.init()
	wi.Serve()
}
