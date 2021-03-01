package main

import (
	"io"
	"net/http"
)

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Bark! Bark! Bark!")
}

func me(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello Byungwook!")
}

func index(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Index Page")
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
