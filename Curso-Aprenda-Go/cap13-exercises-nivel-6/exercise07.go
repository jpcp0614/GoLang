/*
Atribua uma função a uma variável.
- Chame essa função.
*/

package main

import (
	"fmt"
)

func variavel(word string) string {
	return word
}


func main() {
	x := variavel("Hello world")
	fmt.Println(x)
}