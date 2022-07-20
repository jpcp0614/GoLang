/*
Crie um programa que utilize a declaração switch, onde o switch statement
seja uma variável do tipo string com identificador "esporteFavorito".
*/

package main

import (
	"fmt"
)

func main() {
	esporteFavorito := "Fute"

	switch esporteFavorito {
		case "Vôlei":
			fmt.Println("Vôlei")
		case "Futebol":
			fmt.Println("Futebol")
		case "Natação":
			fmt.Println("Natação")
		default:
			fmt.Println("Não tenho esporte favorito")
		}
}