package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

type DbEnv struct {
	User string
	Name string
	Pass string
}

func init() {
	var dbEnv DbEnv
	envconfig.Process("db", &dbEnv)
	// DSN: Data Source Name
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", dbEnv.User, dbEnv.Name, dbEnv.Pass)
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Post{})
	fmt.Println("data.go init done.")
}

type Post struct {
	gorm.Model
	Content string
	Author  string
}

func retrieve(id int) (post Post, err error) {
	post = Post{}
	result := Db.First(&post, id)
	err = result.Error
	if err != nil {
		fmt.Println("Error retrieve:", err)
		return
	}
	return
}

func (p *Post) Create() (err error) {
	result := Db.Create(p)
	err = result.Error
	return
}

func (p *Post) Update() (err error) {
	result := Db.Updates(p)
	err = result.Error
	return
}

func (p *Post) Destroy() (err error) {
	result := Db.Delete(p)
	err = result.Error
	return
}
