package main

import (
	"fmt"
	"strconv"
	"strings"
	//"time"
)

func testaCanal(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto + strconv.Itoa(i)
		//		time.Sleep(time.Second)
	}
	close(canal)
}

func upperCase(canal chan string, output chan string) {
	for message := range canal {
		texto := strings.ToUpper(message)
		output <- texto
	}
	close(output)
}

func main() {
	canal := make(chan string)
	output := make(chan string)
	go testaCanal("olá mundo", canal)
	go func() {
		// Aguarde até que testaCanal feche o canal
		//<-canal
		// Inicie a execução de upperCase
		upperCase(canal, output)
	}()

	for message := range output {
		fmt.Println(message)
	}

}
