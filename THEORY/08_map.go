package main

import "fmt"

func main(){

  wook := map[string]string{"name":"byungwook", "age":"23"}
  fmt.Println(wook)

  for key, value := range(wook){
    fmt.Println(key, value)
  }

}
