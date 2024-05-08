package main

import (
	"fmt"
	"sync"
)

func multiplyByTwo(c chan int, wg *sync.WaitGroup) {
	for value := range c {
		result := value * 2
		fmt.Println("O valor multiplicado por 2 é:", result)
	}
	wg.Done() // Signal completion
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	c := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		c <- i
	}

	close(c)

	go multiplyByTwo(c, &wg)

	wg.Wait()
}
