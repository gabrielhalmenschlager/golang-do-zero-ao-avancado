package main

import "fmt"

func main() {

	m := make(map[string]int)
	m["sp"] = 10000000
	m["cg"] = 900000

	valor, existe := m["rj"]
	if existe {
		fmt.Println(valor)
	} else {
		fmt.Println("chave não existe")
	}

	fmt.Println(m)

}
