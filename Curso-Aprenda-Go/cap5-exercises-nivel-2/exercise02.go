/*
Escreva expressões utilizando os operadores, e atribua seus valores a variáveis.
- Demonstre estes valores.

- =
- !=
- ==
- <=
- >=
- <
- >

*/

package main

import (
	"fmt"
)

const (
	x = 10
	y = 20
)

func main() {
	a := (x == y)
	b := (x <= y)
	c := (x >= y)
	d := (x < y)
	e := (x > y)
	f := (x != y)
	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", a, b, c, d, e, f)
}