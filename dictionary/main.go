package main

import (
	"fmt"

	"./mydict"
)

func main() {

	dictionary := mydict.Dictionary{"hello": "firstword"}
	definition, err := dictionary.Search("hello")
	if err == nil {
		fmt.Println(definition)
	} else {
		fmt.Println(err)
	}
	err2 := dictionary.Add("hi", "Greeting")
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(dictionary)
	dictionary.Delete("hi")
	fmt.Println(dictionary)

}
