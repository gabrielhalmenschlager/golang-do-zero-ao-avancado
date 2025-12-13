package main

import (
	"fmt"
	"strconv"
)

func main() {

	var numero int8 = 127
	var numeroInt int
	numeroInt = int(numero)

	fmt.Println(numero)
	fmt.Println(numeroInt)

	var preco float64 = 15.5
	var preco2 int = int(preco)

	fmt.Println(preco)
	fmt.Println(preco2)

	b, _ := strconv.ParseBool("true")
	fmt.Printf("%T", b)
	fmt.Println(b)

	texto := "42.55"
	i, _ := strconv.ParseFloat(texto, 64)
	fmt.Printf("%T \n", i)
	fmt.Println(i)
}
