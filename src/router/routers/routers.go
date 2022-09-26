package routers

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	URL                    string
	Method                 string
	Function               func(res http.ResponseWriter, req *http.Request)
	RequeredAuthentication bool
}

func Config(router *mux.Router) *mux.Router {
	routers := routersUsers
	routers = append(routers, routerLogin)
	routers = append(routers, routesPosts...)

	for _, route := range routers {
		if route.RequeredAuthentication {
			router.HandleFunc(
				route.URL,
				middlewares.Logger(middlewares.Authenticate(route.Function)),
			).Methods(route.Method)
		} else {
			router.HandleFunc(route.URL, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	return router
}
