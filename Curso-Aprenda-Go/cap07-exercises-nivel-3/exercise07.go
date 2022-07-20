/*
Utilizando a solução anterior, adicione as opções else if e else.
*/

package main

import (
	"fmt"
)

func main() {
	exausto := true
	ocupado := true

	if exausto || !ocupado {
		fmt.Println("Cansado")
	} else if !exausto && !ocupado {
		fmt.Println("De boas")
	}
}