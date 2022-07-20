/*
Crie um programa que:
    - Atribua um valor int a uma variável
    - Demonstre este valor em decimal, binário e hexadecimal
    - Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
    - Demonstre esta outra variável em decimal, binário e hexadecimal
*/

package main

import (
	"fmt"
)

func main() {
	a := 200
	fmt.Printf("%v - %b - %#x", a, a, a)

	fmt.Println("")

	b := a << 1 // dividir ou multiplicar com deslocamento
	fmt.Printf("%v - %b - %#x", b, b, b)
}