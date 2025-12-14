package main

import "fmt"

func main() {
	salario := 990.90
	var salarioMaisOBonus float32
	salarioMaisOBonus = float32(salario)

	if salario < 1000 {
		salarioMaisOBonus = salarioMaisOBonus + 100
	}

	fmt.Println("Salário", salarioMaisOBonus)
}
