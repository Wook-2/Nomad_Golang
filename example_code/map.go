package main

import "fmt"

func main() {
	v := map[string]int{"A": 10}

	val, ok := v["A"]

	if ok {
		fmt.Println("Value:", val)
	} else {
		fmt.Println("key doesn`t have value.")
	}

	a := v["A"]
	a = 20

	fmt.Println(v["A"], a)

}
