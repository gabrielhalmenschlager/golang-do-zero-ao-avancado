package main

import "fmt"

func main() {

	var listaToda = []int{2, 10, 9, 4, 5, 8, 1, 3}
	var segundaLista = listaToda[:3]
	var terceiraLista = listaToda[4:]
	var ultimoItem = listaToda[len(listaToda)-1:]

	fmt.Println("Lista Toda: ", listaToda)
	fmt.Println("Segunda lista: ", segundaLista)
	fmt.Println("Terceira lista: ", terceiraLista)
	fmt.Println("Ultimo item: ", ultimoItem)

}
