package main

import "fmt"

func main() {

	lista := [2]int{6, 5}
	somaTotal := 0

	for i := 0; i < len(lista); i++ {
		somaTotal += lista[i]
	}

	fmt.Println("Soma total: ", somaTotal)

	/* Professor

	lista := []int {2, 3}
	somaTotal := lista[0] + lista[1]
	fmt.Println(somaTotal)

	*/

}
