package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array1 [5]int32
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	array1[3] = 4
	array1[4] = 5

	fmt.Println(array1)

	array2 := [...]int32{1, 2, 3, 4, 5, 5}
	fmt.Println(array2)

	slice1 := []int32{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(slice1)

	fmt.Println(reflect.TypeOf(slice1))
	fmt.Println(reflect.TypeOf(array2))

	slice1 = append(slice1, 18)
	fmt.Println(slice1)

	slice2 := array2[0:2]

	fmt.Println(slice2)

	slice3 := make([]int32, 10, 15)

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	fmt.Println("Slice 4")
	slice4 := make([]int32, 5)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
	slice4 = append(slice1, 5)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
	fmt.Println(slice4)
	fmt.Println(slice4)

}
