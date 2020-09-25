package main

import "fmt"
import "strings"

// return multiple value
func lenAndUpper(name string) (int, string){
  return len(name), strings.ToUpper(name)
}

// return type
func multiply(a, b int) int{
  return a*b
}

// Get multiple argument
func repeatMe(words ...string){
  fmt.Println(words)
}

func main(){
  name := "wook"
  fmt.Println(multiply(2, 2))
  totalLen, upperName := lenAndUpper(name)
  // totalLen, _ := lenAndUpper(name) -> don`t get upperName
  repeatMe("wook", "son", "zzz")
  fmt.Println(totalLen, upperName)
}
