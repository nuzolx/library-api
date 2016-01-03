package router

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"

	"github.com/nuzolx/library-api/app"
)

func NewRouter(appContext *app.AppContext) *mux.Router {

	basicHandlers := alice.New(loggingHandler, recoverHandler)

	router := mux.NewRouter()

	for _, route := range GetRoutes(appContext) {
		router.Methods(route.Method).Path(route.Route).Handler(basicHandlers.ThenFunc(errorHandler(route.Handler)))
	}
	return router
}
