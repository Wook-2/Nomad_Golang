// naked return: 함수 선언과 동시에 return값을 정해주어 새로 변수 정의할필요 x, return ~~명시해줄필요 x
// defer: defer ~~를 통해 함수가 끝난 후 명령을 실행할 수 있음

package main

import (
  "fmt"
  "strings"
)
// "naked return" : new way to return values .
// "defer" : excute something when function is done. !!
func lenAndUpper(name string) (length int, upperName string){
  defer fmt.Println("I'm done") // excute after return . wow 
  length = len(name)
  upperName = strings.ToUpper(name)
  return
}


func main(){
  length, upperName := lenAndUpper("wook")
  fmt.Println(length, upperName)
  
}
