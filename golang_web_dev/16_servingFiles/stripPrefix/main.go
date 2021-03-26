package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	// 현재 경로에서의 assets폴더를 웹서버의 exam폴더로 띄워준다. !
	http.Handle("/exam/", http.StripPrefix("/exam", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/exam/gopher.jpg">`)
}
