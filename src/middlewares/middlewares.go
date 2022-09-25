package middlewares

import (
	"api/src/auth"
	"api/src/responses"
	"log"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		log.Printf("\n %s %s %s", req.Method, req.RequestURI, req.Host)
		next(res, req)
	}
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if erro := auth.ValidateToken(req); erro != nil {
			responses.Error(res, http.StatusUnauthorized, erro)
			return
		}

		next(res, req)
	}
}
