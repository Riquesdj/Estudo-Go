package main

import (
	"errors"
	"fmt"
)

func main() {

	// Inteiros
	var inteiros int = 100000000000

	fmt.Println(inteiros)
	// Fim inteiros

	// Reais

	var real float32 = 45.22

	fmt.Println(real)

	var real2 float64 = 45555.22

	fmt.Println(real2)

	// FIM reais

	// Escrita

	var texto string = "TESTE STRING"
	fmt.Println(texto)

	// FIM Escrita

	//Booleanos

	var testeBool bool
	fmt.Println(testeBool)

	// FIM Booleanos

	//Erros

	var erro error = errors.New("Deu ruim aqui")
	fmt.Println(erro)
}
