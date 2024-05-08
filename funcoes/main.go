package main

import (
	"fmt"
	"strconv"
)

func soma(a, b int32) int32 {
	return a + b
}

func test_void(texto string) {
	fmt.Println(texto)
}

func teste_funcao_com_dois_retorno(a, b int32) (int32, string) {
	if b > a {
		return b - a, "O segundo número era maior que o primeiro"
	}
	return a - b, "O primeiro número era maior que o segundo"
}

func main() {
	fmt.Println("Atividade sobre funcoes no golang")
	resultado := soma(1, 2)
	fmt.Printf("O resultado da soma foi %d\n", resultado)
	resultadoString := strconv.Itoa(int(resultado))
	test_void(resultadoString)
	resultadoSubtracao, texto := teste_funcao_com_dois_retorno(7, 1)
	fmt.Printf("%s\no resultado da subtracao foi %d\n", texto, resultadoSubtracao)

}
