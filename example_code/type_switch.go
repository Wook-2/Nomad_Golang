package main

import "fmt"

func whatAmI(i interface{}) {
	switch t := i.(type) {
	case bool:
		fmt.Println("I`m a bool !")
	case int:
		fmt.Println("I`m a int !")
	default:
		fmt.Printf("Who are you %T?\n", t)
	}
}

func main() {

	whatAmI(true)
	whatAmI(100)
	whatAmI("byungwook")

}
