package main

import "fmt"

type forma interface {
	CalcularArea() float32
}

type quadrado struct {
	lado float32
}

type trianguloRetangulo struct {
	base   float32
	altura float32
}

func (q quadrado) CalcularArea() float32 {
	return q.lado * q.lado
}

func (t trianguloRetangulo) CalcularArea() float32 {
	return (t.base * t.altura) / 2
}

func main() {
	q := quadrado{lado: 4.0}
	t := trianguloRetangulo{base: 3.0, altura: 5.0}
	fmt.Println(q.CalcularArea())
	fmt.Println(t.CalcularArea())

	fmt.Println("Chamando os método por interface")
	fmt.Println(forma.CalcularArea(q))
	fmt.Println(forma.CalcularArea(t))

}
