package myjudge

import (
	"fmt"
	"hash/fnv"
	"log"
	"net/http"
	"regexp"
	"time"

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
		`<table width="100%" cellpadding="0" cellspacing="0" border="0">`,
		`<tr>`, `<td align="center">`,
		`<p style="font-size:26px">`,
		`MyJudge - Online Judge.`,
		`</p>`,
		`</td>`, `</tr>`,
		`<tr height="20">`, `<td align="right">`,
	)
	cookie, err := r.Cookie(AuthInfoCookie)
	var User *User = nil

	if err != nil {
		expires := cookie.Expires
		SessionID := cookie.Value
		wi.DBConnection.DB("myjudge").C("users").Find(bson.M{"lastsessionid": SessionID}).One(&User)

		if User.LogoutDate.Before(expires) {
			cookie.MaxAge = -1
			cookie.Expires = User.LogoutDate
			http.SetCookie(w, cookie)
		}
	}
	if User == nil {
		fmt.Fprint(w,
			`<form action="/register/">`,
			`<input type="submit" value="Register">`,
			`</form>`,
			`<form action="/login/" method="post">`,
			`Login:`,
			`<input type="text" name="login" value size="8">  &nbsp;`,
			`Password:`,
			`<input type="password" name="password" value size="8">  &nbsp;`,
			`<input type="submit" value="Ok">`,
			`</form>`,
		)
	} else {
		fmt.Fprint(w, "You are logged in as <strong>", User.Name, "</strong>")
		fmt.Fprint(w,
			`<form action="/logout/">`,
			`<input type="submit" value="Logout">`,
			`</form>`,
		)
	}

	fmt.Fprint(w, `</td>`, `</tr>`)
}
func (wi *WebInstance) writeFooter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `</table>`)
	fmt.Fprint(w, `</body></html>`)
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
		fmt.Fprint(w, `<head><meta http-equiv="refresh" content="0;/"></head>`)
		return
	} else {
		h := fnv.New64a()
		h.Write([]byte(password))
		hash := int64(h.Sum64())
		user := result[0]
		if user.Password == hash {
			//TODO: Write Cookie and redirect
		} else {
			//Incorrect password - redirect to /
			fmt.Fprint(w, `<head><meta http-equiv="refresh" content="0;/"></head>`)
		}
	}
}
func (wi *WebInstance) logoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<head><meta http-equiv="refresh" content="0;/"></head>`)
	http.SetCookie(w, &http.Cookie{Name: AuthInfoCookie, MaxAge: -1})
}
func (wi *WebInstance) registerHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	reasonCookie, err := r.Cookie("badreason")
	var reason string

	if err == nil {
		reason = reasonCookie.Value
		reasonCookie.Path = "/"
		reasonCookie.Value = ""
		http.SetCookie(w, reasonCookie)
	}

	wi.writeHeader(w, r)
	if reason != "" {
		fmt.Fprintf(w, `<tr><td align="center"><font color="red"><strong>%s</strong></font></td></tr>`, reason)
	}

	fmt.Fprint(w, "",
		`<tr height="100"+"%">`, `<td align="center">`,
		`<form action="/register/check/" method="post">`,
		`<table cellpadding="0" cellspacing="10px" border="0">`,
		`<tr><td>Login*:</td>`,
		`<td><input type="text" name="login" size="20"></td><td>3+ symbols</td></tr><p>`,
		`<tr><td>Password*:</td>`,
		`<td><input type="password" name="password" size="20"></td><td>6+ symbols</td></tr><p>`,
		`<tr><td>Password check*:</td>`,
		`<td><input type="password" name="passwordcheck" size="20"></td><td>password again</td></tr><p>`,
		`<tr><td>Email*:</td>`,
		`<td><input type="text" name="email" size="20"></td><td>Email example@domain.com</td></tr><p>`,
		`<tr><td>Name:</td>`,
		`<td><input type="text" name="name" size="20"></td></tr><p>`,
		`</table>`,
		`<input type="submit" value="Ok"><p>`,
		`</form>`,
		`<font color="red">* - require</font>`,
	)

	fmt.Fprint(w, `</td></tr>`)

	wi.writeFooter(w, r)
}

func (wi *WebInstance) registerCheckHandler(w http.ResponseWriter, r *http.Request) {
	bad := func(reason string) {
		cookie := http.Cookie{}
		//cookie.MaxAge = 3600
		cookie.Expires = time.Now().AddDate(0, 0, 1)
		cookie.Value = reason
		cookie.Name = "badreason"
		cookie.Path = "/"
		http.SetCookie(w, &cookie)

		http.Redirect(w, r, "/register", http.StatusFound)
		//fmt.Fprint(w, `<head><meta http-equiv="refresh" content="0;/register"></head>`)
	}

	r.ParseForm()
	login := r.PostFormValue("login")
	password := r.PostFormValue("password")
	passwordCheck := r.PostFormValue("passwordcheck")
	email := r.PostFormValue("email")
	name := r.PostFormValue("name")

	users := wi.DBConnection.DB("myjudge").C("users")
	result := []User{}

	regexpEmail, _ := regexp.Compile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if len(login) <= 2 || len(password) < 6 || password != passwordCheck || !regexpEmail.MatchString(email) {
		bad("Some fields filled with bad data...")
		return
	}

	users.Find(bson.M{"login": login}).All(&result)
	if len(result) != 0 {
		bad("Login already used")
		return
	}

	users.Find(bson.M{"email": email}).All(&result)
	if len(result) != 0 {
		bad("Email already used")
		return
	}

	hasher := fnv.New64a()
	hasher.Write([]byte(password))
	PasswordHash := int64(hasher.Sum64())

	newbie := &User{Login: login, Password: PasswordHash, Name: name, ID: bson.NewObjectId(), Email: email, LogoutDate: time.Now()}

	err := users.Insert(newbie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, `<head><meta http-equiv="refresh" content="5;/"></head>`)
	fmt.Fprint(w, `<body><font color="green">Registration successed</font><p>You will be automaticaly redirected to the main page in 5 seconds</body>`)
}

func (wi *WebInstance) init() {
	http.HandleFunc("/", wi.rootHandler)
	http.HandleFunc("/login/", wi.loginHandler)
	http.HandleFunc("/logout/", wi.logoutHandler)
	http.HandleFunc("/register/", wi.registerHandler)
	http.HandleFunc("/register/check/", wi.registerCheckHandler)
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
