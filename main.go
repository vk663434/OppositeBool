package main

import (
	"fmt"
	_ "math"
)

func boolOpposite(myboolen *bool) {
   *myboolen =! *myboolen 
  
  
}

func main(){
  truth := true
  boolOpposite(&truth)
  fmt.Println(truth)
  
}