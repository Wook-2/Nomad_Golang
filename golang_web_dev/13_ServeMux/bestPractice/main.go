package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "bark! bark! bark!")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "meow meow meow")
}

func main() {
	http.HandleFunc("/dog/", d) // "/dog/"뒤에 다른 path가 추가되어도 catch해줌
	http.HandleFunc("/cat", c)

	http.ListenAndServe(":8080", nil)
}
