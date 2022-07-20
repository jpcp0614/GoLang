/*
- For
    - Repetição hierárquica
    - Exemplos: relógio, calendário 
*/

package main

import (
	"fmt"
)

// func main() {
// 	for horas := 0; horas <= 12; horas ++ {
// 		fmt.Println("Hora: ", horas)
// 		for x := 0; x < 60; x++ {
// 			fmt.Print(" ", x)
// 		}
// 		fmt.Println("")
// 	}
// }

func main() {
	for mes := 1; mes <= 12; mes ++ {
		fmt.Println("Mês:", mes)
		for x := 1; x <= 30; x++ {
			fmt.Print(" ", x)
		}
		fmt.Println("")
	}
}