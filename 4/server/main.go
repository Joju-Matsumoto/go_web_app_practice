package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"text/template"
	"time"
)

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	for k, v := range h {
		fmt.Fprintln(w, k, v)
	}
	fmt.Fprintln(w, h.Get("Accept"))
	fmt.Fprintln(w, h.Get("Accept-Encoding"))
}

func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func form(w http.ResponseWriter, r *http.Request) {
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

func formPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/form.html")
	t.Execute(w, nil)
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>Go Web Programming - writeExample</title></head>
<body><h1>Hello World</h1></body>
</html>`
	// 	str := `{
	// 	"hello": "world"
	// }
	// `
	w.Write([]byte(str))
}

func writeHeader(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service is available.")
}

func redirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://google.com")
	w.WriteHeader(302)
}

type Post struct {
	User    string
	Threads []string
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "joju",
		Threads: []string{"first thread", "second thread", "third thread"},
	}
	jdata, _ := json.Marshal(post)
	w.Write(jdata)
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "row_cookie",
		Value:    "need to be baked",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "baked_cookie",
		Value:    "not bad",
		HttpOnly: true,
	}
	c3 := http.Cookie{
		Name:     "burned_cookie",
		Value:    "too bitter",
		HttpOnly: true,
	}
	w.Header().Set("Set-Cookie", c1.String())
	w.Header().Add("Set-Cookie", c2.String())
	http.SetCookie(w, &c3)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, "<h3>Your Cookies</h3>")
	// h := r.Header.Get("Cookie")
	// fmt.Fprintln(w, h)
	cs := r.Cookies()
	if len(cs) > 0 {
		for _, c := range cs {
			fmt.Fprintf(w, "<p>%v</p>", c)
		}
	} else {
		fmt.Fprintln(w, "<p>Oops! You lost cookies...</p>")
	}
}

func setMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Hello World!")
	c := http.Cookie{
		Name:  "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c)
}

func showMessage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "no messages.")
		}
	} else {
		rc := http.Cookie{
			Name:    "flash",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/headers", headers)
	http.HandleFunc("/body", body)
	http.HandleFunc("/form", form)
	http.HandleFunc("/formpage", formPage)

	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeader)
	http.HandleFunc("/redirect", redirect)
	http.HandleFunc("/json", jsonExample)

	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)

	http.HandleFunc("/set_message", setMessage)
	http.HandleFunc("/show_message", showMessage)

	server.ListenAndServe()
}
