package models

type Submission struct {
	Id       Id     `bson:"_id"`
	Problem  Id     `bson:"problem"`
	Owner    Id     `bson:"owner"`
	Language Id     `bson:"language"` // Not UUID! Common names like "G++ 5.1 C++ 11", "Java 8.??", "Go 1.6", "Python 3.4.3", etc.
	Source   string `bson:"source"`
	Result   Id     `bson:"result"`
	State    int8   `bson:"state"`
}

const (
	Submission_Queued = iota
	Submission_Compiling
	Submission_Running
	Submission_Tested
)
