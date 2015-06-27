package myjudge

import "gopkg.in/mgo.v2/bson"

//User definition
type User struct {
	Login    string        `bson:"login"`
	Password uint64        `bson:"passwordHash"`
	Name     string        `bson:"name"`
	ID       bson.ObjectId `bson:"_id"`
	Email    string        `bson:"email"`
}
