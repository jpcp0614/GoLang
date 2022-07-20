/*
If: bool
- If: o operador não → "!"
- If: declaração de inicialização
*/

package main

import (
	"fmt"
)

// func main() {
// 	x := 10

// 	// operador not (!)
// 	if !(x < 100) {
// 		fmt.Println("Hello!!!")
// 	} else {
// 		fmt.Println("World!!!")
// 	}
// }

func main() {
	if x := 10; x < 100 {
		fmt.Println("Hello!!!")
	} else {
		fmt.Println("World!!!")
	}
}