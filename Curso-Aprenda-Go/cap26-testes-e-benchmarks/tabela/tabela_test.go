/*
Podemos escrever testes em série para testar variedades de situações.
- Exemplo:
    - struct test, fields: data []int, answer int
    - tests := []test{[]int{}, int}
    - range tests
*/

package main

import (
	"testing"
)

type test struct {
	data		[]int
	answer	int
}

func TestSomaEmTabela(t *testing.T) {
	tests := []test{
		{data: []int{1, 2, 3}, answer: 6},
		{data: []int{1, 2, 3, 4, 5}, answer: 15},
		{data: []int{-1, -2, -3}, answer: -6},
	}

	for _, test := range tests {
		teste := soma(test.data...)
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

// func TestSoma(t *testing.T) {
// 	teste := soma(1, 2, 3)
// 	resultado := 6
// 	if teste != resultado {
// 		t.Error("Expected:", resultado, "Got:", teste)
// 	}
// }

// func TestMultiplicacao(t *testing.T) {
// 	teste := multiplicacao(1, 2, 3)
// 	resultado := 6
// 	if teste != resultado {
// 		t.Error("Expected:", resultado, "Got:", teste)
// 	}
// }