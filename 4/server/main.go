package main

import (
	"net/http"

	"github.com/Joju-Matsumoto/go_web_app_practice/4/server/handlers"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/headers", handlers.Headers)
	http.HandleFunc("/body", handlers.Body)
	http.HandleFunc("/form", handlers.Form)
	http.HandleFunc("/formpage", handlers.FormPage)

	http.HandleFunc("/write", handlers.WriteExample)
	http.HandleFunc("/writeheader", handlers.WriteHeader)
	http.HandleFunc("/redirect", handlers.Redirect)
	http.HandleFunc("/json", handlers.JsonExample)

	http.HandleFunc("/set_cookie", handlers.SetCookie)
	http.HandleFunc("/get_cookie", handlers.GetCookie)

	http.HandleFunc("/set_message", handlers.SetMessage)
	http.HandleFunc("/show_message", handlers.ShowMessage)

	server.ListenAndServe()
}
