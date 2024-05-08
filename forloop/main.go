package main

import (
	"fmt"
)

func main() {
	slice := []string{"João", "davi", "Lucas"}
	mapa := make(map[string]string)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		fmt.Println("novo loop apenas para teste")
	}

	for indice, nome := range slice {
		fmt.Printf("O indice é: %d o nome é %s\n", indice, nome)
	}

	for _, nome := range slice {
		fmt.Println(nome)
	}

	mapa["vinicius"] = "estudante de análise e desenvolvimento de sistemas"
	mapa["Manoel"] = "professor de análise e desenvolvimento de sistemas"

	for chave, valor := range mapa {
		fmt.Println(chave)
		fmt.Println(valor)
	}

	i := 0

	for i < 10 {
		fmt.Println("Aumentou 1 ")
		i++
	}
}
