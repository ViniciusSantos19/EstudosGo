package main

import "fmt"

func dobrandoValorArray(array *[5]int) {
	fmt.Println("Funcao que dobra o valor das posicoes de uma array")
	for i := 0; i < len(array); i++ {
		array[i] *= 2
	}
}

func imprimirArray(array *[5]int) {
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
}

func main() {
	array := [5]int{}
	for i := 0; i < len(array); i++ {
		fmt.Printf("Digite um valor para essa poisção\n")
		fmt.Scanf("%d", &array[i])
	}
	imprimirArray(&array)
	dobrandoValorArray(&array)
	imprimirArray(&array)
}
