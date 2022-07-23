/*
Canais são o Jeito Certo® de fazer sincronização e código concorrente.
- Eles nos permitem transmitir valores entre goroutines.
- Servem pra coordenar, sincronizar, orquestrar, e buffering.
- Na prática:
    - make(chan type, b)
- Canais bloqueiam:
    - Eles são como corredores em uma corrida de revezamento
    - Eles tem que "passar o bastão" de maneira sincronizada
    - Se um corredor tentar passar o bastão pro próximo, mas o próximo corredor não estiver lá...
    - Ou se um corredor ficar esperando receber o bastão, mas ninguém entregar...
    - ...não dá certo.
- Exemplos:
    - Poe um valor num canal e faz um print. Block.
        - Código acima com goroutine.
        - Ou com buffer. Via de regra: má ideia; é legal em certas situações, mas em geral é melhor sempre passar o bastão de maneira sincronizada.
- Interessante: ref/spec → types
- Código: 
    - Block: https://play.golang.org/p/dClS7vQlYE (não roda!)
    - Go routine: https://play.golang.org/p/ZbNCwUuiPi
    - Buffer: https://play.golang.org/p/32vYvCR7qn
    - Buffer block: https://play.golang.org/p/smeW6vigAT
    - Mais buffer: https://play.golang.org/p/Pe2pcboGiA
*/

package main

import (
	"fmt"
)

/*
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()

func main() {
	canal := make(chan int)
	canal <- 42
	fmt.Println(<-canal)
}
*/

// com goroutine
// func main() {
// 	canal := make(chan int)

// 	go func() {
// 		canal <- 42
// 	}()

// 	fmt.Println(<-canal)
// }


// com buffer
func main() {
	canal := make(chan int, 1)
	canal <- 42

	fmt.Println(<-canal)
	// fmt.Println(<-canal) fatal error
}