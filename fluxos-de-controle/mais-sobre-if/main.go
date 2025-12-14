package main

import "fmt"

func main() {
	var ehCarro bool
	ehCarro = true
	var valorDoAutomovel = 1000.00

	if ehCarro == true {
		valorDoAutomovel += 155.50
		fmt.Println("É carro")
	} else {
		valorDoAutomovel += 20.50
		fmt.Println("Não é carro")
	}

	fmt.Println("Valor final: ", valorDoAutomovel)

}

/*
	< menor
	> maior
	<= menor igual
	=> maior igual

*/
