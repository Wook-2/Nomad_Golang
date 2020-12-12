package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	data := map[string]string{
		"AFruits":  "Apple",
		"BFruits2": "Banana",
		"CFlower":  "Catus",
		"DBoss":    "Diablo",
		"EName":    "Evan",
	}

	fmt.Println(data)
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
