/*
Crie um struct "erroEspecial" que implemente a interface builtin.error. 
- Crie uma função que tenha um valor do tipo error como parâmetro. 
- Crie um valor do tipo "erroEspecial" e passe-o para a função da instrução anterior.
*/

package main

import (
	"fmt"
)

type erroEspecial struct {
	qualquerCoisa string
}

func (e erroEspecial) Error() string {
	return "Deu ruim!!!"
}

func erroComoParametro(e error) {
	fmt.Println("Passaram esse erro como argumento:", e)
}

func main() {
	errEsp := erroEspecial{qualquerCoisa: "Que droga!!!"}

	erroComoParametro(errEsp)
}