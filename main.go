package main

import (
	"log"
	"net/http"

	"github.com/nuzolx/library-api/Godeps/_workspace/src/gopkg.in/mgo.v2"

	"github.com/nuzolx/library-api/app"
	"github.com/nuzolx/library-api/router"
)

var appC app.AppContext

func init() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	appC = app.AppContext{DB: session.DB("library")}
}

func main() {

	r := router.NewRouter(&appC)
	log.Fatal(http.ListenAndServe(":4242", r))
}
