package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func main() {
	db, err := InitDB()
	if err != nil {
		panic(err)
	}
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.Handle("/posts/", handleRequest(&Post{Db: db}))
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

func getID(ps httprouter.Params) (id int, err error) {
	id, err = strconv.Atoi(ps.ByName("id"))
	if err != nil {
		fmt.Println("Error strconv.Atoi:", err)
		return
	}
	return
}

func handleRequest(t Text) http.Handler {
	router := httprouter.New()

	router.GET("/posts/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		err := fetchRequestID(p, t)
		if err != nil {
			return
		}
		handleGet(w, r, t)
	})
	router.POST("/posts", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		err := handlePost(w, r, t)
		if err != nil {
			return
		}
	})
	router.PUT("/posts/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		err := fetchRequestID(p, t)
		if err != nil {
			return
		}
		// fmt.Println("PUT:")
		handlePut(w, r, t)
	})
	router.DELETE("/posts/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		err := fetchRequestID(p, t)
		if err != nil {
			return
		}
		handleDelete(w, r, t)
	})

	return router
}

func fetchRequestID(p httprouter.Params, t Text) (err error) {
	id, err := getID(p)
	if err != nil {
		return
	}
	err = t.Fetch(id)
	return
}

func handleGet(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	// http.NotFound(w, r)
	return writeJson(w, post)
}

func handlePost(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	// fmt.Println("HandlePost")
	body, err := readBody(r)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &post)
	if err != nil {
		return
	}
	err = post.Create()
	if err != nil {
		return
	}
	responseOK(w)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	// fmt.Println("HandlePut")
	body, err := readBody(r)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &post)
	if err != nil {
		return
	}
	err = post.Update()
	if err != nil {
		return
	}
	responseOK(w)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	// fmt.Println("HandleDelete")
	err = post.Destroy()
	if err != nil {
		return
	}
	responseOK(w)
	return
}
