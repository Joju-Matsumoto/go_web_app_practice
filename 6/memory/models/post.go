package models

import "fmt"

type Post struct {
	Id      int
	Content string
	Author  string
}

var PostById map[int]*Post
var PostsByAuthor map[string][]*Post

func Store(posts ...*Post) {
	fmt.Println(posts)
	for _, post := range posts {
		fmt.Println("store:", post, &post)
		PostById[post.Id] = post
		PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], post)
	}
	fmt.Println(PostById)
	fmt.Println(PostsByAuthor)
}
