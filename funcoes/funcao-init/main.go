package main

import "fmt"

func init() {
	fmt.Println("Chamando a função init 1")
}

func init() {
	fmt.Println("Chamando a função init 2")
}

func init() {
	fmt.Println("Chamando a função init 3")
}

func main() {
	soma, subtracao, multiplicacao, divisao := Operacao(2, 3)

	fmt.Println("O resultado da soma é", soma)
	fmt.Println("O resultado da subtração é", subtracao)
	fmt.Println("O resultado da multiplicação é", multiplicacao)
	fmt.Println("O resultado da divisão é", divisao)
}

func Operacao(numero1, numero2 int) (soma int, subtracao int, multiplicacao int, divisao int) {
	soma = numero1 + numero2
	subtracao = numero1 - numero2
	multiplicacao = numero1 * numero2
	divisao = numero1 / numero2

	return
}

func init() {
	fmt.Println("Chamando a função init 4")
}

func init() {
	fmt.Println("Chamando a função init 5")
}
