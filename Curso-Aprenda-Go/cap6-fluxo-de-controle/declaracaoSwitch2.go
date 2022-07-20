/*
Pode avaliar tipos
- Pode haver uma expressão de inicialização
*/

package main

import (
	"fmt"
)

var x interface{}

func main() {

	x = -175

	switch x.(type) {
		case int:
			fmt.Println("é um int")
		case bool:
			fmt.Println("é um bool")
		case string:
			fmt.Println("é uma string")
		case float64:
			fmt.Println("é um float")
	}
}