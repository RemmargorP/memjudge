package models

type Problem struct {
	Id          Id
	Checker     Id // Submission
	Tests       []Test
	Submissions []Id
	Rights      []Id // UserId -> Right
}
