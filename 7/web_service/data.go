package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbEnv struct {
	User string
	Name string
	Pass string
}

func InitDB() (db *gorm.DB, err error) {
	var dbEnv DbEnv
	envconfig.Process("db", &dbEnv)
	// DSN: Data Source Name
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", dbEnv.User, dbEnv.Name, dbEnv.Pass)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error Connect DB:", err)
	}
	db.AutoMigrate(&Post{})
	return
}

type Text interface {
	Fetch(id int) (err error)
	Create() (err error)
	Update() (err error)
	Destroy() (err error)
}

type Post struct {
	gorm.Model
	Db      *gorm.DB `json:"-" gorm:"-"`
	Content string
	Author  string
}

func (post *Post) Fetch(id int) (err error) {
	post.ID = 0
	result := post.Db.First(&post, id)
	err = result.Error
	return
}

func (post *Post) Create() (err error) {
	result := post.Db.Create(post)
	err = result.Error
	return
}

func (post *Post) Update() (err error) {
	result := post.Db.Updates(post)
	err = result.Error
	return
}

func (post *Post) Destroy() (err error) {
	result := post.Db.Delete(post)
	err = result.Error
	return
}
