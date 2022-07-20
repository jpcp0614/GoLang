/*
São valores imutáveis.
- Podem ser tipadas ou não:
    - const oi = "Bom dia"
    - const oi string = "Bom dia"
- As não tipadas só terão um tipo atribuido a elas quando forem usadas.
    - Ex. qual o tipo de 42? int? uint? float64?
    - Ou seja, é uma flexibilidade conveniente.
- Na prática: int, float, string.
    - const x = y
    - const ( x = y )
*/

package main

import (
	"fmt"
)

// dá pra declarar constante assim:
const (
	i = 20
	j = 40
	k = 60
)

const x = 10
var y = 10

var c int
var d float64

func main() {
	c = y
	fmt.Printf("%v - %T", c, c) // sendo usado, int

	fmt.Println(" ")

	//d = y // não pode
	d = x
	fmt.Printf("%v - %T", d, d)
}