package main

import "fmt"

func imprimirEnderecoPosicaoArrays(array *[10]int) {

	for i := 0; i < len(*array); i++ {
		fmt.Printf("o valo da memória nessa posição é: %p\n", &array[i])
	}

}

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	imprimirEnderecoPosicaoArrays(&array)
}
