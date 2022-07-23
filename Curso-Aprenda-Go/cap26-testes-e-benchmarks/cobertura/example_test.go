/*
"Cobertura" em se tratando de testes se refere a quanto do seu código, percentualmente, está sendo testado. (E antes que alguem fique neurótico querendo 100%: em geral, 70-80% tá ótimo.)
- A flag -cover faz a análise de cobertura do nosso código.
- Podemos utilizar a flag -coverprofile [arquivo] para salvar a análise de cobertura em um arquivo.
- Na prática:
    - go test -cover
    - go test -coverprofile c.out
    - go tool cover -html=c.out ← abre no browser
    - go tool cover -h ← para mais detalhes
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


func BenchmarkSoma(b *testing.B) {
	tests := []test{
		{data: []int{1, 2, 3}, answer: 6},
		{data: []int{1, 2, 3, 4, 5}, answer: 120},
		{data: []int{-1, -2, -3}, answer: -6},
	}

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			Soma(test.data...)
		}
	}
}

func BenchmarkMultiplicacao(b *testing.B) {
	tests := []test{
		{data: []int{1, 2, 3}, answer: 6},
		{data: []int{1, 2, 3, 4, 5}, answer: 120},
		{data: []int{-1, -2, -3}, answer: -6},
	}

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			multiplicacao(test.data...)
		}
	}
}