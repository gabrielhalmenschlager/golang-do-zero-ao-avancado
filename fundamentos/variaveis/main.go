package main

import "fmt"

func main() {

	var peso int
	peso = 65
	var mensagem string
	mensagem = "Seu peso é 65"

	fmt.Println(peso)
	fmt.Println(mensagem)
	fmt.Printf("Seu peso é %d Kg", peso)

	idade := 17
	mensagem2 := "Sua idade é 17"

	fmt.Println(idade)
	fmt.Println(mensagem2)
	fmt.Printf("Você tem %d anos", idade)
}
