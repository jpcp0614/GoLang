/*
Crie um loop utilizando a sintaxe: for {}
- Utilize-o para demonstrar os anos desde que você nasceu.
*/

package main

import (
	"fmt"
)

func main() {
	ano := 1988

	for {
		if ano <= 2022 {
			fmt.Printf("%v ", ano)
			ano ++
		} else {
			break
		}
	}

	// for {
	// 	if ano > 2022 {
	// 		break
	// 	}
	// 	fmt.Printf("%v ", ano)
	// 		ano ++
	// }
}