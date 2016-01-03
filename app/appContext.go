package app

import "gopkg.in/mgo.v2"

type AppContext struct {
	DB *mgo.Database
}
