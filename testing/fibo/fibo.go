package fibo

import "fmt"

func GetFibo(n int) int {
	n1, n2 := 1, 1

	for i := 0; i < n-2; i++ {
		n1, n2 = n2, n1+n2
	}

	return n2
}

func GetFibo2(n int) int {
	var slice []int
	slice = []int{1, 1}

	for i := 0; i < n-2; i++ {
		slice = append(slice, slice[i]+slice[i+1])
	}

	return slice[len(slice)-1]
}

func PrintHello() {
	fmt.Println("Hello world!")
}
