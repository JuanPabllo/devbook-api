package routers

import (
	"api/src/controllers"
	"net/http"
)

var routersUsers = []Router{
	{
		URL:                    "/users",
		Method:                 http.MethodPost,
		Function:               controllers.CreateUser,
		RequeredAuthentication: false,
	},
	{
		URL:                    "/users",
		Method:                 http.MethodGet,
		Function:               controllers.SearchUsers,
		RequeredAuthentication: false,
	},
	{
		URL:                    "/users/{userId",
		Method:                 http.MethodGet,
		Function:               controllers.SearchUserById,
		RequeredAuthentication: false,
	},
	{
		URL:                    "/users/{userId}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdateUserbyId,
		RequeredAuthentication: false,
	},
	{
		URL:                    "/users/{userId}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteUserById,
		RequeredAuthentication: false,
	},
}
