/*
Utilize a declaração defer de maneira que demonstre que sua execução só ocorre ao final do contexto ao qual ela pertence.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Funções")
	defer fmt.Println(um())
	fmt.Println(dois())
	fmt.Println(tres())
	fmt.Println(quatro())
}

func um() string {
	return "Função 1"
}

func dois() string {
	return "Função dois"
}

func tres() string {
	return "Função três"
}

func quatro() string {
	return "Função quatro"
}