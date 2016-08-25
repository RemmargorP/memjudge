package models

type Contest struct {
	Id       Id   `bson:"_id"`
	Problems []Id `bson:"problems"`
	Users    []Id `bson:"users"`
	Rights   []Id `bson:"rights"` // UserId -> Right
}
