package main

import "fmt"

func maiorValor(primeiro, segundo *int) {
	if *primeiro < *segundo {
		*primeiro, *segundo = *segundo, *primeiro
	}
}

func imprime(a, b *int) {
	maiorValor(a, b)
	fmt.Printf("O maior valor é %d\no menor valor é %d\n", *a, *b)
}

func main() {
	var a, b int
	fmt.Printf("Digite o primeiro valor:\n")
	fmt.Scanf("%d", &a)
	fmt.Printf("Digite o segundo valor:\n")
	fmt.Scanf("%d", &b)
	imprime(&a, &b)
}
