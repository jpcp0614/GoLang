/*
Variável declarada em um code block é undefined em outro
- Para variáveis com uma abrangência maior, package level scope, utilizamos `var`
- Funciona em qualquer lugar
- Prestar atenção: chaves, colchetes, parênteses
*/

package main

import (
	"fmt"
)

var y = 10

func main() {
	z := 20 // declarada dentro do code block
	qualquerCoisa(z)
}


func qualquerCoisa (x int) {
	fmt.Println(y)
	fmt.Println(x)
}