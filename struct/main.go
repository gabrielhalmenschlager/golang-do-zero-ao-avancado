package main

import "fmt"

type endereco struct {
	rua    string
	numero int
	cidade string
}

func main() {
	fmt.Println("Iniciando...")

	endereco := endereco{
		rua:    "Rua X",
		numero: 200,
		cidade: "Santa Cruz do Sul",
	}

	fmt.Println(endereco)

	endereco.numero = 150

	fmt.Println(endereco)

}
