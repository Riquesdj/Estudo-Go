package main

import "fmt"

type pessoa struct {
	nome   string
	idade  uint8
	altura uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {

	p1 := pessoa{"Henrique", 23, 168}

	fmt.Println(p1)

	e1 := estudante{p1, "Tecnologia da informação", "Univesp"}

	fmt.Println(e1)

	fmt.Println(e1.nome)
}
