package main

import (
	"fmt"
	"os"
)

func ShowText() {
	fmt.Println("Finalizando a manipulação do arquivo")
}

func main() {

	file, err := os.Create("./panic-defer-recover/defer/settings.txt")

	// defer file.Close()
	defer ShowText()

	if err != nil {
		panic(err)
	}

	_, err = file.Write([]byte("Teste"))

	if err != nil {
		panic(err)
	}
}
