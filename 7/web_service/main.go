package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.POST("/posts", handlePost)
	router.GET("/posts/:id", handleGet)
	router.PUT("/posts/:id", handlePut)
	router.DELETE("/posts/:id", handleDelete)

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router,
	}
	server.ListenAndServe()
}

func responseOK(w http.ResponseWriter) {
	w.WriteHeader(200)
}

func readBody(r *http.Request) (body []byte, err error) {
	len := r.ContentLength
	body = make([]byte, len)
	r.Body.Read(body)
	return
}

func writeJson(w http.ResponseWriter, data interface{}) (err error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	return
}

func getPostByParams(ps httprouter.Params) (post Post, err error) {
	post = Post{}

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		fmt.Println("Error strconv.Atoi:", err)
		return
	}

	post, err = retrieve(id)
	if err != nil {
		return
	}

	return
}

func handleGet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// fmt.Println("HandleGet")
	post, err := getPostByParams(ps)
	if err != nil {
		return
	}
	writeJson(w, post)
}

func handlePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// fmt.Println("HandlePost")
	body, err := readBody(r)
	if err != nil {
		return
	}
	var post Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		return
	}
	err = post.Create()
	if err != nil {
		return
	}
	responseOK(w)
}

func handlePut(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// fmt.Println("HandlePut")
	post, err := getPostByParams(ps)
	if err != nil {
		return
	}
	body, err := readBody(r)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &post)
	if err != nil {
		return
	}
	// fmt.Println(post)
	err = post.Update()
	if err != nil {
		return
	}
	responseOK(w)
}

func handleDelete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// fmt.Println("HandleDelete")
	post, err := getPostByParams(ps)
	if err != nil {
		return
	}
	err = post.Destroy()
	if err != nil {
		return
	}
	responseOK(w)
}
