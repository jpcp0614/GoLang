/*
Agora vamos resolver a race condition do programa anterior utilizando mutex.
- Mutex é mutual exclusion, exclusão mútua.
- Utilizando mutex somente uma thread poderá utilizar a variável contador de cada vez, e as outras deve aguardar sua vez "na fila."
- Na prática:
    - type Mutex
        - func (m *Mutex) Lock()
        - func (m *Mutex) Unlock()
- RWMutex
*/

package main

import (
	"fmt"
	"sync"
	"runtime"
)

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	contador := 0
	totalDeGoroutines := 1000

	var wg sync.WaitGroup
	wg.Add(totalDeGoroutines)

	var mu sync.Mutex

	for i := 0; i < totalDeGoroutines; i++ {
		go func() {
			mu.Lock()
			v := contador
			// yield
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Valor final:", contador)
}