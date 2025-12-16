package main

import (
	"fmt"
	"golang-do-zero-ao-avancado/exercicios2/exercicio1/models"
	"time"
)

func main() {

	var nomeDosItens []string
	nomeDosItens = append(nomeDosItens, "Arroz")
	nomeDosItens = append(nomeDosItens, "Feijão")
	nomeDosItens = append(nomeDosItens, "Carne")

	compra, err := models.NovaCompra("Mercadinho", time.Now(), nomeDosItens)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Mercado: %s\n", compra.Mercado)
		fmt.Printf("Data: %s\n", compra.Data.Format("02/01/2006 15:04"))
		fmt.Println("Itens:")

		for _, item := range compra.Itens {
			fmt.Printf("- %s\n", item.Nome)
		}
	}

}
