package main

import "fmt"

func imprimirArray(array *[5]int) {
	for i := 0; i < len(array); i++ {
		if array[i]%2 == 0 {
			fmt.Printf("O valor da posicao na memória é %p o valor da posicao é %d\n", &array[i], array[i])
		}
	}
}

func main() {
	array := [5]int{}
	for i := 0; i < len(array); i++ {
		fmt.Printf("Digite um valor para essa poisção\n")
		fmt.Scanf("%d", &array[i])
	}

	imprimirArray(&array)

}
