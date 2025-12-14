package main

import "fmt"

func main() {
	var ehCarro bool
	ehCarro = true
	var valorDoAutomovel = 1000.00

	if ehCarro == true {
		valorDoAutomovel = valorDoAutomovel + 155.50
		fmt.Println("É carro")
	}

	if ehCarro == false {
		valorDoAutomovel = valorDoAutomovel + 55.50
		fmt.Println("É moto")
	}

	fmt.Println("Valor final: ", valorDoAutomovel)

}

/*
	< menor
	> maior
	<= menor igual
	=> maior igual

*/
