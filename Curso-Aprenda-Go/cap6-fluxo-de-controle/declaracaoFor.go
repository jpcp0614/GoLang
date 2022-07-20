/*
For: inicialização, condição, pós
- For: condição ("while")
- For: ...ever? (http servers)
- For: break
*/

package main

import (
	"fmt"
)

// func main() {
// 	for x := 0; x < 10; x++ {
// 		fmt.Println(x)
// 	}
// }

// func main() {
// 	x := 0

// 	for x < 10 {
// 		fmt.Println(x)
// 		x++
// 	}
// }

func main() {
	x := 0

	for {
		if x < 10 {
			x++
			fmt.Println(x)
		} else {
		break
		}
	}
	fmt.Println("Quebrou")
}