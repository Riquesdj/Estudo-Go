package main

import "fmt"

func main() {

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	var variavel3 int = 100
	var ponteiro *int = &variavel3

	fmt.Println(variavel3, ponteiro)

	variavel3++

	fmt.Println(variavel3, *ponteiro)
}
