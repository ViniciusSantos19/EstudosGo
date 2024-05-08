package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

type Endereco struct {
	Rua    string
	Bairro string
	cidade string
	Estado string
	Cep    string
}

func main() {
	fmt.Println("Aula de structs")
	vinicius := Cliente{
		Nome:  "Vinicius",
		Idade: 22,
		Ativo: true,
		Endereco: Endereco{
			Rua:    "Praça João Amaral",
			Bairro: "Amaralina",
			cidade: "Salvador",
			Estado: "Bahia",
			Cep:    "41900205"},
	}
	vinicius.Endereco.cidade = "São Paulo"
	vinicius.Cep = "555111208"
	fmt.Println(vinicius)
}
