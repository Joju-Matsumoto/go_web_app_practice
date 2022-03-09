package main

import (
	"fmt"

	"github.com/Joju-Matsumoto/go_web_app_practice/6/sqlx/models"
)

func main() {
	post := models.Post{Content: "Hello", AuthorName: "joju"}
	err := post.Create()
	if err != nil {
		panic(err)
	}
	fmt.Println(post)
	redPost, err := models.GetPost(post.Id)
	if err != nil {
		panic(err)
	}
	fmt.Println(redPost)
}
