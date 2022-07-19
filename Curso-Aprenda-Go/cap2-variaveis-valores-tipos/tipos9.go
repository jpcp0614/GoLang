/*
Conversão de tipos é o que soa.
- Em Go não se diz casting, se diz conversion.
- a = int(b)
- ref/spec#Conversions
- Fim da sessão. Parabéns! Dicas, motivação e exercícios.
*/

package main

import (
	"fmt"
)

type hotDog int
var b hotDog = 10

func main() {
	x := 10
	fmt.Printf("%v, %T\n", x, x)

	x = int(b)
	fmt.Printf("%v, %T\n", x, x)
}