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
