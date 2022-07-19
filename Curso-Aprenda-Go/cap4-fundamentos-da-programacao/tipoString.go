/*
Strings são sequencias de bytes.
- Imutáveis.
- Uma string é um "slice of bytes" (ou, em português, uma fatia de bytes).
- Na prática:
    - %v %T
    - Raw string literals
    - Conversão para slice of bytes: []byte(x)
    - %#U, %#x
*/

package main

import (
	"fmt"
)

func main() {
	s := "Hello, World"
	fmt.Printf("%v, \n%T", s, s)

	// sb := []byte(s)

	// o range numa string me dá caractere por caractere
	for _, v := range s {
		fmt.Printf("%v - %T - %#U - %#x\n", v, v, v, v)
	}

	fmt.Println(" ")

	// aqui ele dá a resposta por byte
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v - %T - %#U - %#x\n", s[i], s[i], s[i], s[i])
	}


}