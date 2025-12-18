package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println("Iniciando...")

	p := problemaDeNetwork{
		rede:     true,
		hardware: false,
	}

	ExibiError(errors.New("a error"))
	ExibiError(p)
}

type problemaDeNetwork struct {
	rede     bool
	hardware bool
}

func (p problemaDeNetwork) Error() string {
	if p.rede {
		return "problema de rede"
	} else if p.hardware {
		return "problema de hardware"
	} else {
		return "outro problema"
	}
}

func ExibiError(err error) {
	fmt.Println(err.Error())
}
