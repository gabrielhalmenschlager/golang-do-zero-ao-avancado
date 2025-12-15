package main

import "fmt"

func main() {

	var itens = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var somaUmACinco = 0
	var somaSeisADez = 0

	for i := 0; i < len(itens); i++ {
		if itens[i] <= 5 {
			somaUmACinco += itens[i]
		} else {
			somaSeisADez += itens[i]
		}
	}

	fmt.Println("A soma total de um a cinco é:", somaUmACinco)
	fmt.Println("A soma total de seis a dez é:", somaSeisADez)

	/* Professor

	lista := []int{2, 8, 3, 10, 5, 4, 7, 9}
	numeroAte5 = 0
	numeroAte10 = 0

	for i := 0; i < len(itens); i++ {
		if itens[i] <= 5 {
			numeroAte5 = numeroAte5 + lista[i]
		} else {
			numeroAte10 = numeroAte10 + lista[i]
		}
	}

	fmt.Println(numeroAte5)
	fmt.Println(numeroAte10)

	*/

}
