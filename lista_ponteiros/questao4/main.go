package main

import "fmt"

func troca(a, b *int) {
	aux := *a
	*a = *b
	*b = aux
}

func main() {
	fmt.Println("Troca com ponteiros")
	var a, b int
	a = 1
	b = 2
	fmt.Printf("O valor do primeiro número é %d\n", a)
	fmt.Printf("O valor do segundo número é %d\n", b)

	troca(&a, &b)

	fmt.Printf("O valor do primeiro número é %d\n", a)
	fmt.Printf("O valor do segundo número é %d\n", b)
}
