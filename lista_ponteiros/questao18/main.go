package main

import (
	"fmt"
	"math"
)

func calcularAreaEVolume(raio float64, area *float64, volume *float64) {
	*area = 4 * 3.14 * math.Pow(raio, 2)
	*volume = (4 / 3) * 3.14 * math.Pow(raio, 3)
}

func main() {
	raio := 3.6
	var volume, area float64
	calcularAreaEVolume(raio, &area, &volume)
	fmt.Printf("A área é %f o volume é %f\n", area, volume)
}
