package main

import "fmt"

func main() {
	arrayNormal := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println(arrayNormal)
	fmt.Printf("len(%d) cap(%d)\n", len(arrayNormal), cap(arrayNormal))
	fmt.Printf("len(%d) cap(%d)\n", len(arrayNormal[:0]), cap(arrayNormal[:0]))
	fmt.Println(arrayNormal[:0])
	fmt.Printf("len(%d) cap(%d)\n", len(arrayNormal[2:]), cap(arrayNormal[2:]))
	fmt.Println(arrayNormal[2:])
	arrayNormal = append(arrayNormal, 5)
	fmt.Println(arrayNormal)
	fmt.Printf("len(%d), cap(%d)", len(arrayNormal), cap(arrayNormal))
}
