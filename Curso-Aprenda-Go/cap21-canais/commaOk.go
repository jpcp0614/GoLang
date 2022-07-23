/*
v, ok := ←chan
- Se receber valor: v, true
- Canal fechado, nada, etc.: zero v, false
- Agora vamos resolver o problema do exercício anterior usando comma ok.
*/

package main

import (
	"fmt"
)

// func main() {
// 	canal := make(chan int)

// 	go func() {
// 		canal <- 42
// 		close(canal)
// 	}()

// 	v, ok := <-canal
// 	fmt.Println(v, ok)

// 	v, ok = <-canal
// 	fmt.Println(v, ok)
// }


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
		case v := <-par:
			fmt.Println("Recebido do canal par:", v)
		case v := <-impar:
			fmt.Println("Recebido do canal impar:", v)
		case v, ok := <-quit:
			if !ok {
				fmt.Println("Deu zebra:", v)
			} else {
				fmt.Println("Comma ok:", v)
			}
			return
		}
	}
}
