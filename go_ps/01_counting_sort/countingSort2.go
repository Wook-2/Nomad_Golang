package main

import "fmt"

// 0~10의 값을 갖는 배열을 정렬하시오
func main() {
	arr := []int{5, 1, 3, 2, 4, 2, 6, 8, 2, 0, 4, 5, 1, 6, 7, 2, 9, 2, 10}

	var count [11]int

	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}

	sorted := make([]int, len(arr))

	for i := 1; i < 11; i++ {
		count[i] += count[i-1]
	}

	for i := 0; i < len(arr); i++ {
		sorted[count[arr[i]]-1] = arr[i]
		count[arr[i]]--
	}
	fmt.Println("arr: ", arr)
	fmt.Println("count: ", count)
	fmt.Println(sorted)
}
