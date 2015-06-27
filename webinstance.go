package myjudge

import (
	"fmt"
	"hash/fnv"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	HTTPPORT        = ":8080"
	TITLE           = "Online Judge"
	PublicDir       = "/var/www/myjudge/public/"
	MONGODBUSER     = "myjudge"
	MONGODBPASSWORD = ""
)

type WebInstance struct {
	Submits      chan SubmitionProtocol
	DBConnection *mgo.Session
}

func (wi *WebInstance) writeHeader(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><title>%s</title><body>", TITLE)
	fmt.Fprint(w, "",
		"<table width=\"100%\" cellpadding=\"0\" cellspacing=\"0\" border=\"0\">",
		"<tr height=\"30\">", "<td align=\"center\">",
		"<p style=\"font-size:26px\">",
		"MyJudge - Online Judge.",
		"</p>",
		"</td>", "</tr>",
		"<tr height=\"20\">", "<td align=\"right\">",
		"</td>", "</tr>",
	)
	//TODO: check cookie user info
}
func (wi *WebInstance) writeFooter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "</table>")
	fmt.Fprint(w, "</body></html>")
}

func (wi *WebInstance) rootHandler(w http.ResponseWriter, r *http.Request) {
	wi.writeHeader(w, r)

	wi.writeFooter(w, r)
}
func (wi *WebInstance) loginHandler(w http.ResponseWriter, r *http.Request) {
	val := r.PostForm
	login := val.Get("login")
	password := val.Get("password")

	Collection := wi.DBConnection.DB("myjudge").C("users")

	result := []User{}

	Collection.Find(bson.M{"login": login}).All(&result)

	if len(result) == 0 {
		//Unknown login - redirect to /
		fmt.Fprint(w, "<meta http-equiv=\"refresh\" content=\"0; url=/\" />")
		return
	} else {
		h := fnv.New64a()
		h.Write([]byte(password))
		hash := h.Sum64()
		user := result[0]
		if user.Password == hash {
			//TODO: Write Cookie and redirect
		} else {
			//Incorrect password - redirect to /
			fmt.Fprint(w, "<meta http-equiv=\"refresh\" content=\"0; url=/\" />")
		}
	}

}

func (wi *WebInstance) init() {
	http.HandleFunc("/", wi.rootHandler)
	http.HandleFunc("/login/", wi.loginHandler)
	log.Fatal(http.ListenAndServe(HTTPPORT, nil))
}

func (wi *WebInstance) Start(submits chan SubmitionProtocol) {
	var err error
	wi.Submits = submits

	wi.DBConnection, err = mgo.Dial("mongodb://localhost/myjudge")
	if err != nil {
		log.Fatal(err)
	}
	defer wi.DBConnection.Close()
	wi.DBConnection.SetSafe(&mgo.Safe{})

	wi.init()
}
