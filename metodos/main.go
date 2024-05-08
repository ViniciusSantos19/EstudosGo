package main

import "fmt"

type ususario struct {
	nome  string
	idade uint8
}

func (u ususario) salvar() {
	fmt.Println("Dentro do método de salvar")
	fmt.Println("Usuáiro salvo no banco de dados")
}

func (u *ususario) atualizarIdade() {
	u.idade++
}

func main() {
	user := ususario{
		nome:  "Vinicius",
		idade: 22,
	}

	user.salvar()
	user.atualizarIdade()
	fmt.Println(user)
}
