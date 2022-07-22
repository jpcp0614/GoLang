/*
Crie um tipo struct "pessoa" que contenha os campos:
    - nome
    - sobrenome
    - idade
- Crie um método para "pessoa" que demonstre o nome completo e a idade;
- Crie um valor de tipo "pessoa";
- Utilize o método criado para demonstrar esse valor.
*/

package main

import (
	"fmt"
)

type pessoa struct {
	nome				string
	sobrenome		string
	idade				int
}

func (p pessoa) nomeComp() {
	fmt.Println("nome completo:", p.nome, p.sobrenome)
}

func main() {
	pedro := pessoa {
		nome: "Pedro",
		sobrenome: "Henrique",
		idade: 42,
	}

	pedro.nomeComp()


}