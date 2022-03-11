package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  Author `json:"author"`
}

func decodeUnmarshal() {
	jsonData, err := ioutil.ReadFile("post.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	var post Post
	err = json.Unmarshal(jsonData, &post)
	if err != nil {
		fmt.Println("Error Unmarshaling JSON data:", err)
		return
	}

	fmt.Println(post)
}

func decodeDecoder() {
	jsonFile, err := os.Open("post.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	for {
		var post Post
		err := decoder.Decode(&post)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}

		fmt.Println(post)
	}
}

func newPost() Post {
	return Post{
		Id:      1,
		Content: "Hello World!",
		Author: Author{
			Id:   1,
			Name: "joju",
		},
		Comments: []Comment{
			{
				Id:      3,
				Content: "Hi",
				Author: Author{
					Id:   2,
					Name: "Alice",
				},
			},
			{
				Id:      6,
				Content: "Bye",
				Author: Author{
					Id:   3,
					Name: "Bob",
				},
			},
		},
	}
}

func encodeMarshal() {
	post := newPost()

	jsonData, err := json.MarshalIndent(post, "", "    ")
	if err != nil {
		fmt.Println("Error Marshaling to JSON data:", err)
		return
	}

	err = ioutil.WriteFile("new_post.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}
}

func encodeEncoder() {
	post := newPost()

	jsonFile, err := os.Create("encode_post.json")
	if err != nil {
		fmt.Println("Error creating new file:", err)
		return
	}

	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding JSON to file:", err)
		return
	}
}

func main() {
	decodeUnmarshal()
	// decodeDecoder()
	// encodeMarshal()
	// encodeEncoder()
}
