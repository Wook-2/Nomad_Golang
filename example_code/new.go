package main

import "fmt"

func maketen(i *int){
	*i = 10
}

func main(){

	ptr := new(int)
	fmt.Println(*ptr)
	maketen(ptr)
	fmt.Println(*ptr)



}
