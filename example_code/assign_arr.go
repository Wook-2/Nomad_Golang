package main

import "fmt"

func main() {
	// 선언
	var a []byte = make([]byte, 5)
	// a := []byte{'h', 'e'}

	copy(a, "aa")
	fmt.Println(a)

	a = []byte("hello world")

	fmt.Printf("%s\n", a)

	// this is exactly what I want!

	var temp []byte = make([]byte, 0)

	str := "hello golang!"

	temp = []byte(str)
	fmt.Printf("%s\n", temp)

}
