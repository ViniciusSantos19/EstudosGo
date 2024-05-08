package main

import "fmt"

func retornaDia(num int) string {
	switch num {
	case 1:
		return "domingo"
	case 2:
		return "segunda-feira"
	case 3:
		return "terca-feira"
	case 4:
		return "quarta-feira"
	case 5:
		return "quinta-feira"
	case 6:
		return "sexta-feria"
	case 7:
		return "sábado"
	default:
		return "número maior que 7 ou menor que 1"

	}
}

func main() {
	fmt.Println(retornaDia(1))
}
