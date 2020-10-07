package main

import "fmt"

func main(){
	a := 10
	test_defer(&a)
	fmt.Println(a)
}

func add(i **int){
	**i ++
}

func test_defer(a *int){
	//defer add(&a)
	defer *a = 15
	fmt.Println("defer:",*a)
}
