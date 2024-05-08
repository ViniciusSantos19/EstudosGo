package main

import (
	"fmt"
	"log"
	"os"
	"projetoBuscaCepCli/app"
)

func main() {
	fmt.Println("Ponto de partido da aplicação")
	aplicacao := app.Gerar()

	if err := aplicacao.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
