package matematica

import "fmt"

type Numerical interface {
	int | float64
}

type Carro struct {
	Marca      string
	Velocidade float64
}

func Soma[T Numerical](num1, num2 T) T {
	return num1 + num2
}

func (c *Carro) Andar() {
	fmt.Println("o carro andou 10 metros")
}
