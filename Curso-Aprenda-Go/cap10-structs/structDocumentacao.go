/*
É importante se familiarizar com a documentação da linguagem Go.
- Neste vídeo vamos ver um pouco sobre o que a documentação diz sobre structs.
- Veremos:
    - ref/spec
        - Já vimos mais da metade dos tipos em Go!
        - Struct types.
            - x, y int
            - anonymous fields
            - promoted fields
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

	// mais conciso
	pessoa3 := pessoa{"Maurício", 40}

	pessoa4 := profissional{pessoa{"Antônio", 42}, "Mecânico", 4000}

	fmt.Println(pessoa1.nome) // João
	fmt.Println(pessoa2.pessoa.idade) // 30
	fmt.Println(pessoa2.idade) // se não tiver conflito -> 30
	fmt.Println(pessoa3)
	fmt.Println(pessoa4)
}