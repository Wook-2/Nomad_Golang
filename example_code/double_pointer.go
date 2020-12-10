package main

import "fmt"

func secondFunc(number **int) {
	**number += 10
}

func firstFunc(number *int) {
	secondFunc(&number)
}

func main() {
	number := 10

	firstFunc(&number)

	fmt.Println(number)
}
