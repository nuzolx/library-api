package main

import (
	"log"
	"net/http"
	"os"

	"github.com/nuzolx/library-api/Godeps/_workspace/src/gopkg.in/mgo.v2"

	"github.com/nuzolx/library-api/app"
	"github.com/nuzolx/library-api/router"
)

var appC app.AppContext

func init() {

	url := os.Getenv("MONGOLAB_URI")
	if len(url) <= 0 {
		url = "localhost"
	}

	session, err := mgo.Dial(url)

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	appC = app.AppContext{DB: session.DB("library")}
}

func main() {

	r := router.NewRouter(&appC)

	port := os.Getenv("PORT")
	if len(port) <= 0 {
		port = "4242"
	}
	log.Fatal(http.ListenAndServe(":"+port, r))
}
