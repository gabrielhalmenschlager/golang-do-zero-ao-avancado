package main

import "fmt"

func main() {

	texto := "Palavra"
	fmt.Println("Quantidade: ", len(texto))
	tamanho := len(texto)
	i := 0

	for i < tamanho {
		if string(texto[i]) == "r" {
			break
		}
		fmt.Println(string(texto[i]))
		i++
	}

}
