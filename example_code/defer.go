package main

import "fmt"

func main() {
	a := 10
	testDefer(&a)
	fmt.Println(a)
}

func add(i **int) {
	**i++
}

func testDefer(a *int) {
	defer add(&a)
	fmt.Println("defer:", *a)
}
