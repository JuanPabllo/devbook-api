package routers

import (
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

	for _, route := range routers {
		router.HandleFunc(route.URL, route.Function).Methods(route.Method)
	}

	return router
}
