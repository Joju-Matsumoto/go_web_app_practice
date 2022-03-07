package main

import (
	"net/http"

	"github.com/Joju-Matsumoto/go_web_app_practice/5/server/handlers"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/process", handlers.Process)
	http.HandleFunc("/rand", handlers.ProcessRand)
	http.HandleFunc("/with", handlers.ProcessWith)
	http.HandleFunc("/include", handlers.ProcessInclude)
	http.HandleFunc("/map", handlers.ProcessMap)
	http.HandleFunc("/pipe", handlers.ProcessPipeline)

	http.HandleFunc("/custom_func", handlers.CustomFunc)

	http.HandleFunc("/context", handlers.Context)

	http.HandleFunc("/xss", handlers.Xss)
	http.HandleFunc("/xss_form", handlers.XssForm)

	http.HandleFunc("/nest", handlers.Nest)

	server.ListenAndServe()
}
