package main

import (
	"fmt"
	"golang-do-zero-ao-avancado/exercicios2/exercicio1/models"
	"time"
)

func main() {

	var itens []models.ItemDaCompra
	itens = append(itens, models.ItemDaCompra{Nome: "Arroz", Quantidade: 2})
	itens = append(itens, models.ItemDaCompra{Nome: "Feijão", Quantidade: 6})
	itens = append(itens, models.ItemDaCompra{Nome: "Batata", Quantidade: 8})

	compra := models.Compra{
		Mercado: "Mercado Super Legal",
		Data:    time.Now(),
		Itens:   itens,
	}

	fmt.Printf("Mercado: %s\n", compra.Mercado)
	fmt.Printf("Data: %s\n", compra.Data.Format("02/01/2006 15:04"))
	fmt.Println("Itens:")

	for _, item := range compra.Itens {
		fmt.Printf("- %s - Quantidade: %d\n", item.Nome, item.Quantidade)
	}
}
