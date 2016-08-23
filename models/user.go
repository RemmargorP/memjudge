package models

type User struct {
	Id            Id
	Login         string
	Email         string
	Password      string // MD5 Hash
	FirstName     string
	LastName      string
	Submissions   []Submission
	Contests      []Id
	Administrator bool
}
