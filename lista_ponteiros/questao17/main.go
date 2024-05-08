package main

import "fmt"

func fracao(num float32, inteiro *int, frac *float32) {
	*inteiro = int(num)
	*frac = num - float32(*inteiro)
}

func main() {
	var real float32
	real = 123.4
	var inteiro int
	var frac float32
	fracao(real, &inteiro, &frac)
	fmt.Printf("Os valores são %f %f e %d\n", real, frac, inteiro)
}
