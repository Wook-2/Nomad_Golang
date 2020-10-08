package main

import "fmt"

func main(){

	res := factorial(6)
	fmt.Println(res)
}

func factorial(x int) int{

	if x == 0{
		return 1
	}
	return x * factorial(x-1)
}
