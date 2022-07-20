/*
Operação módulo: %
- For: break
- For: continue
*/

package main

import (
	"fmt"
)

// func main() {
// 	x := 0

// 	for x < 20 {
// 		x++

// 		if (x % 2 != 0) {
// 			// faz isso quando o número não é par
// 			// continue quebra a iteração do loop ao meio
// 			// e segue para o próximo passo
// 			continue
// 		}
// 		fmt.Println(x)
// 	}
// }

func main() {
	x := 0

	for x < 20 {
		x++

		if (x == 3) {
			// faz isso quando o número não é par
			continue
		}
		fmt.Println(x)
	}
}