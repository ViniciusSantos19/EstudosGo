package main

import (
	"fmt"
	"unicode/utf8"
)

func ocorrenciaString(string1, string2 *string) bool {
	lenString1 := utf8.RuneCountInString(*string1)
	lenString2 := utf8.RuneCountInString(*string2)

	if lenString2 > lenString1 {
		return false
	}

	for i := 0; i <= lenString1-lenString2; i++ {
		substr := (*string1)[i : i+lenString2]

		if substr == *string2 {
			return true
		}
	}
	return false
}

func main() {
	string1 := "ola mundo"
	string2 := "ola"
	fmt.Println(ocorrenciaString(&string1, &string2))
	string1 = "fmt.Println"
	string2 = "ola"
	fmt.Println(ocorrenciaString(&string1, &string2))
}
