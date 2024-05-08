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
		return converteEmString(num1) + " " + fmt.Sprint("%d", *num1)
	}
	return converteEmString(num2) + " " + fmt.Sprint("%d", *num2)

}

func main() {
	var a, b int

	fmt.Print("Digite o primeiro número: ")
	fmt.Scanf("%d", &a)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanf("%d", &b)

	fmt.Println(comparaEnderecos(&a, &b))

}
