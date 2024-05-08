package main

import (
	"fmt"
)

func ordenarValores(num1, num2, num3 int) (int, []int) {
	valores := []int{num1, num2, num3}

	for i := 0; i < len(valores)-1; i++ {
		for j := i + 1; j < len(valores); j++ {
			if valores[i] > valores[j] {
				valores[i], valores[j] = valores[j], valores[i]
			}
		}
	}

	if valores[0] == 0 && valores[1] == 0 && valores[2] == 0 {
		return 1, valores
	}
	return 0, valores

}

func main() {
	array := [3]int{}

	for i := 0; i < len(array); i++ {
		fmt.Printf("Digite o valor dessa posicao do array\n")
		fmt.Scanf("%d", &array[i])
	}
	numero, array2 := ordenarValores(array[0], array[1], array[2])
	fmt.Println(numero)
	fmt.Println(array2)
}
