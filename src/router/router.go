package router

import (
	"api/src/router/routers"

	"github.com/gorilla/mux"
)

func Generate() *mux.Router {
	router := mux.NewRouter()

	return routers.Config(router)
}
