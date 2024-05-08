package main

import "fmt"

func soma(num1 *int, num2 int) {
	soma := *num1 + num2
	*num1 = soma
}

func main() {
	var a, b int
	a = 12
	b = 5
	fmt.Printf("Valores das variáveis são a:%d e b:%d\n", a, b)

	soma(&a, b)

	fmt.Printf("Valores das variáveis são a:%d e b:%d\n", a, b)

	soma(&b, a)

	fmt.Printf("Valores das variáveis são a:%d e b:%d\n", a, b)

}
