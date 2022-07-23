/*
- Exemplo 3:
        - Chans par, ímpar, quit
        - Func send manda números pares pra um canal, ímpares pra outro, e fecha/quit
        - Func receive é um select entre os três canais, encerra no quit
        - Problema!
*/

package main

import (
	"fmt"
)

func main() {
	par := make(chan int)
	impar := make(chan int)
	quit := make(chan bool)

	go mandaNumeros(par, impar, quit)

	receive(par, impar, quit)
}

func mandaNumeros(par, impar chan int, quit chan bool) {
	for i := 0; i < 500; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
	close(par)
	close(impar)
	quit <- true
}

func receive(par, impar chan int, quit chan bool) {
	for {
		select {
		case <-par:
			fmt.Println("Recebido do canal par:", par)
		case <-impar:
			fmt.Println("Recebido do canal impar:", impar)
		case <-quit:
			return
		}
	}
}
