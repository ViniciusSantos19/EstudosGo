package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func lerArquivo(file *os.File) error {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if scanner.Err() != nil {
		return scanner.Err()
	}
	return nil
}

func main() {
	/*	f, err := os.Create("arquvio.txt")
		if err != nil {
			log.Panicln(err)
		}
		defer f.Close()*/
	arquivo, err := os.OpenFile("arquivo.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 10; i++ {
		_, err := arquivo.WriteString(fmt.Sprintf("adicionou uma linha %d\n", i))
		if err != nil {
			fmt.Println(err)
			arquivo.Close()
			return
		}
	}

	// Mova o cursor de volta para o início do arquivo
	_, err = arquivo.Seek(0, io.SeekStart)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = lerArquivo(arquivo)
	if err != nil {
		fmt.Println("O erro foi aqui")
		fmt.Println(err)
	}
	defer arquivo.Close()
}
