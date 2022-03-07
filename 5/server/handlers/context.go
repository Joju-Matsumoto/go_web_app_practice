package handlers

import (
	"net/http"

	"github.com/Joju-Matsumoto/go_web_app_practice/5/server/html"
)

func Context(w http.ResponseWriter, r *http.Request) {
	html.Generate(w, `I asked: <i>"What's up?"</i>`, []string{"context.html"})
}
