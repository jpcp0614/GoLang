/*
Demonstre o funcionamento de um closure.
- Ou seja: crie uma função que retorna outra função, onde esta outra função faz uso de uma variável alem de seu scope.
*/

package main

import (
	"fmt"
)

func first() func() {
	x := 0
	return func() {
		x++
		fmt.Println(x)
	}
}

func main() {
	a := first()
	a()
	a()
	a()
	b := first()
	b()
	b()
	b()
}