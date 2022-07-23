/*
Agora vamos fazer a mesma coisa, mas com atomic ao invés de mutex.
    - atomic.AddInt64
    - atomic.LoadInt64
*/

package main

import (
	"fmt"
	"sync"
	"runtime"
	"sync/atomic"
)

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var contador int64
	contador = 0
	totalDeGoroutines := 10

	var wg sync.WaitGroup
	wg.Add(totalDeGoroutines)

	// var mu sync.Mutex

	for i := 0; i < totalDeGoroutines; i++ {
		go func() {
			// mu.Lock()
			atomic.AddInt64(&contador, 1)
			// yield
			runtime.Gosched()

			// mu.Unlock()
			fmt.Println("Contador:\t", atomic.LoadInt64(&contador))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Valor final:", contador)
}