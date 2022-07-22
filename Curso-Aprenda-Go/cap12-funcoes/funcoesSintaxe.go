/*
Qual a utilidade de funções?
    - Abstrair funcionalidade
    - Reutilização de código
- func (receiver) identifier(parameters) (returns) { code }
- A diferença entre parâmetros e argumentos:
    - Funções são definidas com parâmetros
    - Funções são chamadas com argumentos
- Tudo em Go é *pass by value.*
    - Pass by reference, pass by copy, ... não.
- Parâmetro pode ser ...variádico
- Exemplos:
    - Função básica
    - Função que aceita um argumento
    - Função com retorno
    - Função com múltiplos retornos e parâmetro variádico
*/

package main

import (
	"fmt"
)

func main() {
	simples()
	argumento("manhã")
	retorno(2,2)
	soma, qtd, palavra := multParam(1,2,3,4,5)
	fmt.Println(soma, qtd, palavra)
}

func simples() {
	fmt.Println("Hello World!!!")
}

func argumento(s string) {
	if s == "manhã" {
		fmt.Println("Bom dia")
	} else if s == "tarde" {
		fmt.Println("Boa tarde")
	} else {
		fmt.Println("Boa noite")
	}
}

func retorno(x, y int) int {
	fmt.Println(x+y)
	return x + y
}

func multParam(x ...int) (int, int, string) {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x), "funcionou!"
}