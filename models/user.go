package models

import (
	"crypto/md5"
	"fmt"
	"github.com/RemmargorP/memjudge/utils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io"
	"net/http"
	"time"
)

type User struct {
	Id               bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Login            string        `bson:"login,omitempty" json:"login,omitempty"`
	PasswordHash     string        `bson:"pwd,omitempty" json:"-,omitpempty"`
	Email            string        `bson:"email,omitempty" json:"email,omitempty"`
	LastSID          string        `bson:"lastSID,omitempty" json:"-,omitempty"`
	LastSessionStart time.Time     `bson:"lastSessionStart,omitempty" json:"lastSessionStart,omitempty"`
	LastSessionEnd   time.Time     `bson:"lastSessionEnd,omitempty" json:"lastSessionEnd,omitempty"`
}

func CheckSignUpDataCorrectness(login, email, password string) bool {
	ok := true
	ok = ok && len(login) >= 3 && len(login) <= 64
	matched, _ := utils.ValidateEmail(email)
	ok = ok && matched
	ok = ok && len(password) >= 6 && len(password) <= 64
	return ok
}

func GetUserFromCookie(r *http.Request, db *mgo.Database) *User {
	sid := utils.GetCookieValue(r, "SID")

	if len(sid) == 0 {
		return nil
	}

	var user *User
	db.C("users").Find(bson.M{"lastSID": string(sid)}).One(&user)

	if user == nil {
		return nil
	}

	return user
}

func (u *User) IsLoggedIn() bool {
	if u == nil {
		return false
	}
	return time.Now().Before(u.LastSessionEnd)
}

func GenerateSID(login, ip string, t time.Time) string {
	h := md5.New()
	io.WriteString(h, login+"/"+ip)
	io.WriteString(h, t.String())
	return fmt.Sprintf("%x", h.Sum(nil))
}

func GatherUserInfo(u *User, data *map[string]interface{}) {
	(*data)["user"] = map[string]interface{}{
		"loggedIn": u.IsLoggedIn(),
		"login": func() string {
			if u == nil {
				return ""
			}
			return u.Login
		}(),
	}
}
