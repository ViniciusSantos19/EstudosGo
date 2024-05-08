package soma

func Soma(numero int32, numero2 int32) int32 {
	return (numero + numero2)
}

func Subtracao(numero int32, numero2 int32) int32 {
	if numero2 > numero {
		return numero2 - numero
	}
	return numero - numero2
}
