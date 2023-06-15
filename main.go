package main

import (
	"fmt"
	"github.com/vk663434/pkg"
)
	
func main(){
  fmt.Println("calling function Getfloat")
  fmt.Print("please enter number : ")
  farhrenhit ,err := pkg.Getfloat()
  if err != nil {
    fmt.Println(err)
  }
  cilsius := (farhrenhit-32)*5/9
  fmt.Printf("result is %.3f",cilsius)
  
}