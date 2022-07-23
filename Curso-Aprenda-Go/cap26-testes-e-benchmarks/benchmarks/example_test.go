/*
Benchmarks nos permitem testar a velocidade ou performance do nosso código.
- Na prática:
    - Arquivo: _test.go
    - BET: Testes, Exemplos e...
    - func BenchmarkFunc (b *testing.B) { ... b.N ... }
    - go test -bench . ← todos
    - go test -bench Func ← somente Func
- go help testflag
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