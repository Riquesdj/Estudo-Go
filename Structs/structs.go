package main

import "fmt"

type endereço struct {
	lougadoro string
	numero    uint8
}

type usuario struct {
	nome     string
	idade    uint8
	endereço endereço
}

func main() {

	var u usuario
	u.nome = "Henrique"
	u.idade = 23

	fmt.Println(u)

	e2 := endereço{"Rua pavão dourado", 130}

	u2 := usuario{"Leticia", 25, e2}
	fmt.Println(u2)
}
