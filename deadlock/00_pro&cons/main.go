package main

import (
	"fmt"
	"time"
)

func plus(c chan int) {
	val := <-c
	c <- val + 1
}

func minus(c chan int) {
	val := <-c
	c <- val - 1
}

func display(c <-chan int) {
	for {
		val := <-c
		fmt.Println(val)
		time.Sleep(time.Second)
	}

}

func main() {
	var c chan int = make(chan int)

	c <- 1

	go plus(c)
	go minus(c)
	go display(c)

	var input string
	fmt.Scanln(&input)

}
