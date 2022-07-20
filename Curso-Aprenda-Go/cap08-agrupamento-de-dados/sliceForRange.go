/*
Slices:
    - Tamanho: len(x)
    - Índice específico: x[i] (0-based)
- Para ver todos os itens de uma slice utilizamos o loop for com range.
- Range significa alcance, faixa, extensão.
- For range: for i, v := range x {}
*/

package main

import (
	"fmt"
)

// func main() {
// 	slice := []string{"banana", "maçã", "laranja"}

// 	for index, valor := range slice {
// 		fmt.Println("No índice", index, "temos o valor:", valor)
// 	}
// }

func main() {
	slice := []int{10, 20, 30, 40, 50, 60}

	total := 0

	for _, valor := range slice {
		total += valor
	}
	fmt.Println("No slice, o valor total é:", total)
}