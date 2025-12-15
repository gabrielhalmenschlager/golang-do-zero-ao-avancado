package main

import (
	"fmt"
	"golang-do-zero-ao-avancado/struct/models"
)

func main() {
	fmt.Println("Iniciando...")

	endereco := models.Endereco{
		Rua:    "Rua X",
		Numero: 200,
		Cidade: "Santa Cruz do Sul",
	}

	pessoa := models.Pessoa{
		Nome:     "Gabriel",
		Endereco: endereco,
	}

	endereco.Numero = 150

	fmt.Println(pessoa)

}
