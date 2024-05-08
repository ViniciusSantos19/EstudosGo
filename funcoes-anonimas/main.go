package main

import "fmt"

func main() {
	pegaString := func(texto string) string {
		return texto
	}("olá mundo")

	fmt.Println(pegaString)
}
