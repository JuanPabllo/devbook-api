package routers

import (
	"api/src/controllers"
	"net/http"
)

var routesPosts = []Router{
	{
		URL:                    "/posts",
		Method:                 http.MethodPost,
		Function:               controllers.CreatePosts,
		RequeredAuthentication: true,
	},
	{
		URL:                    "/posts",
		Method:                 http.MethodGet,
		Function:               controllers.GetPosts,
		RequeredAuthentication: true,
	},
	{
		URL:                    "/posts/{postID}",
		Method:                 http.MethodGet,
		Function:               controllers.GetPostByID,
		RequeredAuthentication: true,
	},
	{
		URL:                    "/posts/{postID}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdatePost,
		RequeredAuthentication: true,
	},
	{
		URL:                    "/posts/{postID}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeletePost,
		RequeredAuthentication: true,
	},
}