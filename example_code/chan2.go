package main

import "fmt"

// Channel is First In First Out
func main() {

	c := make(chan int, 2)

	c <- 1
	c <- 2

	fmt.Println(<-c)
	fmt.Println(<-c)
}
