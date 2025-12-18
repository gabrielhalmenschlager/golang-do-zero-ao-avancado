package main

import "fmt"

func main() {

	fmt.Println("Iniciando...")

	var a interface{}
	a = 5
	fmt.Println(a)

	var lista []interface{}
	lista = append(lista, 10)
	lista = append(lista, true)
	lista = append(lista, "Teste")

	for _, valor := range lista {
		if v, ok := valor.(string); ok {
			fmt.Println(v + " string")
		} else {
			fmt.Println(valor)
		}
	}
}
