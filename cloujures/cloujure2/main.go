package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func soma_recursiva(num, vezes int) int {
	if vezes == 0 {
		return 0
	}
	return num + soma_recursiva(num, vezes-1)
}

func main() {
	fmt.Println("Teste de clouures")
	adicionar := adder()

	/*	for i := 0; i < 10; i++ {
		fmt.Println(adicionar(i))
	}*/
	fmt.Println(adicionar(4))
	fmt.Println(adicionar(2))
	fmt.Println(soma_recursiva(5, 2))
	fmt.Println(soma_recursiva(5, 0))
	fmt.Println(soma_recursiva(5, 5))

}
