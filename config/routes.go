package config

import (
	"github.com/naoina/kocha"
	"kocha-sample/app/controller"
)

type RouteTable kocha.RouteTable

var routes = RouteTable{
	{
		Name:       "root",
		Path:       "/",
		Controller: &controller.Root{},
	}, {
		Name:       "post/index",
		Path:       "/post",
		Controller: &controller.Post{},
	}, {
		Name:       "post/detail",
		Path:       "/post/:id",
		Controller: &controller.PostDetail{},
	}, {
		Name:       "post/edit",
		Path:       "/post/edit/:id",
		Controller: &controller.PostEdit{},
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
