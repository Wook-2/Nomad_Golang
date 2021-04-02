package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	fmt.Println("Listen . . .")
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	c := req.FormValue("c")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<form method="post">
		<input type="text" name="q">
		<input type="checkbox" name="c">
		<input type="submit">
	</form>
	<br>`+v+c)

}

// form 양식의 method를 post로 하면 폼양식으로 데이터 전송 시 url이 바뀌지 않지만 get으로 해주면 ?q=input 이 url뒤에 붙게된다.
