package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n2 - n1
	return
}

func main() {

	soma, subtracao := calculosMatematicos(10, 5)

	fmt.Println(soma, subtracao)

}
