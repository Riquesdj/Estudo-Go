package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	// passando o texto e o canal que nós criamos logo acima
	go escrever("Olá mundo", canal)

	// := range canal serve para que enquanto o canal estiver aberto ele irá executar o comando
	for mensagem := range canal {
		fmt.Println(mensagem)
	}
}

// função recebe uma string e um canal que só recebe valores string
func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		// seta apontada para o canal quer dizer que ele esta recebendo o valor
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}
