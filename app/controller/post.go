package controller

import (
	"github.com/naoina/kocha"
	"fmt"
	"kocha-sample/db"
	"kocha-sample/app/model"
	"kocha-sample/app/repository"
)

type Post struct {
	*kocha.DefaultController
}

func (this *Post) GET(c *kocha.Context) error {
	posts, err := repository.GetPost().FindAll()
  if err != nil {
   	return c.RenderError(500, nil, nil)
  }

	return c.Render(map[string]interface{}{
		"ControllerName": "Post",
		"posts": posts,
	})
}

func (this *Post) POST(c *kocha.Context) error {
	post := &model.Post{
		Title: c.Params.Get("title"),
		Content: c.Params.Get("content"),
	}

	_, err := db.Get("default").Insert(post)	
	if err != nil {
		c.App.Logger.Error("Post Insert error: ", err)
		return c.RenderError(500, nil, nil)
	}

	return c.Redirect("/post", false)
}


// PostDetail Controller
type PostDetail struct {
	*kocha.DefaultController
}

func (this *PostDetail) GET(c *kocha.Context) error {
	post, err := repository.GetPost().FindOne(c.Params.Get("id"))
  if err != nil {
  	c.App.Logger.Error("Query error: ", err)
   	return c.RenderError(500, nil, nil)
  }

  if (post.Id == 0) {
  	return c.RenderError(404, nil, nil)
  }

	return c.Render(map[string]interface{}{
		"ControllerName": "PostDetail",
		"post": post,
	})
}

// PostEdit Controller
type PostEdit struct {
	*kocha.DefaultController
}

func (this *PostEdit) GET(c *kocha.Context) error {
	post, err := repository.GetPost().FindOne(c.Params.Get("id"))
  if err != nil {
  	c.App.Logger.Error("Query error: ", err)
   	return c.RenderError(500, nil, nil)
  }

  if (post.Id == 0) {
  	return c.RenderError(404, nil, nil)
  }

	return c.Render(map[string]interface{}{
		"ControllerName": "PostEdit",
		"post": post,
	})
}

func (this *PostEdit) POST(c *kocha.Context) error {
	post, err := repository.GetPost().FindOne(c.Params.Get("id"))
  if err != nil {
  	c.App.Logger.Error("Query error: ", err)
   	return c.RenderError(500, nil, nil)
  }

  if (post.Id == 0) {
  	return c.RenderError(404, nil, nil)
  }

  post.Title = c.Params.Get("title")
  post.Content = c.Params.Get("content")
  _, updateError := db.Get("default").Update(&post)

  if updateError != nil {
  	c.App.Logger.Error("Post update error: ", updateError)
   	return c.RenderError(500, nil, nil)
  }

	return c.Redirect(fmt.Sprintf("/post/%s", c.Params.Get("id")), false)
}

