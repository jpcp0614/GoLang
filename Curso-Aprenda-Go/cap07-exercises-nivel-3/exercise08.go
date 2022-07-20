/*
Crie um programa que utilize a declaração switch, sem switch statement especificado.
*/

package main

import (
	"fmt"
)

func main() {
	x := -1

	switch {
		case x >= 0:
			fmt.Println("x é positivo")
		case x < 0:
			fmt.Println("x é negativo")
		}
}