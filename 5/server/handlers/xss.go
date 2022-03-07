package handlers

import (
	"html/template"
	"net/http"

	"github.com/Joju-Matsumoto/go_web_app_practice/5/server/html"
)

func XssForm(w http.ResponseWriter, r *http.Request) {
	html.Generate(w, nil, []string{"xss.html"})
}

func Xss(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	html.Generate(w, template.HTML(r.FormValue("comment")), []string{"simple.html"})
}
