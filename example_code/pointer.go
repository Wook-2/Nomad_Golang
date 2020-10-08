package main

import "fmt"

func zero(a *int){
	*a = 0
}

func main(){

	x := 10
	zero(&x)
	fmt.Println(x)
}
