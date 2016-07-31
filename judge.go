package memjudge

import (
	"gopkg.in/mgo.v2"
)

type Judge struct {
	DB *mgo.Database
}

func (j *Judge) Start(id int, stop <-chan interface{}, db *mgo.Database) {

}
