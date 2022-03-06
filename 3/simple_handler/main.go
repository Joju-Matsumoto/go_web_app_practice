package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, world!")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler Func called - " + name)
		h(w, r)
	}
}

func protect(h http.Handler, dh http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("pass") == "hello" {
			fmt.Println("Authenticate: OK")
			h.ServeHTTP(w, r)
		} else {
			fmt.Println("Authenticate: NG")
			dh.ServeHTTP(w, r)
		}
	})
}

type DefaultHandler struct{}

func (h DefaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is default page.")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	dh := DefaultHandler{}
	http.Handle("/hello", protect(log(hello), dh))
	server.ListenAndServe()
}
