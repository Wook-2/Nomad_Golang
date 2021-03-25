package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/gopher.jpg", pic)

	fmt.Println("Serving ...")
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utg-8")
	io.WriteString(w, `<img src="/gopher.jpg">`)
}

func pic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "gopher.jpg")
}
