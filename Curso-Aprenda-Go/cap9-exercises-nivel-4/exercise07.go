/*
Crie uma slice contendo slices de strings ([][]string).
Atribua valores a este slice multi-dimensional da seguinte maneira:
    - "Nome", "Sobrenome", "Hobby favorito"
- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
*/

package main

import (
	"fmt"
)

func main() {
	slice := [][]string{
		{"João", "Pereira", "Ler"},
		{"Gustavo", "Henrique", "Correr"},
		{"Ana", "Luísa", "Assistir séries"},
	}

	fmt.Printf("Nome\t\tSobrenome\t\tHobby favorito\n")
	for _, i := range slice {
		for _, value := range i {
			fmt.Printf("%v\t", value)
		}
	}
}