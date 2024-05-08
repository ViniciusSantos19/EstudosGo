package main

import "fmt"



func generica(inter interface{}) {
  fmt.Println(inter)
}

func main() {
  numero :=1
  array := []int{1,2,3,4,5,6,7,8,9,10}
  generica(array)
  generica(numero)
}
