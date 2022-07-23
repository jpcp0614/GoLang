package main

import (
	"testing"
)

// soma(i ...int) int

func TestSoma(t *testing.T) {
	teste := soma(1, 2, 3)
	resultado := 6
	if teste != resultado {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}

func TestMultiplicacao(t *testing.T) {
	teste := multiplicacao(1, 2, 3)
	resultado := 6
	if teste != resultado {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}