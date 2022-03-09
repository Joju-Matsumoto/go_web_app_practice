package main

import (
	"fmt"

	"github.com/Joju-Matsumoto/go_web_app_practice/6/postgresql/models"
)

func main() {
	post := models.Post{Content: "Hello, Posts!", Author: "joju"}

	fmt.Println(post)
	err := post.Create()
	if err != nil {
		panic(err)
	}
	fmt.Println(post)

	fmt.Println(models.Posts(10))

	redPost, err := models.GetPost(post.Id)
	if err != nil {
		panic(err)
	}
	fmt.Println(redPost)

	redPost.Content = "こんにちは"
	redPost.Author = "JOJU"
	redPost.Update()

	posts, err := models.Posts(10)
	if err != nil {
		panic(err)
	}
	fmt.Println(posts)

	// redPost.Delete()
	// fmt.Println(models.Posts(10))

	comment := models.Comment{Content: "hi", Author: "matsumoto", Post: &redPost}
	err = comment.Create()
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)

	comments, err := redPost.Comments()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(comments)

	for i := 0; i < 1000; i++ {
		redPost.Comments()
	}

	posts, _ = models.Posts(10)
	for _, post := range posts {
		err = post.Delete()
		if err != nil {
			panic(err)
		}
	}
}
