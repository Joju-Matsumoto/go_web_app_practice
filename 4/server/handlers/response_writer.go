package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>Go Web Programming - writeExample</title></head>
<body><h1>Hello World</h1></body>
</html>`
	// 	str := `{
	// 	"hello": "world"
	// }
	// `
	w.Write([]byte(str))
}

func WriteHeader(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service is available.")
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://google.com")
	w.WriteHeader(302)
}

type Post struct {
	User    string
	Threads []string
}

func JsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "joju",
		Threads: []string{"first thread", "second thread", "third thread"},
	}
	jdata, _ := json.Marshal(post)
	w.Write(jdata)
}
