/*
- 2. Com throttling! Ou seja, com um número máximo de go funcs.
        - Idem acima, mas a func que lança go funcs é assim:
        - Cria X go funcs, cada uma com um range no primeiro canal que, para cada item, poe o retorno de trabalho() no canal dois.
*/

package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	canal1 := make(chan int)
	canal2 := make(chan int)
	funcoes := 10

	go manda(20, canal1)
	go outra(funcoes, canal1, canal2)

	for v := range canal2 {
		fmt.Println(v)
	}
}

func manda(n int, canal chan int) {
	for i := 0; i < n; i++ {
		canal <- i
	}
	close(canal)
}

func outra(funcoes int, canal1, canal2 chan int) {
	var wg sync.WaitGroup

	for i := 0; i < funcoes; i++ {
		wg.Add(1)
		go func(x int) {
			for v := range canal1 {
				canal2 <- trabalho(v)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	close(canal2)
}

func trabalho(n int) int {
	time.Sleep(1000 * time.Millisecond)
	return n
}