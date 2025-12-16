package models

import "time"

type Compra struct {
	Mercado string
	Data    time.Time
	Itens   []ItemDaCompra
}

type ItemDaCompra struct {
	Nome       string
	Quantidade int
}
