package main

import (
	"fmt"
	"sync"
	"time"
)

func teste(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

func main() {
	var waitGroup sync.WaitGroup
	texto := "texto que vai rodar dentro da funcao"
	texto2 := "texto do main"
	waitGroup.Add(2)
	go func() {
		teste(texto)
		waitGroup.Done()
	}()

	go func() {
		teste(texto2)
		waitGroup.Done()
	}()

	waitGroup.Wait()
}
