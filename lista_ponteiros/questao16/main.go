package main

import "fmt"

func main() {
	var A int
	var B *int
	var C **int
	var D ***int

	A = 5
	fmt.Println(A)

	B = &A
	*B *= 2

	fmt.Println(A)
	C = &B
	**C *= 3
	fmt.Println(A)

	D = &C
	***D *= 4
	fmt.Println(A)
}
