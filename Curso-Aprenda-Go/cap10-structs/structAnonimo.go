/*
São structs sem identificadores.
- x := struct { name type }{ name: value }
*/

package main

import (
	"fmt"
)

func main() {
	x := struct {
		nome	string
		idade int
	}{
		nome: "Pedro",
		idade: 20,
	}

	fmt.Println(x)
}