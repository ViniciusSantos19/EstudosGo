package main

import "fmt"

func main() {
	var numero interface{} = 0
	var texto interface{} = "olá mundo"
	imprimeInterface(texto)
	imprimeInterface(numero)
}

func imprimeInterface(t interface{}) {
	fmt.Printf("o tipo da variável é %T o valor da variável é %v\n", t, t)
}
