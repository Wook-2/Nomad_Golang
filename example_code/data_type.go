package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	slc := []int{100, 200, 300}
	fmt.Println(strconv.Itoa(slc[0]))

	b := byte('a')
	fmt.Printf("%c\n", b)

	bs := []byte{'h', 'a'}
	fmt.Println("type:", reflect.TypeOf(bs))
	bs[0] = 'c'

	for i := range bs {
		fmt.Printf("%c", bs[i])
	}
	fmt.Printf("%s\n", bs)
	// fmt.Println(strconv.Itoa(slc))
	a := "a"
	str := "hello"
	bs = append(bs, str...)
	fmt.Printf("bs = %s\n", bs)
	bs[4] = 'w'
	fmt.Printf("bs = %s\n", bs)
	fmt.Println(string(bs[0]), string(bs[:4]), reflect.TypeOf(bs))
	bs[4] = 'a'

	copy(bs[:4], "aaaa")
	fmt.Printf("using copy() , bs = %s\n", bs)
	// str[0] = "a"
	// string은 값을 할당할 수 없음.

	// 형변환
	fmt.Println(string(str[0]))
	fmt.Println(int(a[0]))
}
