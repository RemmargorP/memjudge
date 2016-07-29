package memjudge

import (
	"encoding/json"
	"fmt"
	"github.com/RemmargorP/memjudge/web_handlers"
	"github.com/gorilla/sessions"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

const MasterPort = ":45100"

type ServerConfig struct {
	NumJudges       int
	NumWebInstances int
}

func (s *ServerConfig) SpreadThreads(threads int) {
	if threads < 3 {
		s.NumWebInstances = 1
		s.NumJudges = 1
		return
	}
	s.NumWebInstances = threads / 3
	s.NumJudges = threads - s.NumWebInstances
}

func DefaultServerConfig() *ServerConfig {
	s := &ServerConfig{}
	s.SpreadThreads(runtime.NumCPU())
	return s
}

type Server struct {
	Config       *ServerConfig
	Judges       map[int]chan interface{} // chans used to kill specified judges
	WebInstances map[int]chan interface{} // or web instances
	lastThreadID int
	DB           *mgo.Database
}

func (s *Server) init() {
	db_auth_json, err := ioutil.ReadFile("DB_auth.json")
	if err != nil {
		log.Fatal(err)
	}

	var db_auth struct {
		Url             string
		DB              string
		User            string
		Pass            string
		CookieStoreSalt string
	}
	err = json.Unmarshal(db_auth_json, &db_auth)
	if err != nil {
		log.Fatal(err)
	}

	var session *mgo.Session
	session, err = mgo.Dial(db_auth.Url)
	if err != nil {
		log.Fatal(err)
	}
	session.SetSafe(&mgo.Safe{})

	s.DB = session.DB(db_auth.DB)

	err = s.DB.Login(db_auth.User, db_auth.Pass)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Successfully connected to the DB.\n")

	s.Config = DefaultServerConfig()

	log.Printf("Using Default Server Config:\n  judges: %d\n  web instances: %d\n  total threads: %d\n",
		s.Config.NumJudges, s.Config.NumWebInstances, s.Config.NumJudges+s.Config.NumWebInstances)

	s.Judges = make(map[int]chan interface{})
	s.WebInstances = make(map[int]chan interface{})

	for i := 0; i < s.Config.NumJudges; i++ {
		routine := &Judge{}
		stop := make(chan interface{}, 1)
		go routine.Start(stop, s.DB)
		s.Judges[s.lastThreadID] = stop
		s.lastThreadID += 1
	}

	cookieStore := sessions.NewCookieStore([]byte(db_auth.CookieStoreSalt))
	for i := 0; i < s.Config.NumWebInstances; i++ {
		routine := &memjudgeweb.WebInstance{}
		stop := make(chan interface{}, 1)
		go routine.Start(stop, s.DB, cookieStore)
		s.WebInstances[s.lastThreadID] = stop
		s.lastThreadID += 1
	}
}

func Stop(seconds int64) {
	time.Sleep(time.Duration(seconds) * time.Second)
	os.Exit(0)
}

func (s *Server) Serve() {
	w, _ := os.Create("log")
	log.SetOutput(w)
	s.init()
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	HttpServer := &http.Server{
		Addr:           MasterPort,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(HttpServer.ListenAndServe())
}

func handler(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	option := req.Form.Get("option")

	fmt.Fprintf(rw, "<html><title>Memjudge Control</title><body>")

	switch option {
	case "stop":
		fmt.Fprintf(rw, "<p><strong>Server gonna be stopped now.</strong></p>")
		log.Println("Shutdown initiated...")
		go Stop(2)
	default:
		fmt.Fprintf(rw, "<p>Unknown option: <strong>%s</strong></p>", option)
	}

	fmt.Fprintf(rw, "</body></html>")
}
