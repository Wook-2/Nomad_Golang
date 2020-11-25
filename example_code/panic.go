package main

import "fmt"

func test() int {
	arr := []int{}

	defer func() {
		v := recover()
		fmt.Println("panic recover: ", v)
	}()

	element := arr[5]
	return element
}

func main() {
	result := test()
	fmt.Println("result: ", result)
}
