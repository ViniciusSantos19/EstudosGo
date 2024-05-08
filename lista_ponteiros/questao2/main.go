package main

import (
	"fmt"
)

func converteEmString(end *int) string {
	end_string := fmt.Sprintf("%p", end)
	return end_string
}

func comparaEnderecos(num1, num2 *int) string {
	if *num1 > *num2 {
		return converteEmString(num1)
	}
	return converteEmString(num2)
}

func main() {
	var a, b int

	a = 2
	b = 1

	fmt.Println(&a)
	fmt.Println(&b)
	fmt.Println(comparaEnderecos(&a, &b))

}
