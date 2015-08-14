package myjudge

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"hash/fnv"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"time"
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
	Header, _ := getPage("header.html")
	Header.Data["TITLE"] = TITLE
	Header.Data["HeaderTitle"] = "MyJudge - Online Judge."

	tmp, _ := template.New("").Parse(Header.Page)

	badReasonCookie, err := r.Cookie("badreason")
	var badReason string

	if err == nil {
		badReason = badReasonCookie.Value
		badReasonCookie.Path = "/"
		badReasonCookie.Value = ""
		badReasonCookie.Expires = time.Now().AddDate(0, 0, -1)
		http.SetCookie(w, badReasonCookie)
	}

	if badReason != "" {
		Header.Data["BadReason"] = badReason
	}

	cookie, err := r.Cookie(AuthInfoCookie)
	var User *User

	if err == nil {
		SessionID := cookie.Value
		wi.DBConnection.DB("myjudge").C("users").Find(bson.M{"lastsessionid": SessionID}).One(&User)

		if User != nil && User.LogoutDate.Before(time.Now()) {
			cookie.Value = ""
			cookie.Expires = User.LogoutDate
			http.SetCookie(w, cookie)
			User = nil
		}
	}
	if User != nil {
		Header.Data["isUserLoggined"] = true
		Header.Data["User"] = User
	}
	tmp.Execute(w, Header.Data)
}
func (wi *WebInstance) writeFooter(w http.ResponseWriter, r *http.Request) {
	WP, _ := getPage("footer.html")
	tmp, _ := template.New("").Parse(WP.Page)

	tmp.Execute(w, WP.Data)
}

func (wi *WebInstance) rootHandler(w http.ResponseWriter, r *http.Request) {
	wi.writeHeader(w, r)
	wi.writeFooter(w, r)
}
func (wi *WebInstance) loginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	val := r.PostForm
	login := val.Get("login")
	password := val.Get("password")

	Collection := wi.DBConnection.DB("myjudge").C("users")

	result := []User{}

	Collection.Find(bson.M{"login": login}).All(&result)

	if len(result) == 0 { //Login not found
		fmt.Fprint(w, `<head><meta http-equiv="refresh" content="0;/"></head>`)
		return
	}

	h := fnv.New64a() //password hasher
	h.Write([]byte(password))
	hash := int64(h.Sum64())

	user := result[0]
	if user.Password == hash {

		Session := GenerateSessionID()
		cookie := http.Cookie{}
		cookie.Name = AuthInfoCookie
		cookie.Path = "/"
		cookie.Expires = time.Now().AddDate(0, 0, 7)
		cookie.Value = Session
		cookie.HttpOnly = true

		http.SetCookie(w, &cookie)

		user.LastSessionDate = time.Now()
		user.LastSessionID = Session
		user.LogoutDate = time.Now().AddDate(0, 0, 7)

		Collection.Update(bson.M{"login": user.Login}, user)
	}

	fmt.Fprint(w, `<head><meta http-equiv="refresh" content="0;/"></head>`)
}

func (wi *WebInstance) logoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(AuthInfoCookie)
	if err != nil {
		return
	}
	cookie.Path = "/"
	cookie.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, cookie)
	fmt.Fprint(w, `<head><meta http-equiv="refresh" content="0;/"></head>`)
}
func (wi *WebInstance) registerHandler(w http.ResponseWriter, r *http.Request) {
	wi.writeHeader(w, r)
	r.ParseForm()

	WP, _ := getPage("forms/register.html")
	log.Println(WP.Page)
	templ, err := template.New("").Parse(WP.Page)
	if err != nil {
		log.Println(err)
	}
	templ.Execute(w, WP.Data)

	wi.writeFooter(w, r)
}

func (wi *WebInstance) registerCheckHandler(w http.ResponseWriter, r *http.Request) {
	bad := func(reason string) {
		cookie := http.Cookie{}
		cookie.Expires = time.Now().AddDate(0, 0, 1)
		cookie.Value = reason
		cookie.Name = "badreason"
		cookie.Path = "/"
		http.SetCookie(w, &cookie)

		http.Redirect(w, r, "/register", http.StatusFound)
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
