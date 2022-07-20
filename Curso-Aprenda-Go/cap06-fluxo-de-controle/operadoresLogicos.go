/*
&&
- ||
- !
- Qual o resultado de fmt.Println...
    - true && true
    - true && false
    - true || true
    - true || false
    - !true
*/

package main

import (
	"fmt"
)

func main() {
	x := 4

	if (x % 2 == 0 || x % 3 == 0) {
		fmt.Println("x é múltiplo de 2 ou 3")
	}

	if (x % 2 == 0 && x % 3 == 0) {
		fmt.Println("x é múltiplo de 2 e 3")
	}
}