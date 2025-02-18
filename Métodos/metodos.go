package main

import "fmt"

type usuario struct {
	nome  string
	idade uint
}

func (u usuario) salvar() {
	fmt.Println(u.nome)
}

func (u usuario) verificaUsuarioMaiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {

	usuario1 := usuario{"Claudio", 46}
	fmt.Println(usuario1)
	usuario1.salvar()

	fmt.Println(usuario1.verificaUsuarioMaiorDeIdade())
	usuario1.fazerAniversario()
}
