package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {

	generica("STRING")
	generica(1)
	generica(true)
}
