/*
Utilize atomic para consertar a condição de corrida do exercício #3.
*/

package main

import (
	"fmt"
	"sync"
	"runtime"
	"sync/atomic"
)

var wg sync.WaitGroup

const quantidadeDeGoroutines = 500
var contador int64

func main() {
	criarGoroutines(quantidadeDeGoroutines)
	wg.Wait()
	fmt.Println("Total de goroutines:", quantidadeDeGoroutines)
	fmt.Println("Contador:", contador)
}

func criarGoroutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			atomic.AddInt64(&contador, 1)

			runtime.Gosched()
			
			fmt.Println("Contador:\t", atomic.LoadInt64(&contador))
			wg.Done()
		}()
	}
}