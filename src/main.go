package main

import (
	"net/http"
	"log"

	"gopkg.in/mgo.v2"

	"github.com/nuzolx/library-api/src/router"
	"github.com/nuzolx/library-api/src/app"
)

var appC app.AppContext

func init() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	appC = app.AppContext{DB:session.DB("library")}
}

func main() {

	r := router.NewRouter(&appC)
	log.Fatal(http.ListenAndServe(":4242", r))
}
