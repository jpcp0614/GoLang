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

	for i := 0; i < totalDeGoroutines; i++ {
		go func() {
			v := contador
			// yield
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Valor final:", contador)
}