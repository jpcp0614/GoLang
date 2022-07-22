/*
Um método é uma função anexada a um tipo.
- Quando se anexa uma função a um tipo, ela se torna um método desse tipo.
- Pode-se anexar uma função a um tipo utilizando seu receiver.
- Utilização: valor.método()
- Exemplo: o tipo "pessoa" pode ter um método oibomdia()
*/


package main

import (
	"fmt"
)

type pessoa struct {
	nome	string
	idade	int
}

func (p pessoa) oibomdia() {
	fmt.Printf("oi, bom dia")
}

func main() {
	mauricio := pessoa{nome: "mauricio", idade: 30}

	mauricio.oibomdia()
}

// func (receiver) identifier(parameters) (returns) { code }