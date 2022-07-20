/*
Demonstre o resto da divisão por 4 de todos os números entre 10 e 100
*/

package main

import (
	"fmt"
)

func main() {
	for x := 0; x <= 100; x++ {
		if (x % 4 == 0) {
			fmt.Printf("%v\n", x)
		}
		fmt.Printf("%v\n", x % 4)
	}
}