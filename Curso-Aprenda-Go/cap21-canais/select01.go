/*
Select é como switch, só que pra canais, e não é sequencial.
- "A select blocks until one of its cases can run, then it executes that case.
	It chooses one at random if multiple are ready." — https://tour.golang.org/concurrency/5
- Na prática:
    - Exemplo 1:
        - Duas go funcs enviando X/2 números cada uma pra um canal
        - For loop X valores, select case ←x
    - Exemplo 2:
        - Func 1 recebe X valores de canal, depois manda qualquer coisa pra chan quit
        - Func 2 for infinito, select: case envia pra canal, case recebe de quit
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

	c1 := make(chan int)
	c2 := make(chan int)

	x := 500

	go func(x int) {
		for i := 0; i < x; i++ {
			c1 <- i
		}
	}(x/2)

	go func(x int) {
		for i := 0; i < x; i++ {
			c2 <- i
		}
	}(x/2)

	for i := 0; i < x; i++ {
		select {
		case <-c1:
			fmt.Println("Recebido do canal 1:", i)
		case <-c2:
			fmt.Println("Recebido do canal 2:", i)
		}
	}
}