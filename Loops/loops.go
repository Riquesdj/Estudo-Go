package main

import "fmt"

func main() {

	// for fazendo a função do while

	/*i := 0

	for i < 10 {
		time.Sleep(time.Second)
		i++
		fmt.Println("Incrementando i")

	}


	for j := 0; j < 10; j++ {

		fmt.Println("Incrementando j", j)
		//time.Sleep(time.Second)

	*/
	/*
		nomes := [3]string{"João", "Guilherme", "Henrique"}

		// o termo do array, o valor do array := range nomeDoArray
		for indice, nome := range nomes {
			fmt.Println(indice, nome)
		}

		for _, letra := range "PALAVRA" {
			fmt.Println(string(letra))
		}

	*/
	usuario := map[string]string{
		"Nome":      "Henrique",
		"Sobrenome": "Sousa",
	}

	for nome, sobrenome := range usuario {
		fmt.Println(nome, sobrenome)
	}
}
