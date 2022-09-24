package routers

import (
	"api/src/controllers"
	"net/http"
)

var routerLogin = Router{
	URL:                    "/login",
	Method:                 http.MethodPost,
	Function:               controllers.Login,
	RequeredAuthentication: false,
}
