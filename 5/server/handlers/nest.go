package handlers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/Joju-Matsumoto/go_web_app_practice/5/server/html"
)

func Nest(w http.ResponseWriter, r *http.Request) {
	data := "Hello World!"
	rand.Seed(time.Now().Unix())
	fn := []string{"layout.html"}
	if tmp := rand.Intn(3); tmp == 0 {
		fn = append(fn, "red.html")
	} else if tmp == 1 {
		fn = append(fn, "blue.html")
	}
	html.Generate(w, data, fn, html.BaseTemplate("layout"), html.BaseDir("templates/nest"))
}
