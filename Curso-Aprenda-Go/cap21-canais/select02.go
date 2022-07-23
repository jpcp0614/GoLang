/*
Exemplo 2:
        - Func 1 recebe X valores de canal, depois manda qualquer coisa pra chan quit
        - Func 2 for infinito, select: case envia pra canal, case recebe de quit
*/

package main

import (
	"fmt"
)

func main() {
	canal := make(chan int)
	quit := make(chan int)
	
	go recebeQuit(canal, quit)
	enviaParaCanal(canal, quit)

}

func recebeQuit(canal chan int, quit chan int) {
	for i := 0; i < 500; i++ {
		fmt.Println("Recebido:", canal)
	}
	quit <- 0
}

func enviaParaCanal(canal chan int, quit chan int) {
	qualquerCoisa := 10
	
	for {
		select {
			case canal <- qualquerCoisa:
				qualquerCoisa++
			case <-quit:
				return
		}
	}
}