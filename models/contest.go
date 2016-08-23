package models

type Contest struct {
	Id      Id
	Problem []Id
	Users   []Id
	Rights  []Id // UserId -> Right
}
