package memjudgemodels

import (
	"github.com/gorilla/sessions"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const SIDLength = 32

type User struct {
	Id              bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Login           string        `bson:"login,omitempty" json:"login,omitempty"`
	PasswordHash    string        `bson:"pwd,omitempty" json:"-,omitpempty"`
	Email           string        `bson:"email,omitempty" json:"email,omitempty"`
	LastSID         string        `bson:"lastSID,omitempty" json:"-,omitempty"`
	LastLoginDate   time.Time     `bson:"lastLoginDate,omitempty" json:"lastLoginDate,omitempty"`
	LastLoginMaxAge time.Duration `bson:"lastLoginMaxAge,omitempty" json:"lastLoginMaxAge,omitempty"`
}

func CheckUserLoginInfo(s *sessions.Session, db *mgo.Database) *User {
	sid_raw := s.Values["SID"]
	var sid string
	switch sid_raw.(type) {
	case string:
		sid = string(sid_raw.(string))
	default:
		return nil
	}
	if len(string(sid)) != SIDLength {
		return nil
	}
	var user *User
	db.C("users").Find(bson.M{"lastSID": string(sid)}).One(user)

	if user == nil {
		return nil
	}

	return user
}

func (u *User) IsLoggedIn() bool {
	if u == nil {
		return false
	}
	return time.Now().Before(u.LastLoginDate.Add(u.LastLoginMaxAge))
}
