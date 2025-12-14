package main

import "fmt"

func main() {

	var listaTodaArray = [10]int{2, 10, 9, 4, 5, 8, 1, 3, 2, 11}
	fmt.Printf("%T \n", listaTodaArray)
	fmt.Println("Array: ", listaTodaArray)

	var listaTodaSlice = []int{2, 10, 9, 4, 5, 8, 1, 3, 2, 11}
	fmt.Printf("%T \n", listaTodaSlice)
	fmt.Println("Slice: ", listaTodaSlice)

}
