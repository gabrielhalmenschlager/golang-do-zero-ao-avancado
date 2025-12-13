package main

import "fmt"

func main() {

	texto := "Bom dia!"
	// texto = 15 não pode alterar o tipo
	texto = "Boa noite!"

	fmt.Println(texto)

	// Tipos de dados
	var texto2 string
	var numero int
	var preco float32
	var ehMasculino bool

	fmt.Println("\nSem Valores")
	fmt.Println(texto2)
	fmt.Println(numero)
	fmt.Println(preco)
	fmt.Println(ehMasculino)

	texto2 = "Boa tarde!"
	numero = 15
	preco = 17.5
	ehMasculino = true

	fmt.Println("\nCom Valores")
	fmt.Println(texto2)
	fmt.Println(numero)
	fmt.Println(preco)
	fmt.Println(ehMasculino)
}
