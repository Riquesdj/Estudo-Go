package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup
	//Diz quantas goroutines vão ser sincronizadas no ()
	waitGroup.Add(2)

	go func() {
		escrever("Olá mundo")
		waitGroup.Done() // Diminui 1 na contagem da waitGroup
	}()
	go func() {
		escrever("Programando em GO")
		waitGroup.Done() // Diminui 1 na contagem da waitGroup
	}()

	//Aguarda o contador chegar a 0 para encerrar o código
	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
