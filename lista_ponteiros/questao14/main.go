package main

import (
	"fmt"
	"unsafe"
)

func preencherArrayComValor(arr *int, valor int, tam int) {
	for i := 0; i < tam; i++ {
		*arr = valor
		arr = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(arr)) + uintptr(unsafe.Sizeof(*arr))))

	}
}

func main() {
	var array [5]int
	tam := len(array)
	preencherArrayComValor(&array[0], 5, tam)
	fmt.Println(array)
}
