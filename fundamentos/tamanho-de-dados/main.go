package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var i int
	var f float64
	var s string
	var b bool

	fmt.Printf("Tamanho de 'i' (int): %d bytes\n", unsafe.Sizeof(i))
	fmt.Printf("Tamanho de 'f' (float64): %d bytes\n", unsafe.Sizeof(f))
	fmt.Printf("Tamanho de 's' (string): %d bytes\n", unsafe.Sizeof(s))
	fmt.Printf("Tamanho de 'b' (bool): %d bytes\n", unsafe.Sizeof(b))
}
