package main

import "fmt"

type person struct{
  name string
  age int
  favFood []string
}

func main(){

  wook := person{name:"byungwook", age:23, favFood:[]string{"tangerine", "ramen"}}
  fmt.Println(wook)
  fmt.Println(wook.name)

}
