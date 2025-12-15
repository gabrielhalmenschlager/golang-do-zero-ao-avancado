package main

import "fmt"

func main() {

	x := 5
	y := &x
	*y = 10

	ImprimirValor(&x, y)

	fmt.Println("\n===== Main =====")
	fmt.Println("O valor de X é", x)
	fmt.Println("O valor de Y é", *y)

	fmt.Println("A Referência de memoria de X é", &x)
	fmt.Println("A Referência de memoria de Y é", y)

}

func ImprimirValor(x, y *int) {

	*x = 20

	fmt.Println("\n===== Imprimir valores =====")

	fmt.Println("O valor de X é", x)
	fmt.Println("O valor de Y é", y)

	fmt.Println("A Referência de memoria de X é", &x)
	fmt.Println("A Referência de memoria de Y é", &y)
}
