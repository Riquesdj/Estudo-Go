package main

import "fmt"

func main() {
	var variavel1 string = "Variavel tipo string"
	variavel2 := "Variavel 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "Variavel 3"
		variavel4 string = "Variavel 4"
	)

	fmt.Println(variavel3)
	fmt.Println(variavel4)

	variavel5, variavel6 := "Variavel5", "Variavel6"

	fmt.Println(variavel5)
	fmt.Println(variavel6)
}
