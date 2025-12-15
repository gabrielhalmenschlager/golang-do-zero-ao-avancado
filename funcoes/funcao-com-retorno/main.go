package main

import "fmt"

func main() {
	resultado := Soma(2, 3)
	fmt.Println("O resultado da soma é: ", resultado)
}

func Soma(numero1, numero2 int) int {
	resultado := numero1 + numero2
	return resultado
}
