/*
Outra maneira é fazer testes como exemplos.
- Estes exemplos são os mesmos que aparecem na documentação.
- Para exemplos o formato é "func ExampleFuncao()"
- Deve haver um comentário "// Output: resultado", que é o que será testado
- Para visualizar seu exemplo na documentação, fazemos o de sempre:
    - godoc -http :8080
- Tanto para testes quanto para exemplos podemos utilizar: go test ./...
*/

package main

import (
	"fmt"
	"testing"
)

type test struct {
	data   []int
	answer int
}

func TestSomaEmTabela(t *testing.T) {
	tests := []test{
		{data: []int{1, 2, 3}, answer: 6},
		{data: []int{1, 2, 3, 4, 5}, answer: 15},
		{data: []int{-1, -2, -3}, answer: -6},
	}

	for _, test := range tests {
		teste := Soma(test.data...)
		if teste != test.answer {
			t.Error("Expected:", test.answer, "Got:", teste)
		}
	}
}

func TestMultiplicacaoEmTabela(t *testing.T) {
	tests := []test{
		{data: []int{1, 2, 3}, answer: 6},
		{data: []int{1, 2, 3, 4, 5}, answer: 120},
		{data: []int{-1, -2, -3}, answer: -6},
	}

	for _, test := range tests {
		teste := multiplicacao(test.data...)
		if teste != test.answer {
			t.Error("Expected:", test.answer, "Got:", teste)
		}
	}
}

func ExampleSoma() {
	fmt.Println(Soma(1, 2, 3))
	fmt.Println(Soma(1, 2, 3, 4, 5))
	// Output: 6
	// 15
}
