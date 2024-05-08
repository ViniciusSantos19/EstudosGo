package main

import "fmt"

type pessoa struct {
  nome string
  idade int32
  data_nascimento string
}

type aluno struct {
  pessoa
  id_aluno int32
  curso string
}

/*func main(){

  fmt.Println("'Herença'")

  pessoa1:= pessoa{"João", 18,"12/5/2001"}

  aluono1 := aluno{pessoa1, 12, "ads"}

  fmt.Println(pessoa1)
  fmt.Println(aluono1)

}*/
