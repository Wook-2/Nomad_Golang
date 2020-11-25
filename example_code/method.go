package main

import "fmt"

type rectangle struct {
	height, width float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func main() {

	r := rectangle{10, 20}

	fmt.Println(r.area())

}
