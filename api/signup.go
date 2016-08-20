package api

import (
	"errors"
	"github.com/RemmargorP/memjudge/models"
	"github.com/RemmargorP/memjudge/utils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func SignUp(db *mgo.Database, login, email, password string) (*models.User, error) {
	if !models.CheckUserDataCorrectness(login, email, password) {
		return nil, errors.New("Invalid login, email or password")
	}

	var user *models.User

	db.C("users").Find(bson.M{"login": login}).One(&user)
	if user != nil {
		return nil, errors.New("Login already taken")
	}

	db.C("users").Find(bson.M{"email": email}).One(&user)
	if user != nil {
		return nil, errors.New("Email already taken")
	}

	user = &models.User{bson.NewObjectId(), login, utils.Encrypt(password), email, "", time.Unix(0, 0), 1 * time.Millisecond}

	db.C("users").Insert(user)

	return user, nil
}
