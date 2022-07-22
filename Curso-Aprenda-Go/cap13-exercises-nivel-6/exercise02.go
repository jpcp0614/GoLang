/*
Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
- Passe um valor do tipo slice of int como argumento para a função;
- Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
- Passe um valor do tipo slice of int como argumento para a função.
*/

package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(params(s...))
	fmt.Println(sliceInt(s))
}

func params(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}

func sliceInt(x []int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}