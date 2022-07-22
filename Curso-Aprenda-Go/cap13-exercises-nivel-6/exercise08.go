/*
Crie uma função que retorna uma função.
- Atribua a função retornada a uma variável.
- Chame a função retornada.
*/

package main

import (
	"fmt"
)

func retorno(x int) func(int) float64 {
	return func(y int) float64 {
		return float64(y + x)
	}
}

func main() {
	fmt.Println(retorno(2)(3))
}