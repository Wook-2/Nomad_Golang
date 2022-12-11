## Counting Sort 알고리즘

### 제한 조건만 맞으면 가장 빠른 정렬 알고리즘
제한 조건 맞는 경우가 드물긴함..  

Quick sort
- 최적: O(n)
- mean: O(nlogn)
- worst: O(n^2)

Merge sort
- best: O(nlogn)
- worst: O(nlogn)

Counting sort
- best: O(n+k) (k는 값의 범위)
- 따라서 k의 값이 작은경우에 효율적이다. k가 n^2이면 결국 O(n^2)이랑 다를거 없음 

`countSort1.go`

```go
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
```

`countSort2.go`

```go
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

```

`소문자로 이루어진 문자열에서 가장 많이나온 알파벳을 구하시오`

```go
package main

import "fmt"

// 문자열에서 가장 많이나온 알파벳을 구하시오
func main() {
	str := "sadashbhvxzcjvhcxzvptmdxxxx"

	var count [26]int
	for i := 0; i < len(str); i++ {
		count[str[i]-'a']++
		// fmt.Println(str[i] - 'a')
	}

	maxCount := 0
	var maxCh byte

	for i := 0; i < len(count); i++ {
		if count[i] >= maxCount {
			maxCount = count[i]
			maxCh = byte('a' + i)
		}
	}

	fmt.Printf("letter: %c\ncount: %d\n", maxCh, maxCount)
}

```

`학생들의 키가 소수점 한자리 까지 주어질 때 주어진 범위 사이에 속하는 학생들을 출력하시오`

```go
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

```

### 제한 조건만 맞으면 가장 빠른 정렬 알고리즘

제한 조건이 맞는 경우가 드물지만 `O(n + k)` 알고리즘이고 k가 작으면 가장 빠른 알고리즘이다.!
