package main

import (
	"fmt"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("Should be + f:%g", f)
	}
	return math.Sqrt(f), nil
}

func main() {
	sqrt, err := Sqrt(5)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("sqrt = ", sqrt)
}
