package myjudge

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type SubmitionProtocol struct {
	Id     bson.ObjectId
	Data   string
	Result int
}

const MasterPort = ":45100"
const packetBufferSize = 100

type ServerConfig struct {
	numGoRoutines   int
	numJudges       int
	numWebInstances int
}

func (s *ServerConfig) run() {
	if s.numJudges == 0 {
		s.numJudges = 1
	}
	if s.numWebInstances == 0 {
		s.numWebInstances = 1
	}

	if s.numJudges+1+s.numWebInstances < runtime.NumCPU() {
		s.numJudges = runtime.NumCPU() - 1 - s.numWebInstances
	}

	s.numGoRoutines = 1 + s.numJudges + s.numWebInstances
}

func DefaultServerConfig() *ServerConfig {
	s := &ServerConfig{}
	s.run()
	return s
}

type Server struct {
	Config  *ServerConfig
	submits chan SubmitionProtocol
}

func (s *Server) init() {
	runtime.GOMAXPROCS(s.Config.numGoRoutines)
	fmt.Print(s.Config.numWebInstances)
	s.submits = make(chan SubmitionProtocol, packetBufferSize)
	for i := 0; i < s.Config.numJudges; i++ {
		routine := &Judge{}
		go routine.Start(s.submits)
	}
	for i := 0; i < s.Config.numWebInstances; i++ {
		routine := &WebInstance{}
		go routine.Start(s.submits)
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

	fmt.Fprintf(rw, "<html><title>Myjudge Control</title><body>")

	switch option {
	case "stop":
		fmt.Fprintf(rw, "<p><strong>Server gonna be stopped now.</strong></p>")
		go Stop(2)
	default:
		fmt.Fprintf(rw, "<p>Unknown option: <strong>%s</strong></p>", option)
	}

	fmt.Fprintf(rw, "</body></html>")
}
