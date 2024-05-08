package main

import (
	"fmt"
	"time"
)

func teste(texto string, canal chan string) {
	time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}

func main() {
	canal := make(chan string)
	go teste("olá mundo", canal)
	fmt.Println("Depois da funcao de teste")
	for mensagem := range canal {
		fmt.Println(mensagem)
	}
	fmt.Println("fim do programa")
}
