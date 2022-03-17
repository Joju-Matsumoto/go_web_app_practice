package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type FakePost struct {
	ID      uint
	Content string
	Author  string
}

func (post *FakePost) Fetch(id int) (err error) {
	post.ID = uint(id)
	post.Content = "Hello World!"
	post.Author = "joju"
	return
}

func (post *FakePost) Create() (err error) {
	return
}

func (post *FakePost) Update() (err error) {
	return
}

func (post *FakePost) Destroy() (err error) {
	return
}

func TestHandleGet(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/posts/", HandleRequest(&FakePost{}))

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/posts/1", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.ID != 1 {
		t.Error("Cannnot retrieve JSON post")
	}
}
