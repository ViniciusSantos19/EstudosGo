package main

import (
	"fmt"
	"unsafe"
)

func imprimirArray(array *int, tam int) {
	for i := 0; i < tam; i++ {
		fmt.Printf("O valor dessa posicao %d é %d\n", i, *array)
		array = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(array)) + uintptr(unsafe.Sizeof(*array))))
	}
}

func iniciarArray(array []int) {
	for i := 0; i < len(array); i++ {
		array[i] = i + 1
	}
}

func main() {
	slice := make([]int, 10)
	iniciarArray(slice)
	imprimirArray(&slice[0], len(slice))
}
