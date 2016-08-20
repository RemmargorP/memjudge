package web

import (
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	TITLE     = "MemJudge"
	PublicDir = "./public/"
)

type WebInstance struct {
	DB        *mgo.Database
	Router    *mux.Router
	Stop      <-chan bool
	Templates *template.Template
	Port      int
	Id        int
	ticker    *time.Ticker
}

func (wi *WebInstance) init() {
	wi.Router = mux.NewRouter()
	// Static files
	wi.Router.PathPrefix("/static/css/").Handler(http.StripPrefix("/static/css", http.FileServer(http.Dir(PublicDir+"css/"))))
	wi.Router.PathPrefix("/static/js/").Handler(http.StripPrefix("/static/js", http.FileServer(http.Dir(PublicDir+"js/"))))

	// Pages
	wi.Router.HandleFunc("/", wi.HomeHandler)
	wi.Router.HandleFunc("/signup", wi.SignUpHandler)
	wi.reloadTemplates()

	// API
	wi.Router.HandleFunc("/api/signup", wi.APISignUpHandler)

	wi.ticker = time.NewTicker(500 * time.Millisecond) //every 0.5s
	go wi.checkTicker()
}

func (wi *WebInstance) checkTicker() {
	for {
		select {
		case <-wi.ticker.C:
			wi.reloadTemplates()
		case <-wi.Stop:
			wi.stop()
			return
		}
	}
}

var funcMap = template.FuncMap{ // Custom template functions
	"not": func(a interface{}) interface{} {
		b := bool(a.(bool))
		return interface{}(!b)
	},
}

func (wi *WebInstance) reloadTemplates() {
	var err error
	wi.Templates, err = template.New("").Funcs(funcMap).ParseGlob(PublicDir + "html/*.html")
	if err != nil {
		log.Fatal(err)
	}
}

func (wi *WebInstance) stop() {
	wi.ticker.Stop()
}

func (wi *WebInstance) Serve() {
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(wi.Port), wi.Router))
}

func (wi *WebInstance) Start(id int, port int, stop <-chan bool, db *mgo.Database) {
	wi.Port = port
	wi.Id = id
	wi.DB = db
	wi.Stop = stop
	wi.init()
	wi.Serve()
}
