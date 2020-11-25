package main

import "fmt"

func main() {

	const (
		x = iota
		y
		z
	)

	const (
		a = iota
		_
		b
		c
	)

	fmt.Println(x, y, z)
	fmt.Println(a, b, c)

}
