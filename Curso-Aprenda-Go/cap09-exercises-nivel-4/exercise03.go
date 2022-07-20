/*
Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
    - Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
    - Do quinto ao último item do slice (incluindo o último item!)
    - Do segundo ao sétimo item do slice (incluindo o sétimo item!)
    - Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
    - Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item
*/

package main

import (
	"fmt"
)

func main() {
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	slice13 := slice[:3]
	fmt.Println(slice13)

	slice510 := slice[4:]
	fmt.Println(slice510)

	slice27 := slice[1:7]
	fmt.Println(slice27)

	slice39 := slice[2:9]
	fmt.Println(slice39)

	slicePenultimo := slice[len(slice) - 2]
	fmt.Println(slicePenultimo)

	slice39Penultimo := slice[2:len(slice) - 1]
	fmt.Println(slice39Penultimo)
}