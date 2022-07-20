/*
Structs dentro de structs dentro de structs.
- Exemplo: um corredor de fórmula 1 é uma pessoa (nome, sobrenome, idade)
						e também* um competidor (nome, equipe, pontos).
*/

package main

import (
	"fmt"
)

type pessoa struct {
	nome	string
	idade	int
}

type profissional struct {
	pessoa
	titulo	string
	salario	int
}

func main() {
	pessoa1 := pessoa{
		nome: "João",
		idade: 34,
	}

	pessoa2 := profissional{
		pessoa: pessoa{
			nome: "Pedro",
			idade: 30,
		},
		titulo: "Pizzaiolo",
		salario: 2500,
	}

	fmt.Println(pessoa1)
	fmt.Println(pessoa2)
}