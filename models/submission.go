package models

type Submission struct {
	Id       Id
	Problem  Id
	Owner    Id
	Language Id // Not UUID! Common names like "G++ 5.1 C++ 11", "Java 8.??", "Go 1.6", "Python 3.4.3", etc.
	Source   string
	Result   Id
	State    int8
}

const (
	Submission_Queued = iota
	Submission_Compiling
	Submission_Running
	Submission_Tested
)
