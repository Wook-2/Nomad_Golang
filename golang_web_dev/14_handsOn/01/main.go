package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Bark! Bark! Bark!")
}

func me(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln("error in parsing template", err)
	}

	err = tpl.ExecuteTemplate(res, "tpl.gohtml", "byungwook")

	if err != nil {
		log.Fatalln("error executing template", err)
	}

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
