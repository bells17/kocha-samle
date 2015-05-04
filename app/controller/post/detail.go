package post

import (
	"github.com/naoina/kocha"
	"fmt"
	"kocha-sample/db"
	"kocha-sample/app/model"
)

type Detail struct {
	*kocha.DefaultController
}

func (this *Detail) GET(c *kocha.Context) error {
	// FIXME: auto-generated by kocha
	fmt.Println("post detail GET called\n")

	id := c.Params.Get("id")

	var posts []model.Post
	err := db.Get("default").Select(&posts, db.Get("default").Where("id", "=", id))
	fmt.Printf("select err: %v\n", err)
  if err != nil {
   	return c.RenderError(500, nil, nil)
  }

  fmt.Printf("posts all %#v\n", posts)

  if len(posts) != 1 {
		return c.RenderError(404, nil, nil)  	
  }

	return c.Render(map[string]interface{}{
		"ControllerName": "Detail",
		"post": posts[0],
	})
}