package main

import "fmt"


func superAdd(numbers ...int) int{
  total := 0
  sum := 0
  
  //using range
  for _, number := range numbers{
    sum += number
  }
  
  //for loop
  for i:=0; i< len(numbers); i++{
    total += numbers[i]
  }
  defer fmt.Println("sum is", sum)
  return sum
  
}

func main(){
  result := superAdd(1, 2, 3, 4, 5)
  fmt.Println(result)
  
}
