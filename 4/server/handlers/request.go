package handlers

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"text/template"
)

func Headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	for k, v := range h {
		fmt.Fprintln(w, k, v)
	}
	fmt.Fprintln(w, h.Get("Accept"))
	fmt.Fprintln(w, h.Get("Accept-Encoding"))
}

func Body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func Form(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	r.ParseForm()

	// fmt.Fprintln(w, "encrypt")

	fmt.Fprintln(w, "<h3>http.Request.Form</h3>")
	fmt.Fprintln(w, "<p>フォーム，URLに同一のキーがある場合は両方がスライスに入れられる<br>フォームの値が優先して先に入る</p>")
	for k, v := range r.Form {
		fmt.Fprintf(w, "%v: %v<br>", k, v)
	}

	fmt.Fprintln(w, "<h3>http.Request.PostForm</h3>")
	fmt.Fprintln(w, "<p>フォームから送信された値のみ入る</p>")
	for k, v := range r.PostForm {
		fmt.Fprintf(w, "%v: %v<br>", k, v)
	}

	fmt.Fprintln(w, "<h3>Uploaded File</h3>")
	r.ParseMultipartForm(1024)

	var file multipart.File
	var err error

	// fileHeaders := r.MultipartForm.File["uploaded"]
	// if len(fileHeaders) > 0 {
	// 	file, err = fileHeaders[0].Open()
	// }

	// 簡単なやり方．fileが1つの場合のみ
	file, _, err = r.FormFile("uploaded")

	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

func FormPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/form.html")
	t.Execute(w, nil)
}
