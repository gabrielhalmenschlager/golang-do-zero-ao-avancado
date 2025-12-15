package main

import "fmt"

func main() {
	ImprimirMensagem()
	ImprimirMensagem2("Olá ")
}

func ImprimirMensagem() {
	fmt.Println("Hello, World!")
}

func ImprimirMensagem2(mensagem string) {
	mensagem += "Bom dia"
	fmt.Println(mensagem)
}
