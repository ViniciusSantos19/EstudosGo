package main

import "fmt"

func main() {
	fmt.Println("Cloujures em golang")

	total := func() int {
		return (1 + 2 + 3 + 4 + 5 + 6) * 2
	}()
	fmt.Println(total)
}
