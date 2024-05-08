package main

import "fmt"

func dobraSoma(num1, num2 *int) int {
	*num1 *= 2
	*num2 *= 2
	fmt.Println(num1)
	return *num1 + *num2
}

func main() {
	var num1, num2 int
	num1 = 12
	num2 = 7
	fmt.Println(dobraSoma(&num1, &num2))
	fmt.Printf("O novo valor de num1 é %d\n o novo valor de num2 é %d\n", num1, num2)
}
