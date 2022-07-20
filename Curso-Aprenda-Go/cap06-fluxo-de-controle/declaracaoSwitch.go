/*
Switch:
    - pode avaliar uma expressão 
        - switch statement == case (value)
        - default switch statement == true (bool)
    - não há fall-through por padrão
    - criando fall-through
    - default
    - cases compostos
*/

package main

import (
	"fmt"
)

// func main() {
// 	x := 4

// 	switch {
// 		case x < 5:
// 			fmt.Println("x é menor que 5")
// 		case x == 5:
// 			fmt.Println("x é igual 5")
// 		case x > 5:
// 			fmt.Println("x é maior que 5")
// 	}
// }

// func main() {
// 	x := 4

// 	switch {
// 		case x <= 5:
// 			fmt.Println("x é menor que 5")
// 			// criar o fall through
// 			fallthrough
// 		case x == 5:
// 			fmt.Println("x é igual 5")
// 		case x > 5:
// 			fmt.Println("x é maior que 5")
// 		default:
// 			fmt.Println("x não é um número")
// 	}
// }

// func main() {
// 	x := "Maria"

// 	switch x {
// 		case "Zezinho", "Maria":
// 			fmt.Println("Time 1")
// 		case "Joana", "Pedrinho":
// 			fmt.Println("Time 2")
// 		default:
// 			fmt.Println("Não tem time hoje")
// 	}
// }

func main() {
	x := 5

	switch {
		case (x == 4), (x == 8):
			fmt.Println("é 4 ou 8")
		case (x < 10):
			fmt.Println("x é menor que 10")
	}
}