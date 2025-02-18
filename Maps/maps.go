package main

import "fmt"

func main() {

	fmt.Println("----MAPS----")

	// [Tipo da Chave] - Valor da chave{

	usuario := map[string]string{
		"nome":      "Pedro",
		"Sobrenome": "Carlos",
	}
	fmt.Println(usuario)

	usuario2 := map[string]map[string]string{
		"nome": {
			"Primeiro Nome": "Carlos",
			"Segundo Nome":  "Jorge",
		},
		"curso": {
			"Nome do Curso": "Engenharia",
			"Campus":        "Campus IFSP",
		},
	}

	fmt.Println(usuario2)

}
