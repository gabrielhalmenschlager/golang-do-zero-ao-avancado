package models

import (
	"errors"
	"time"
)

type Compra struct {
	Mercado string
	Data    time.Time
	Itens   []ItemDaCompra
}

type ItemDaCompra struct {
	Nome string
}

func NovaCompra(mercado string, data time.Time, nomeDosItens []string) (*Compra, error) {

	if mercado == "" {
		return nil, errors.New("mercado é obrigátorio")
	}

	if len(nomeDosItens) == 0 {
		return nil, errors.New("itens são obrigátorio")
	}

	var itens []ItemDaCompra

	for _, nome := range nomeDosItens {
		itens = append(itens, ItemDaCompra{Nome: nome})
	}

	return &Compra{
		Mercado: "Mercado Super Legal",
		Data:    time.Now(),
		Itens:   itens,
	}, nil

}
