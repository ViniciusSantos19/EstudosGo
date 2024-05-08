package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")
	num1 := 12
	ponteiro := &num1

	*ponteiro = 11

	fmt.Println(num1)
	fmt.Println(ponteiro)
	fmt.Println(&num1)
	fmt.Println(&ponteiro)
}
