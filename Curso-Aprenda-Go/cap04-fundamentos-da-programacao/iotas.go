/*
- https://go.dev/ref/spec#Iota
 Numa declaração de constantes, o identificador iota representa números sequenciais.
- Na prática.
    - iota, iota + 1, a = iota b c, reinicia em cada const, _
*/

package main

import (
	"fmt"
)

/*
const (
	a = iota
	_ = iota
	c = iota
	d = iota
	_ = iota
)
*/

// outra forma de escrever
const (
	a = iota * 10
	_
	c
	d
	_
)

func main() {
	fmt.Println(a, c, d)
}