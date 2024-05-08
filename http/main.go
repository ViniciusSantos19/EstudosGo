package main

import (
	"encoding/json"
	"fmt"
	"http/models"
	"io"
	"net/http"
	"os"
)

func main() {
	chamada, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println(err)
	}
	result, err := io.ReadAll(chamada.Body)
	fmt.Println(string(result))
	defer chamada.Body.Close()

	pessoa := models.Pessoa{
		Nome:           "Vinicius",
		Sobrenome:      "Rodrigues",
		DataNascimetno: "05/11/2001",
	}
	resultado, err := json.Marshal(pessoa)
	fmt.Println(string(resultado))
	json.NewEncoder(os.Stdout).Encode(pessoa)
	var pessoaX models.Pessoa
	err = json.Unmarshal(resultado, &pessoaX)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pessoaX)
}
