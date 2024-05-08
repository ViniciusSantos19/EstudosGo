package main

import "fmt"

func frequenciaMaior(array []int, maior *int, maisFrequente *int) {

	contagem := make(map[int]int)

	for _, num := range array {
		contagem[num]++
	}

	fmt.Println(contagem)

	var maiorChave, maiorValor, chaveMaisFrequente int

	for chave, valor := range contagem {
		if chave > maiorChave {
			maiorChave = chave
		}

		if valor > maiorValor {
			maiorValor = valor
			chaveMaisFrequente = chave
		}

	}
	*maisFrequente = chaveMaisFrequente
	*maior = maiorChave
}

func main() {
	array := []int{1, 1, 2, 3, 4, 5, 1, 6, 7, 8}
	var maior, maisFrequente int
	frequenciaMaior(array, &maior, &maisFrequente)
	fmt.Printf("O maior valor é %d o valor mais frequente é %d", maior, maisFrequente)
}
