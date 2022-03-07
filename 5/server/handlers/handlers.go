package handlers

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"

	"github.com/Joju-Matsumoto/go_web_app_practice/5/server/html"
)

func Process(w http.ResponseWriter, r *http.Request) {
	// ファイルから
	// t, _ := template.ParseFiles("templates/simple.html")
	//　エラーでパニック
	t := template.Must(template.ParseFiles("templates/simple.html"))

	// // 文字列から
	// tmp1 := `
	// <html>
	// <head>
	// <title>Hello world</title>
	// </head>
	// <body>
	// {{ . }}
	// </body>
	// </html>`
	// t := template.New("simple.html")
	// t, _ = t.Parse(tmp1)

	t.Execute(w, "Hello World!")
	// t.ExecuteTemplate(w, "simple.html", "Hello World!")
}

func ProcessRand(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/rand.html"))
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 4)
}

func ProcessWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/with.html"))
	t.Execute(w, "hello")
}

func ProcessInclude(w http.ResponseWriter, r *http.Request) {
	html.Generate(w, "Hello World!", []string{"t1.html", "t2.html"})
}

func ProcessMap(w http.ResponseWriter, r *http.Request) {
	m := map[string]string{}
	m["hello"] = "world"
	m["joju"] = "matsumoto"
	html.Generate(w, m, []string{"map.html"})
}

func ProcessPipeline(w http.ResponseWriter, r *http.Request) {
	html.Generate(w, nil, []string{"pipeline.html"})
}
