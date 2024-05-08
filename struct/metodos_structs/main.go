package main

import "fmt"

type Endereco struct {
	Rua        string
	Cep        string
	Cidade     string
	Estado     string
	Logradoruo string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (cliente *Cliente) Desativar() {
	cliente.Ativo = false
	fmt.Printf("O cliente %s foi desativado\n", cliente.Nome)
}

func main() {
	vinicius := Cliente{
		Nome:  "vinicius",
		Idade: 22,
		Ativo: true,
		Endereco: Endereco{
			Rua:        "Praca João Amaral",
			Cep:        "41900205",
			Cidade:     "Salvador",
			Estado:     "Bahia",
			Logradoruo: "Amaralina",
		},
	}
	fmt.Println(vinicius)
	vinicius.Desativar()
	fmt.Println(vinicius)
}
