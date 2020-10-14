package main

import "fmt"

func main() {
	a := 10
	test_defer(&a)
	fmt.Println(a)
}

func add(i **int) {
	**i++
}

func test_defer(a *int) {
	defer add(&a)
	fmt.Println("defer:", *a)
}
