package router

import (
	"github.com/nuzolx/library-api/app"
	"github.com/nuzolx/library-api/controllers"
)

type Route struct {
	Route   string
	Method  string
	Handler Handler
}

type Routes []Route

func GetRoutes(appContext *app.AppContext) Routes {

	server := &controllers.Server{Context: appContext}

	routes := Routes{
		Route{
			"/users",
			"POST",
			server.CreateUser,
		},
		Route{
			"/users/{userName}",
			"GET",
			server.GetUser,
		},
	}

	return routes
}
