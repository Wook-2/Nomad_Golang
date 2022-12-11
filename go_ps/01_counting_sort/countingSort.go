package main

import "fmt"

// 0~10의 값을 갖는 배열을 정렬하시오
func main() {
	arr := []int{5, 1, 3, 2, 4, 2, 6, 8, 2, 0, 4, 5, 1, 6, 7, 2, 9, 2, 10}

	var count [11]int

	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}

	sorted := make([]int, 0, len(arr))

	for i := 0; i < 11; i++ {
		for c := 0; c < count[i]; c++ {
			sorted = append(sorted, i)
		}
	}
	fmt.Println("count: ", count)
	fmt.Println(sorted)
}
