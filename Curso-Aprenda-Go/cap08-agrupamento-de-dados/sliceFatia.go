/*
x[:], x[a:], x[:b], x[a:b]
- "a" é incluso;
- "b" não é.
- Exemplo: cabeça magnética de um disco rígido (relógio, fita).
    - Off-by-one error.
- Go Playground: https://play.golang.org/p/i5ZOLKb3Fi
- É fatiando que se deleta um item de uma slice. Na prática:
    - x := append(x[:i], x[:i]...)
	
		- Exercício: tente acessar todos os itens de uma slice *sem* utilizar range.
*/

package main

import (
	"fmt"
)

func main() {
	sabores := []string{"Pepperoni", "Mozzarella", "Frango com Catupiri", "Marguerita", "Quatro Queijos"}

	fatia1 := sabores[:2]
	fmt.Println(fatia1)

	fatia2 := sabores[1:3]
	fmt.Println(fatia2)

	// items do slice sem usar range
	fatia3 := sabores[:]
	fmt.Println(fatia3)

	for i := 0; i < len(sabores); i++ {
		fmt.Println(sabores[i])
	}

	// excluir um sabor
	sabores = append(sabores[:2], sabores[3:]...)
	fmt.Println(sabores)
}