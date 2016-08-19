package memjudge

import (
	"encoding/json"
	"fmt"
	"github.com/RemmargorP/memjudge/judge"
	"github.com/RemmargorP/memjudge/web"
	"github.com/gorilla/sessions"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"time"
)

const MasterPort = ":45100"
const WebPort = ":8080"

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
	s.NumWebInstances = 2 // TODO DELETE (TESTING)
	s.NumJudges = threads - s.NumWebInstances
}

func DefaultServerConfig() *ServerConfig {
	s := &ServerConfig{}
	s.SpreadThreads(runtime.NumCPU())
	return s
}

type Server struct {
	Config       *ServerConfig
	Judges       map[int]chan bool // chans used to kill specified judges
	WebInstances map[int]chan bool // or web instances
	lastThreadId int
	DB           *mgo.Database
	Proxy        *httputil.ReverseProxy
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

	s.Judges = make(map[int]chan bool)
	s.WebInstances = make(map[int]chan bool)

	for i := 0; i < s.Config.NumJudges; i++ {
		routine := &judge.Judge{}
		stop := make(chan bool, 1)
		go routine.Start(s.lastThreadId, stop, s.DB)
		s.Judges[s.lastThreadId] = stop

		s.lastThreadId += 1
	}

	cookieStore := sessions.NewCookieStore([]byte(db_auth.CookieStoreSalt))
	var proxyTargets []*url.URL
	for i := 0; i < s.Config.NumWebInstances; i++ {
		routine := &web.WebInstance{}
		stop := make(chan bool, 1)
		go routine.Start(s.lastThreadId, 9000+s.lastThreadId, stop, s.DB, cookieStore)
		s.WebInstances[s.lastThreadId] = stop
		url, err := url.Parse("http://127.0.0.1:" + strconv.Itoa(9000+s.lastThreadId))
		if err != nil {
			log.Fatal(err)
		}
		proxyTargets = append(proxyTargets, url)

		s.lastThreadId += 1
	}

	s.Proxy = &httputil.ReverseProxy{
		Director: func(r *http.Request) {
			target := proxyTargets[rand.Int()%len(proxyTargets)]
			r.URL.Scheme = target.Scheme
			r.URL.Host = target.Host
			//r.URL.Path = target.Path
		},
	}
}

func (s *Server) Stop(seconds int64) {
	for _, j := range s.Judges {
		j <- true
	}
	for _, wi := range s.WebInstances {
		wi <- true
	}
	time.Sleep(time.Duration(seconds) * time.Second)
	os.Exit(0)
}

func (s *Server) Serve() {
	w, _ := os.Create("log")
	log.SetOutput(w)
	s.init()

	go func() { // PROXY (RANDOMIZED LOAD BALANCER)
		log.Fatal(http.ListenAndServe(WebPort, s.Proxy))
	}()
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handler)
	HttpServer := &http.Server{
		Addr:           MasterPort,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(HttpServer.ListenAndServe())
}

func (s *Server) handler(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	option := req.Form.Get("option")

	fmt.Fprintf(rw, "<html><title>Memjudge Control</title><body>")

	switch option {
	case "stop":
		fmt.Fprintf(rw, "<p><strong>Server gonna be stopped now.</strong></p>")
		log.Println("Shutdown initiated...")
		go s.Stop(1)
	default:
		fmt.Fprintf(rw, "<p>Unknown option: <strong>%s</strong></p>", option)
	}

	fmt.Fprintf(rw, "</body></html>")
}
