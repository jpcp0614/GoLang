/*
Struct é um tipo de dados composto que nos permite armazenar valores de tipos *diferentes.*
- Seu nome vem de "structure," ou estrutura.
- Declaração: type x struct { y: z }
- Acesso: x.y
- Exemplo: nome, idade, fumante.
*/

package main

import (
	"fmt"
)

type cliente struct {
	nome				string
	sobrenome		string
	fumante			bool
}

func main() {
	cliente1 := cliente{
		nome: "João",
		sobrenome: "Pereira",
		fumante: false,
	}

	fmt.Println(cliente1)
}