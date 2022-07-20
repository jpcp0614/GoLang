/*
Isso tudo aqui a gente já viu:
- Toda slice tem um array subjacente.
- Um slice é: um ponteiro/endereço para um array, mais len e cap (que é o len to array).
- Exemplo:
    - x := []int{...números}
    - y := append(x[:i], x[:i]...)
    - pkg/builtin/#append: "If it has sufficient capacity,
			the destination is resliced to accommodate the new elements.
			If it does not, a new underlying array will be allocated."
    - Ou seja, y utiliza o mesmo array subjacente que x.
    - O que nos dá um resultado inesperado.
- Ou seja, bom saber de antemão pra não ter que aprender na marra.
*/

package main

import (
	"fmt"
)

func main() {
	primeiroSlice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(primeiroSlice)

	segundoSlice := append(primeiroSlice[:2], primeiroSlice[4:]...)
	fmt.Println(segundoSlice)

	// array subjacent
	fmt.Println(primeiroSlice)
}