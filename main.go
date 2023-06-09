package main

import (
	"fmt"
	"math"
  
)

func squre(number float64) (float64,error){
  if number < 0 {
  return 0,fmt.Errorf("invalid number %0.2f",number)
  }
  return math.Sqrt(number), nil
}


func main() {
	result,err := squre(-4)
  if err != nil{
    fmt.Println(err)
  }
  fmt.Println(result)
}
