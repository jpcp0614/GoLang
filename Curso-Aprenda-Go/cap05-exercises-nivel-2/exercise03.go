/*
Crie constantes tipadas e não-tipadas.
- Demonstre seus valores.
*/

package main

import (
	"fmt"
)

const (
	a = 100.0
	b int = 20
)

func main() {
	fmt.Printf("%v, %T\n%v, %T", a, a, b, b)
}