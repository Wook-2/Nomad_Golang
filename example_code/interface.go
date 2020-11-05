package main

import (
	"fmt"
	"math"
)

//Circle struct
type Circle struct {
	r float64
}

//Area return circle area
func (c Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

//Rect struct
type Rect struct {
	width  float64
	height float64
}

// Area return rect area
func (r Rect) Area() float64 {
	return r.height * r.width
}

// Shape interface
type Shape interface {
	Area() float64
}

func main() {

	var c = Circle{5}
	var r = Rect{5, 10}
	var s Shape = c

	fmt.Println("using interface: ", s.Area())
	fmt.Println("area of circle: ", c.Area())
	fmt.Println("area of rectangle: ", r.Area())

}
