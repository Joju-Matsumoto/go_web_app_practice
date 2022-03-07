package handlers

import (
	"html/template"
	"net/http"
	"time"
)

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func CustomFunc(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("custom_func.html").Funcs(funcMap)
	t, _ = t.ParseFiles("templates/custom_func.html")
	t.Execute(w, time.Now())
}
