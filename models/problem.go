package models

type Problem struct {
	Id          Id     `bson:"_id"`
	Checker     Id     `bson:"checker"` // Submission
	Tests       []Test `bson:"tests"`
	Submissions []Id   `bson:"submissions"`
	Rights      []Id   `bson:"rights"` // UserId -> Right
}
