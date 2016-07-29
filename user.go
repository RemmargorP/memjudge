package memjudge

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"time"

	"gopkg.in/mgo.v2/bson"
)

//User definition
type User struct {
	Login           string        `bson:"login"`
	Password        int64         `bson:"passwordHash"`
	Name            string        `bson:"name"`
	ID              bson.ObjectId `bson:"_id"`
	Email           string        `bson:"email"`
	LastSessionDate time.Time     `bson:"lastsessiondate"`
	LastSessionID   string        `bson:"lastsessionid"`
	LogoutDate      time.Time     `bson:"logoutdate"`
}

func GenerateSessionID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
