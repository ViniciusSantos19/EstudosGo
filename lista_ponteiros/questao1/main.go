package main

import "fmt"

func main() {
	interio := 1
	ponteiro_interio := &interio
	real := 2.4
	ponteiro_real := &real
	character := 'c'
	ponteito_character := &character

	//imprimindo caracteres
	fmt.Println(interio)
	fmt.Println(real)
	fmt.Printf("%c\n", character)

	//alterando valores das variaveis com ponterios
	*ponteiro_interio = 12
	*ponteiro_real = 12.9
	*ponteito_character = 'a'

	//imprimindo caracteres
	fmt.Println(interio)
	fmt.Println(real)
	fmt.Printf("%c\n", character)
}
