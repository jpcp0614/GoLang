/*
Utilize mutex para consertar a condição de corrida do exercício anterior.
*/

package main

import (
	"fmt"
	"sync"
	"runtime"
)

var wg sync.WaitGroup

const quantidadeDeGoroutines = 5000
var contador int

var mu sync.Mutex

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
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
	}
}