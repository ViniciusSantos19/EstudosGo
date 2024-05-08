package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", BuscaCep)
	http.ListenAndServe(":8080", nil)
}

func BuscaCep(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("dentro do método")
	if request.URL.Path != "/" {
		fmt.Println("erro no path")
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	cepParam := request.URL.Query().Get("cep")
	if cepParam == "" {
		fmt.Println("cep vazio")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("olá mundo"))
}
