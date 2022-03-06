package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!", p.ByName("name"))
}

func defaultHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(w, "default page.")
}

func main() {
	mux := httprouter.New()

	mux.GET("/", defaultHandler)
	mux.GET("/hello/:name", hello)

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
