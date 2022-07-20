/*
Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
    - Nome
    - Sobrenome
    - Sabores favoritos de sorvete
- Crie dois valores do tipo "pessoa" e demonstre estes valores,
	utilizando range na slice que contem os sabores de sorvete.
*/

package main

import (
	"fmt"
)

type pessoa struct {
	nome				string
	sobrenome		string
	sabores			[]string
}

func main() {
	pessoa1 := pessoa {
		nome: "João",
		sobrenome: "Pereira",
		sabores: []string{
			"Flocos",
			"Passas ao rum",
		},
	}

	pessoa2 := pessoa {
		nome: "Ana",
		sobrenome: "Paula",
		sabores: []string{
			"Napolitano",
			"Coco",
			"Chocolate",
		},
	}

	fmt.Printf("%v:\n", pessoa1.nome)
	for i, s := range pessoa1.sabores {
		fmt.Println("\t", i+1, "-", s)
	}

	fmt.Printf("%v:\n", pessoa2.nome)
	for i, s := range pessoa2.sabores {
		fmt.Println("\t", i+1, "-", s)
	}
}