/*
Deslocamento de bits é quando deslocamos digitos binários para a esquerda ou direita.
*/

package main

import (
	"fmt"
)

/*
const (
	a = iota + 100000
	_
	c
	d
	_
	f
)
*/

/*
func main() {
	y := 24

	x := y << 2

	z := y >> 2

	fmt.Printf("%b\n", y)
	fmt.Printf("%b\n", x)
	fmt.Printf("%b\n", z)
}
*/

const (
	_ = iota // 0
	KB = 1 << (iota * 10) // 1 << 1 * 10
	MB
	GB
	TB
)

func main() {
	fmt.Println("binary\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\tdecimal")
	fmt.Printf("%b\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t\t\t\t\t\t\t\t\t\t\t\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t\t\t\t\t\t\t\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t\t\t", TB)
	fmt.Printf("%d\n", TB)
}