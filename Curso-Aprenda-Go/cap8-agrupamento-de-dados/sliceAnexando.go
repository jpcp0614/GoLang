/*
Effective Go: append (package builtin)
- x = append(slice, ...values)
- x = append(slice, slice...)
- Todd: unfurl → desdobrar, desenrolar
- Nome oficial: enumeration
*/

package main

import (
	"fmt"
)

func main() {
	umaSlice := []int{1, 2, 3, 4}
	outraSlice := []int{9, 10, 11, 12}

	umaSlice = append(umaSlice, 5, 6, 7, 8)
	fmt.Println(umaSlice)

	fmt.Println(outraSlice)

	// ... pegar o conteúdo (items da fatia)
	outraSlice = append(outraSlice, umaSlice...)
	fmt.Println(outraSlice)
}