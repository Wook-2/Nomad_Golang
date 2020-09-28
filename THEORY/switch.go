package main

import "fmt"


func canIDrink(age int) bool{
  // can declare variable at the head of switch
  switch koreanAge := age + 2; koreanAge{
    case 10:
      return true
    case 18:
      return false
  }
  return true

}

func main(){
  fmt.Println(canIDrink(16))
}
