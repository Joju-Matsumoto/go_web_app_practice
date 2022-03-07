package main

import (
	"fmt"

	"github.com/Joju-Matsumoto/go_web_app_practice/6/memory/models"
)

func main() {
	models.PostById = make(map[int]*models.Post)
	models.PostsByAuthor = make(map[string][]*models.Post)

	p1 := models.Post{Id: 1, Content: "Hello", Author: "joju"}
	p2 := models.Post{Id: 2, Content: "World!", Author: "joju"}
	p3 := models.Post{Id: 3, Content: "Go", Author: "matsumoto"}
	p4 := models.Post{Id: 4, Content: "Programming", Author: "matsumoto"}
	p5 := models.Post{Id: 5, Content: "Language", Author: "matsumoto"}

	models.Store(&p1, &p2, &p3, &p4, &p5)

	fmt.Println(models.PostById[1])
	fmt.Println(models.PostById[2])

	for _, post := range models.PostsByAuthor["joju"] {
		fmt.Println(post)
	}
	for _, post := range models.PostsByAuthor["matsumoto"] {
		fmt.Println(post)
	}
}
