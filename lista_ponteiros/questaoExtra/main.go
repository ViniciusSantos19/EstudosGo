package main

import "fmt"

func main() {
	slice1 := make([]int, 10, 10)

	for i := 0; i < len(slice1); i++ {
		slice1[i] = i
	}

	slice2 := slice1[2:len(slice1)]
	fmt.Println(slice2)
}
