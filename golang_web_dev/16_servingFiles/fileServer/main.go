package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/gopher", gopher)
	fmt.Println("Serving...")
	http.ListenAndServe(":8080", nil)
}

func gopher(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="gopher.jpg">`)
}
