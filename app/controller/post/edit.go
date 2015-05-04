package post

import (
	"github.com/naoina/kocha"
	"fmt"
	"kocha-sample/db"
	"kocha-sample/app/model"
)

type Edit struct {
	*kocha.DefaultController
}

func (this *Edit) GET(c *kocha.Context) error {
	// FIXME: auto-generated by kocha
	fmt.Println("post edit GET called\n")

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
		"ControllerName": "Edit",
		"post": posts[0],
	})
}

func (this *Edit) POST(c *kocha.Context) error {
	// FIXME: auto-generated by kocha
	fmt.Println("post edit POST called\n")

	id := c.Params.Get("id")
	title := c.Params.Get("title")
	content := c.Params.Get("content")

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

  obj := posts[0]
  obj.Title = title
  obj.Content = content

  count, updateError := db.Get("default").Update(&obj)

  fmt.Printf("update err: %v\n", updateError)
  fmt.Printf("update count: %v\n", count)

	return c.Redirect(fmt.Sprintf("/post/%s", id), false)
}