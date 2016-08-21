package api

import (
	"errors"
	"github.com/RemmargorP/memjudge/models"
	"github.com/RemmargorP/memjudge/utils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func Login(db *mgo.Database, login, password, ip string, sessionDuration time.Duration) (*models.User, error) {
	var user *models.User

	db.C("users").Find(bson.M{"login": login}).One(&user)
	if user == nil {
		return nil, errors.New("User with given login doesn't exist")
	}

	if user.PasswordHash != utils.Encrypt(password) {
		return nil, errors.New("Invalid password")
	}

	user.LastSID = models.GenerateSID(login, ip, time.Now().UTC())
	user.LastSessionStart = time.Now().UTC()
	user.LastSessionEnd = time.Now().UTC().Add(sessionDuration)

	db.C("users").UpdateId(user.Id, bson.M{
		"$set": bson.M{
			"lastSID":          user.LastSID,
			"lastSessionStart": user.LastSessionStart,
			"lastSessionEnd":   user.LastSessionEnd,
		},
	})

	return user, nil
}
