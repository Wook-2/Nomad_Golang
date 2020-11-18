package main

import "fmt"

func main(){
	for i := 0; i<5; i++{
		fmt.Println("%d", i)
		defer fmt.Println("defer: %d", i)
	}



}
