/*
Um uint16, por exemplo, vai de 0 a 65535.
- Que acontece se a gente tentar usar 65536?
- Ou se a gente estiver em 65535 e tentar adicionar mais 1?
*/

package main

import (
	"fmt"
)

func main() {
	var i uint16
	i = 65535
	fmt.Println(i) // 65535
	i++
	fmt.Println(i) // 0
	i++
	fmt.Println(i) // 1
	i++
	fmt.Println(i) // 2
}