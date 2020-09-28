package main

import "fmt"

func canIDrink(age int) bool{
  //declare variable in if
  if koreanAge := age+1; koreanAge<18{
    return false
  }
  return true
}

func main(){
  fmt.Println(canIDrink(16))
}
