package main

import (
	"fmt"
	"time"
)

func teste(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

func main() {
	texto := "texto que vai rodar dentro da funcao"
	texto2 := "texto do main"
	go teste(texto)
	teste(texto2)
}
