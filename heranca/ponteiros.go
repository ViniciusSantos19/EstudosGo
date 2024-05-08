package main

import "fmt"

func main() {
	var numero1 int32 = 12
	var numero2 int32 = numero1

	fmt.Println(numero1, numero2)

	numero1++

	fmt.Println(numero1, numero2)

	//ponteiros
	var numero3 int = 2
	var ponteiro *int = &numero3

	fmt.Println(numero3, ponteiro)

	*ponteiro++
	fmt.Println(numero3, ponteiro)
}
