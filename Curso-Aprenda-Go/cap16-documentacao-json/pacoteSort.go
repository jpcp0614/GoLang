/*
Sort serve para ordenar slices.
    - Como faz?
    - golang.org/pkg/ → sort
    - godoc.org/sort → examples
    - Sort altera o valor original!
- Exemplo: Ints, Strings.
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	ss := []string{"abacaxi", "caqui", "maçã", "tomate", "laranja"}
	sort.Strings(ss)

	fmt.Println(ss)

	si := []int{2, 15, 7, 18, 10, 12, 4}
	sort.Ints(si)

	fmt.Println(si)
}