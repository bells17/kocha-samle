package config

import (
	"github.com/naoina/kocha"
	"kocha-sample/app/controller"
	"kocha-sample/app/controller/post"
)

type RouteTable kocha.RouteTable

var routes = RouteTable{
	{
		Name:       "root",
		Path:       "/",
		Controller: &controller.Root{},
	}, {
		Name:       "post",
		Path:       "/post",
		Controller: &controller.Post{},
	}, {
		Name:       "post_detail",
		Path:       "/post/:id",
		Controller: &post.Detail{},
	}, {
		Name:       "post_edit",
		Path:       "/post/edit/:id",
		Controller: &post.Edit{},
	},
}

func init() {
	AppConfig.RouteTable = kocha.RouteTable(append(routes, RouteTable{
		{
			Name:       "static",
			Path:       "/*path",
			Controller: &kocha.StaticServe{},
		},
	}...))
}
