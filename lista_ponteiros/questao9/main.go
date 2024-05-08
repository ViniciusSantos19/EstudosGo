package main

import "fmt"

func imprimeMatriz(matriz *[3][3]int) {
	colunas := len(matriz[0])
	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[0]); j++ {
			fmt.Printf("o valor dessa posição %d da matrix é %p\n", i*colunas+j+1, &matriz)
		}
	}
}

func main() {
	matrix := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	imprimeMatriz(&matrix)
}
