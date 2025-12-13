package main

import (
	"fmt"
	"reflect"
)

func main() {

	num1 := 10
	num2 := 5
	result := 0

	result = num1 + num2
	fmt.Println(result)

	result = num1 - num2
	fmt.Println(result)

	result = num1 * num2
	fmt.Println(result)

	result = num1 / num2
	fmt.Println(result)

	result = num1 % num2
	fmt.Println(result)

	fmt.Println(reflect.TypeOf(result))

	texto1 := "Texto 1"
	texto2 := "Texto 2"

	result2 := texto1 + texto2
	fmt.Println(result2)

	fmt.Println(reflect.TypeOf(result))
}
