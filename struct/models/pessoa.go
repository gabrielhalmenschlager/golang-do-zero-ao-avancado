package models

import (
	"time"
)

type Pessoa struct {
	Nome             string
	Endereco         Endereco
	DataDeNascimento time.Time
	Idade            int
}

// Método
func (p *Pessoa) CalculaIdade() {
	anoDeNascimento := p.DataDeNascimento.Year()
	anoAtual := time.Now().Year()
	p.Idade = anoAtual - anoDeNascimento
}

/* Função
func IdadeAtual(p Pessoa) int {
	anoDeNascimento := p.DataDeNascimento.Year()
	anoAtual := time.Now().Year()
	return anoAtual - anoDeNascimento
}
*/
