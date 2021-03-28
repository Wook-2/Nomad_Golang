package main

import (
	"log"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./serving")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
