package main

import (
	"fmt"
	"pacotes/matematica"
)

func main() {
	carro := matematica.Carro{
		Marca:      "chevrolet",
		Velocidade: 10,
	}
	carro.Andar()
	soma := matematica.Soma(1, 2.0)
	fmt.Println(soma)
}
