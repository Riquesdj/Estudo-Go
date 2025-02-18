package main

import "fmt"

func main() {

	var array1 [5]int
	array1[0] = 5

	fmt.Println(array1)

	array2 := []int{1, 2, 3, 4, 5, 6}

	slice1 := []int{10, 20, 30, 40, 50, 60, 70, 80}

	fmt.Println(slice1)

	slice1 = append(slice1, 90)
	fmt.Println(slice1)

	// Estamos adicionando a um novo slice uma parte do array utilizando o comando =array2[0:3]
	slice2 := array2[0:3]
	fmt.Println(slice2)

	// Array internos

	//	tipo do slice, quantiade de termos, máximo de termos
	fmt.Println("-----------------------------------------")
	slice3 := make([]float32, 10, 11)

	fmt.Println(slice3)

	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 8)
	fmt.Println(slice3)

	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

}
