package app

import "github.com/nuzolx/library-api/Godeps/_workspace/src/gopkg.in/mgo.v2"

type AppContext struct {
	DB *mgo.Database
}
