package main

import (
	"fmt"
	enderecos "introducao/Enderecos"
)

func main() {
	TipoEndereco := enderecos.TipoDeEndereco("Rua antonio ambuba")

	fmt.Println(TipoEndereco)
}
