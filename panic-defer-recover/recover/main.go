package main

import (
	"fmt"
	"os"
)

func ReadFile() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado")
		}
	}()

	_, err := os.Open("./panic-defer-recover/recover/settings.txt")

	if err != nil {
		panic("FileNotExist")
	}
}

func main() {

	ReadFile()

	fmt.Println("Fim...")
}
