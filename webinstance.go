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
	AuthInfoCookie  = "authinfo"
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
	)
	cookie, _ := r.Cookie(AuthInfoCookie)
	log.Println(cookie)
	var User *User
	User = nil

	if cookie != nil {
		expires := cookie.Expires
		SessionID := cookie.Value
		wi.DBConnection.DB("myjudge").C("users").Find(bson.M{"lastsessionid": SessionID}).One(&User)

		log.Println(User)
		if User.LogoutDate.Before(expires) {
			cookie.MaxAge = -1
			cookie.Expires = User.LogoutDate
			http.SetCookie(w, cookie)
		}
	}
	log.Println(User)
	if User == nil {
		fmt.Fprint(w,
			"<form action=\"/register/\">",
			"<input type=\"submit\" value=\"Register\" style=\"height:18\">",
			"</form>",
			"<form action=\"/login/\" method=\"post\">",
			"Login:",
			"<input type=\"text\" name=\"login\" value size=\"8\" style=\"height:18\">  &nbsp;",
			"Password:",
			"<input type=\"password\" name=\"password\" value size=\"8\" style=\"height:18\">  &nbsp;",
			"<input type=\"submit\" value=\"Ok\" style=\"height:18\">",
			"</form>",
		)
	} else {
		fmt.Fprint(w, "You are logged in as <strong>", User.Name, "</strong>")
		fmt.Fprint(w,
			"<form action=\"/logout/\">",
			"<input type=\"submit\" value=\"Logout\" style=\"height:18\">",
			"</form>",
		)
	}

	fmt.Fprint(w, "</td>", "</tr>")
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
		fmt.Fprint(w, "<head><meta http-equiv=\"refresh\" content=\"0;/\"></head>")
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
			fmt.Fprint(w, "<head><meta http-equiv=\"refresh\" content=\"0;/\"></head>")
		}
	}
}
func (wi *WebInstance) logoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<head><meta http-equiv=\"refresh\" content=\"0;/\"></head>")
	http.SetCookie(w, &http.Cookie{Name: AuthInfoCookie, MaxAge: -1})
}
func (wi *WebInstance) registerHandler(w http.ResponseWriter, r *http.Request) {
	wi.writeHeader(w, r)

	fmt.Fprint(w,
		"<tr height=\"100"+"%\">", "<td align=\"center\">",
		"<form action=\"/register/check/\" method=\"post\">",
		"<table cellpadding=\"0\" cellspacing=\"10px\" border=\"0\">",
		"<tr><td>Login:</td>",
		"<td><input type=\"text\" name=\"login\" value size=\"20\"></td></tr><p>",
		"<tr><td>Password:</td>",
		"<td><input type=\"password\" name=\"password\" value size=\"20\"></td></tr><p>",
		"<tr><td>Password check:</td>",
		"<td><input type=\"password\" name=\"passwordcheck\" value size=\"20\"></td></tr><p>",
		"<tr><td>Email:</td>",
		"<td><input type=\"text\" name=\"email\" value size=\"20\"></td></tr><p>",
		"<tr><td>Name:</td>",
		"<td><input type=\"text\" name=\"name\" value size=\"20\"></td></tr><p>",
		"</table>",
		"<input type=\"submit\" value=\"Ok\" style=\"height:18\"><p>",
		"</form>",
	)

	fmt.Fprint(w, "</td></tr>")

	wi.writeFooter(w, r)
}

func (wi *WebInstance) init() {
	http.HandleFunc("/", wi.rootHandler)
	http.HandleFunc("/login/", wi.loginHandler)
	http.HandleFunc("/logout/", wi.logoutHandler)
	http.HandleFunc("/register/", wi.registerHandler)
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
