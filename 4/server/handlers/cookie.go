package handlers

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
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

func GetCookie(w http.ResponseWriter, r *http.Request) {
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

// Flash Message

func SetMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Hello World!")
	c := http.Cookie{
		Name:  "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c)
}

func ShowMessage(w http.ResponseWriter, r *http.Request) {
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
