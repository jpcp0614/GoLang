/*
Escreva um programa que mostre um número em decimal, binário, e hexadecimal.
*/

package main

import (
	"fmt"
)

func main() {
	a := 300
	fmt.Println("decimal\t\tbinary\t\t\thexadecimal")
	fmt.Printf("%v\t\t\t\t%b\t\t\t%#x", a, a, a)
}