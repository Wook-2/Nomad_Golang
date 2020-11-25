package main

import "fmt"

func sum(a []int, c chan int) {
	total := 0
	// fmt.Println(a)
	for _, v := range a {
		total += v
	}

	c <- total
}

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c := make(chan int)
	go sum(arr[:5], c)
	go sum(arr[5:], c)

	x := <-c
	y := <-c

	fmt.Println(x, y)
}
