package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"

	"github.com/Joju-Matsumoto/go_web_app_practice/6/memory/models"
)

func store(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

func load(data interface{}, filename string) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}

}

func main() {
	post := models.Post{Id: 1, Content: "Hello", Author: "joju"}
	store(post, "post1")
	var redPost models.Post
	load(&redPost, "post1")
	fmt.Println(redPost)
}
