package main

import "fmt"

func testeCloujure() func() {
	texto := "texto dentro da funcao"

	return func() {
		fmt.Println(texto)
	}

}

func main() {
	texto := "texto do escopo dentro da main"
	fmt.Println(texto)
	pegaFuncao := testeCloujure()
	pegaFuncao()
}
