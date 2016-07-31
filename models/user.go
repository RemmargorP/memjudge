package memjudgemodels

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	id           bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	login        string        `bson:"login,omitempty" json:"login,omitempty"`
	passwordHash string        `bson:"pwd,omitempty" json:"-,omitpempty"`
	email        string        `bson:"email,omitempty" json:"email,omitempty"`
}

func (u *User) GetID() bson.ObjectId {
	return u.id
}
