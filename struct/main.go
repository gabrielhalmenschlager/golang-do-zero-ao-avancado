package main

import (
	"fmt"
	"golang-do-zero-ao-avancado/struct/models"
	"time"
)

func main() {
	fmt.Println("Iniciando...")

	endereco := models.Endereco{
		Rua:    "Rua X",
		Numero: 200,
		Cidade: "Santa Cruz do Sul",
	}

	pessoa := models.Pessoa{
		Nome:             "Gabriel",
		Endereco:         endereco,
		DataDeNascimento: time.Date(2007, 12, 14, 0, 0, 0, 0, time.Local),
	}

	fmt.Println(pessoa)
	fmt.Println(endereco)

	// idade := models.CalculaIdade(pessoa) Usando a Função

	pessoa.CalculaIdade()
	fmt.Println(pessoa.Idade)

}
