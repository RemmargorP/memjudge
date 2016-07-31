package memjudgeweb

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"gopkg.in/mgo.v2"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

const (
	TITLE     = "MemJudge"
	PublicDir = "./public/"
)

type WebInstance struct {
	DB        *mgo.Database
	Store     *sessions.CookieStore
	Stop      <-chan interface{}
	Router    *mux.Router
	Templates *template.Template
	Id        int
	Port      int
}

func (wi *WebInstance) init() {
	wi.Router = mux.NewRouter()
	// Static files
	wi.Router.PathPrefix("/static/css/").Handler(http.StripPrefix("/static/css", http.FileServer(http.Dir(PublicDir+"css/"))))
	wi.Router.PathPrefix("/static/js/").Handler(http.StripPrefix("/static/js", http.FileServer(http.Dir(PublicDir+"js/"))))

	// Pages
	wi.Router.HandleFunc("/", wi.HomeHandler)

	var err error
	wi.Templates, err = template.ParseGlob(PublicDir + "html/*.html")
	if err != nil {
		log.Fatal(err)
	}
}

func (wi *WebInstance) Serve() {
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(wi.Port), wi.Router))
}

func (wi *WebInstance) Start(id int, port int, stop <-chan interface{}, db *mgo.Database, cookieStore *sessions.CookieStore) {
	wi.Port = port
	wi.Id = id
	wi.DB = db
	wi.Stop = stop
	wi.Store = cookieStore
	wi.init()
	wi.Serve()
}
