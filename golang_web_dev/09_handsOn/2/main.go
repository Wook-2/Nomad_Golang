package main

import (
	"log"
	"os"
	"text/template"
)

// Add Menu contains Breakfast, Lunch, Dinner

type food struct {
	Name  string
	Price int
}
type restaurant struct {
	Breakfast []food
	Lunch     []food
	Dinner    []food
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	ramen := food{Name: "라면", Price: 5000}
	yogurt := food{Name: "요거트", Price: 2000}

	dongga := food{Name: "돈까스", Price: 8000}
	bob := food{Name: "회덮밥", Price: 11000}

	kim := food{Name: "김치찌개", Price: 15000}
	sabu := food{Name: "샤브샤브", Price: 20000}

	r1 := restaurant{
		Breakfast: []food{ramen, yogurt},
		Lunch:     []food{dongga, bob},
		Dinner:    []food{kim, sabu},
	}

	err := tpl.Execute(os.Stdout, r1)

	if err != nil {
		log.Fatalln(err)
	}
}
