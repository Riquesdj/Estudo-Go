package main

import "fmt"

// func + nomeDaFunção(parametros, parametros) o que função retorna
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// func + nomeDaFunção(parametros, parametros) (tipoQueFunçãoRetorna, tipoQueFunçãoRetorna)
func CalculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao

}

func main() {

	var m1 int8 = 8
	var m2 int8 = 8

	fmt.Println(somar(m1, m2))

	var soma = somar(10, 20)
	println(soma)

	// criar um váriavel recebendo um função
	var f = func(txt string) {

		fmt.Println(txt)

	}

	f("Texto da funcao f")

	fmt.Println(CalculosMatematicos(10, 5))
	// ou
	resultadoSoma, _ := CalculosMatematicos(25, 10)
	fmt.Println(resultadoSoma)

}
