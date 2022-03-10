package main

import (
	"database/sql"
	"fmt"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model `json:"-"`
	Content    string
	Author     string `sql:"not null"`
	Comments   []Comment
}

func (p Post) String() string {
	return fmt.Sprintf("Post{ID: %d, Content: \"%s\", Author: \"%s\"}", p.ID, p.Content, p.Author)
}

type Comment struct {
	gorm.Model `json:"-"`
	Content    string
	Author     string `sql:"not null"`
	PostId     int
}

func (c Comment) String() string {
	return fmt.Sprintf("Comment{ID: %d, Content: \"%s\", Author: \"%s\"}", c.ID, c.Content, c.Author)
}

var Db *gorm.DB

type DbEnv struct {
	User string
	Name string
	Pass string
}

func init() {
	var dbEnv DbEnv
	envconfig.Process("db", &dbEnv)
	var err error
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", dbEnv.User, dbEnv.Name, dbEnv.Pass)
	// データベースに接続
	// Db, err = gorm.Open(postgres.Open(dsn))
	// if err != nil {
	// 	panic(err)
	// }
	// 既存のデータベースに接続
	sqldb, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	Db, err = gorm.Open(postgres.New(postgres.Config{
		Conn: sqldb,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 自動migration
	Db.AutoMigrate(&Post{}, &Comment{})
}

func main() {
	post := Post{Content: "Hello World!", Author: "joju"}

	result := Db.Create(&post)
	fmt.Println(post)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)

	comment := Comment{Content: "Good", Author: "matsumoto"}

	Db.Model(&post).Association("Comments").Append(&comment)

	for _, comment := range post.Comments {
		fmt.Println(comment)
	}

	var redPost Post
	Db.Where("id = $1", post.ID).First(&redPost)
	var comments []Comment
	Db.Model(&redPost).Association("Comments").Find(&comments)
	fmt.Println(comments)
}
