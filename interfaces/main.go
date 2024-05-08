package main

import "fmt"

type Pessoa interface {
	Desativar()
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func (cliente *Cliente) Desativar() {
	cliente.Ativo = false
	fmt.Printf("O cliente %s foi desativado", cliente.Nome)
}

func Destativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	vinicius := Cliente{
		Nome:  "vinicius",
		Idade: 22,
		Ativo: true,
	}

	Destativacao(&vinicius)
	fmt.Println(vinicius)
}
