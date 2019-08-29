package models

import (
	"github.com/globalsign/mgo"
)

//DBSession ...
type DBSession struct {
	Session *mgo.Session
}

//Connection is variable for interaction
var Connection DBSession
