package main

import "fmt"

func main() {
	students := []struct {
		Name   string
		Height float64
	}{
		{Name: "jimmy", Height: 160.1},
		{Name: "chan", Height: 162.3},
		{Name: "ne", Height: 158.0},
		{Name: "gu", Height: 170.1},
		{Name: "keke", Height: 80.1},
	}

	var HeightMap [3000][]string

	for _, student := range students {
		idx := int(student.Height * 10)
		HeightMap[idx] = append(HeightMap[idx], student.Name)
	}

	start, end := 160, 180

	for i := start * 10; i < end*10; i++ {
		for _, name := range HeightMap[i] {
			fmt.Println("name: ", name, "height: ", float64(i)/10)
		}
	}
}
