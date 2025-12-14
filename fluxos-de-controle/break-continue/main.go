package main

import "fmt"

func main() {

	texto := "Palavra"
	fmt.Println("Quantidade: ", len(texto))

	for i := 0; i < len(texto); i++ {
		if string(texto[i]) == "r" {
			break
		}
		fmt.Println(string(texto[i]))
	}

}
