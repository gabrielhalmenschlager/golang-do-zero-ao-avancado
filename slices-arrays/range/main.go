package main

import "fmt"

func main() {

	m := make(map[string]int)
	m["sp"] = 10000000
	m["cg"] = 900000

	for chave, valor := range m {
		fmt.Println("Cidade", chave, "H:", valor)
	}

}
