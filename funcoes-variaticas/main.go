package main

import "fmt"

func soma(numeros ...int) int {
	aux := 0
	for _, numero := range numeros {
		aux += numero
	}
	return aux
}

func main() {
	numeros := []int{1, 2, 3, 4, 5}
	resultado := soma(numeros...)
	fmt.Println(resultado)
}
