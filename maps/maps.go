package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "vinicius",
		"sobrenome": "Rodrigues",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	fmt.Println("map dentro de map")
	usuario2 := map[string]map[string]string{
		"nome": {
			"Primeiro": "Paula",
			"Segundo":  "Adelaide",
		},
		"curso": {
			"Graducao": "Arquitetura",
			"Campus":   "Salvador",
		},
	}
	fmt.Println(usuario2)
	fmt.Println(usuario2["nome"])
	fmt.Println(usuario2["curso"]["Campus"])
}
