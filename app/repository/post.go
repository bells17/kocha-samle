package repository

import (
	"kocha-sample/db"
	"kocha-sample/app/model"
)

type Post struct {

}

func (this *Post) FindAll() ([]model.Post, error) {
	var posts []model.Post
	err := db.Get("default").Select(&posts)
	return posts, err
}

func (this *Post) FindOne(id string) (model.Post, error) {
	var post model.Post
	var posts []model.Post
	err := db.Get("default").Select(&posts, db.Get("default").Where("id", "=", id))

	if len(posts) == 1 {
		post = posts[0]
	}

  return post, err
}


