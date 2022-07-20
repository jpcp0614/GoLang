/*
Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
- Demonstre os valores do map utilizando range.
- Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
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

var meuMap map[string]pessoa

func main() {

	meuMap := make(map[string]pessoa)

	meuMap["Pereira"] = pessoa {
		nome: "João",
		sobrenome: "Pereira",
		sabores: []string{"Flocos", "Passas ao rum"},
	}

	meuMap["Paula"] = pessoa {
		nome: "Ana",
		sobrenome: "Paula",
		sabores: []string{"Napolitano", "Coco", "Chocolate"},
	}

	for _, value := range meuMap {
		fmt.Println("Meu nome é:", value.nome)
		for _, s := range value.sabores {
			fmt.Printf("%v\n", s)
		}
	}

	// outra maneira de fazer
	// meuMap2 := map[string]pessoa {
	// 	"Pereira": {
	// 		nome: "João",
	// 		sobrenome: "Pereira",
	// 		sabores: []string{"Flocos", "Passas ao rum"},
	// 	},
	// 	"Paula": {
	// 		nome: "Ana",
	// 		sobrenome: "Paula",
	// 		sabores: []string{"Napolitano", "Coco", "Chocolate"},
	// 	},
	// }

	// for _, value := range meuMap2 {
	// 	fmt.Println(value)
	// }
}