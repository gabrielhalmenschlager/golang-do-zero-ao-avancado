package main

import "fmt"

func main() {

	m := make(map[string]int)
	m["sp"] = 10000000
	m["cg"] = 900000
	m["rj"] = 6000000
	delete(m, "rj")

	for chave, valor := range m {
		fmt.Println("Cidade", chave, "H:", valor)
	}

}
