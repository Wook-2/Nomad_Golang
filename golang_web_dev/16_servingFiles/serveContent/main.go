package main

import (
	"io"
	"net/http"
	"os"
)

// servecontent는 잘 쓰이진 않는다고 한다.
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/byungwook.jpg", pic)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Contetn-Type", "text/html; charset=utg-8")
	io.WriteString(w, `<img src="/byungwook.jpg">`)
}

func pic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("byungwook.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}
