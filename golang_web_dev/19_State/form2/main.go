package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	fmt.Print("Serving ...")
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	f := req.FormValue("first")
	l := req.FormValue("last")
	s := req.FormValue("subscribe") == "on"

	err := tpl.ExecuteTemplate(w, "index.gohtml", person{f, l, s})
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}
