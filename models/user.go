package models

type User struct {
	Id            Id           `bson:"_id"`
	Login         string       `bson:"login"`
	Email         string       `bson:"email"`
	Password      string       `bson:"password"` // MD5 Hash
	FirstName     string       `bson:"first_name"`
	LastName      string       `bson:"last_name"`
	Submissions   []Submission `bson:"submissions"`
	Contests      []Id         `bson:"contests"`
	Administrator bool         `bson:"administrator"`
}
