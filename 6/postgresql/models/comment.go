package models

import (
	"errors"
)

type Comment struct {
	Id      int
	Content string
	Author  string
	Post    *Post
}

func (comment *Comment) Create() (err error) {
	if comment.Post == nil {
		err = errors.New("comment has no post")
		return
	}
	err = Db.QueryRow("insert into comments (content, author, post_id) values ($1, $2, $3) returning id",
		comment.Content, comment.Author, comment.Post.Id).Scan(&comment.Id)
	return
}

func GetComments(post *Post) (comments []Comment, err error) {
	comments = []Comment{}
	rows, err := Db.Query("select id, content, author from comments where post_id = $1", post.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		comment := Comment{Post: post}
		err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
		if err != nil {
			return
		}
		comments = append(comments, comment)
	}
	return comments, err
}
