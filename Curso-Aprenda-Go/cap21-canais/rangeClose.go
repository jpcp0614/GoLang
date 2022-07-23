/*
Range:
    - go func com for loop com send e close(chan)
    - recebe com range chan
*/

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go meuLoop(10, c)

	prints(c)
}


func meuLoop(t int, s chan<- int) {
	for i := 0; i < t; i++ {
		s <- i
	}
	close(s) // daqui não sai mais nada (encerra o canal)
}

func prints(r <-chan int) {
	for v := range r {
		fmt.Println("Recebido do canal:", v)
	}
}