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
	// Static files
	wi.Router.PathPrefix("/static/css/").Handler(http.StripPrefix("/static/css", http.FileServer(http.Dir(PublicDir+"css/"))))
	wi.Router.PathPrefix("/static/js/").Handler(http.StripPrefix("/static/js", http.FileServer(http.Dir(PublicDir+"js/"))))

	// Pages
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
