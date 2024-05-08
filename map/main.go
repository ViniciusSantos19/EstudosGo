package main

import "fmt"

func main() {
	fmt.Println("aula sobre maps")
	salarios := map[string]int32{"wesley": 1000, "joão": 1200, "Paula": 1000}
	fmt.Println(salarios)
	delete(salarios, "wesley")
	fmt.Println(salarios)
	salario_com_make := make(map[string]int32)
	salario_com_make["gustavo fring"] = 3000
	salario_com_make["Renato Cariani"] = 1000000
	fmt.Println(salario_com_make)
	for nome, salario := range salario_com_make {
		fmt.Printf("O nome é %s e o salário é %d\n", nome, salario)
	}

	for _, salario := range salario_com_make {
		fmt.Printf("O salário é %d\n", salario)
	}

}
