package memjudgeweb

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
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
	Store     *sessions.CookieStore
	Stop      <-chan bool
	Router    *mux.Router
	Templates *template.Template
	Id        int
	Port      int
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
	wi.Router.HandleFunc("/signup/handle", wi.SignUpCheckHandler)
	wi.reloadTemplates()

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

func (wi *WebInstance) Start(id int, port int, stop <-chan bool, db *mgo.Database, cookieStore *sessions.CookieStore) {
	wi.Port = port
	wi.Id = id
	wi.DB = db
	wi.Stop = stop
	wi.Store = cookieStore
	wi.init()
	wi.Serve()
}
