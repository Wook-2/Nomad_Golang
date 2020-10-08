package main

import "fmt"

func main(){
	defer func(){
		str := recover()
		fmt.Println("AA",str)
	}()
	panic("PANICCC")
}
