/*
Crie e use um struct anônimo.
- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
*/

package main

import (
	"fmt"
)

func main() {
	s := struct {
		nome					string
		sabor					string
		local					[]string
		combinaCom		map[string]string
	} {
		nome: "Sorvete",
		sabor: "Flocos",
		local: []string{"Lanchonete", "Padaria"},
		combinaCom: map[string]string{
			"amendoim": "picado",
			"coco": "ralado",
			"cobertura": "chocolate",
		},
	}

	fmt.Println(s)
}