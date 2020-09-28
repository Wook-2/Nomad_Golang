package main

import "fmt"

func main(){
  // length is limited
  names := [5]string{"wook", "wook2"}
  fmt.Println(names)
  // slice type doesn`t limit array length
  slice_name := []string{"a", "b"}
  fmt.Println(slice_name)
  ar := append(slice_name, "zzz")
  fmt.Println("slice_name:", slice_name)
  fmt.Println("ar:", ar)

}
