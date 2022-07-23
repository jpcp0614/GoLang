package main

import (
	"fmt"
)

func main() {
	fmt.Println(Soma(1, 2, 3, 4, 5))
	fmt.Println(multiplicacao(1, 2, 3, 4, 5))
}

func Soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func multiplicacao(i ...int) int {
	total := 1
	for _, v := range i {
		total *= v
	}
	return total
}