package main

import (
	"Ola_mundo/soma"
	"fmt"
)

func main() {
	fmt.Println("ola mundo")
	resultado := soma.Soma(3, 2)
	fmt.Println(resultado)
	resultado = soma.Subtracao(4, 3)
	fmt.Println(resultado)
	resultado = soma.Subtracao(3, 5)
	fmt.Println(resultado)
}
