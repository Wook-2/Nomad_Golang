package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<h1><a href="/set">set a cookie</a></h1>`)
}

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: "value of cookie",
		Path:  "/",
	})
	fmt.Fprintln(w, `<h1><a href="/read">read</a></h1>`)

}

func read(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/set", http.StatusSeeOther)
		return
	}

	fmt.Fprintf(w, `<h1>Your Cookie:<br>%v</h1><h1><a href="/expire">expire</a></h1>`, cookie)

}

func expire(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")

	if err != nil {
		http.Redirect(w, r, "/set", http.StatusSeeOther)
		return
	}
	// 쿠키 삭제
	// cookie.MaxAge = -1

	// 만료기간? 5초
	cookie.MaxAge = 5
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
