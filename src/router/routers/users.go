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
		RequeredAuthentication: true,
	},
	{
		URL:                    "/users/{userId}",
		Method:                 http.MethodGet,
		Function:               controllers.SearchUserById,
		RequeredAuthentication: true,
	},
	{
		URL:                    "/users/{userId}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdateUserbyId,
		RequeredAuthentication: true,
	},
	{
		URL:                    "/users/{userId}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteUserById,
		RequeredAuthentication: true,
	},
	{
		URL:                    "/users/{userId}/follow",
		Method:                 http.MethodPost,
		Function:               controllers.FollowUser,
		RequeredAuthentication: true,
	},
	{
		URL:                    "/users/{userId}/unfollow",
		Method:                 http.MethodPost,
		Function:               controllers.UnFollowUser,
		RequeredAuthentication: true,
	},
	{
		URL:                    "/users/{userId}/followers",
		Method:                 http.MethodGet,
		Function:               controllers.SearchFollowers,
		RequeredAuthentication: true,
	},
	{
		URL:                    "/users/{userId}/following",
		Method:                 http.MethodGet,
		Function:               controllers.SearchFollowing,
		RequeredAuthentication: true,
	},
}
