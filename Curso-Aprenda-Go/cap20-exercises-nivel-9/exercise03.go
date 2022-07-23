/*
Utilizando goroutines, crie um programa incrementador:
    - Tenha uma variável com o valor da contagem
    - Crie um monte de goroutines, onde cada uma deve:
        - Ler o valor do contador
        - Salvar este valor
        - Fazer yield da thread com runtime.Gosched()
        - Incrementar o valor salvo
        - Copiar o novo valor para a variável inicial
    - Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
    - Demonstre que há uma condição de corrida utilizando a flag -race
*/

package main

import (
	"fmt"
	"sync"
	"runtime"
)

var wg sync.WaitGroup

const quantidadeDeGoroutines = 500
var contador int

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
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
		}()
	}
}