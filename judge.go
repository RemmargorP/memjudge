package memjudge

import (
	"gopkg.in/mgo.v2"
)

type Judge struct {
	DB *mgo.Database
}

func (j *Judge) Start(stop <-chan interface{}, db *mgo.Database) {

}
