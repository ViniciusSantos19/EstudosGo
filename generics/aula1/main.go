package main

import (
	"fmt"
)

type Number interface {
	int | float64
}

type Pessoa[T Number] struct {
	nome      string
	sobrenome string
	saldo     T
}

func (p *Pessoa[T]) comparaSalario(num1, num2 T) T {
	if num1 > num2 {
		return num1
	}
	return num2
}

func soma[T Number](mapa map[string]T) T {
	var total T
	for _, valor := range mapa {
		total += valor
	}
	return total
}

func compara[T Number](mapa map[string]T) T {
	var maior T
	for _, v := range mapa {
		if v > maior {
			maior = v
		}
	}
	return maior
}

func compara2[T Number](num1, num2 T) bool {
	if num1 != num2 {
		return false
	}
	return true
}

func maior[T Number](num1, num2 T) bool {
	if num1 > num2 {
		return true
	}
	return false
}

func main() {
	mapa := make(map[string]int)
	mapa["um"] = 1
	mapa["dois"] = 2
	mapa["três"] = 3
	mapa["quatro"] = 4
	pessoa := Pessoa[float64]{
		nome:      "Vinicius",
		sobrenome: "Rodrigues",
		saldo:     1200,
	}
	pessoa2 := Pessoa[int]{
		nome:      "Paula",
		sobrenome: "Moreira",
		saldo:     12000,
	}
	fmt.Println(soma(mapa))
	fmt.Println(compara(mapa))
	fmt.Println(compara2(mapa["dois"], mapa["dois"]))
	fmt.Println(compara2(1, 1.0))
	fmt.Println(maior(1, 2))
	fmt.Println(pessoa.comparaSalario(pessoa.saldo, float64(pessoa2.saldo)))
}
