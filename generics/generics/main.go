package main

import "fmt"

func main() {

	slice1 := []int{5, 1, 2, 3}
	slice2 := []string{"a", "b", "c", "d"}

	newInts := reverse(slice1)
	newStrings := reverse(slice2)

	fmt.Println(newInts)
	fmt.Println(newStrings)
}

func reverse[T int | string](slice []T) []T {

	newInts := make([]T, len(slice))
	newIntsLen := len(slice) - 1

	for i := 0; i < len(slice); i++ {
		newInts[newIntsLen] = slice[i]
		newIntsLen--
	}
	return newInts
}
